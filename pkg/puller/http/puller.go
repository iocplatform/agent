package http

import (
	"archive/zip"
	"bufio"
	"bytes"
	"compress/gzip"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/iocplatform/agent/pkg/dispatcher"
	"github.com/iocplatform/agent/pkg/puller/api"

	"github.com/minio/blake2b-simd"
	"github.com/sethgrid/pester"
	"go.zenithar.org/pkg/log"
	"golang.org/x/xerrors"
)

type httpPuller struct {
	api.DefaultPuller

	dopts      Options
	httpClient *http.Client
	client     *pester.Client
}

// New returns an HTTP puller setted up with given options
func New(opts ...Option) (api.Puller, error) {
	puller := &httpPuller{
		DefaultPuller: api.DefaultPuller{
			Name: "http",
		},
	}

	// Process build options
	for _, opt := range opts {
		opt(&puller.dopts)
	}

	// Internals
	puller.httpClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: puller.dopts.tlsSkipVerify},
			Proxy:           http.ProxyFromEnvironment,
		},
	}

	// Resilient http client
	client := pester.NewExtendedClient(puller.httpClient)
	client.Backoff = pester.ExponentialBackoff
	client.MaxRetries = 5
	client.KeepLog = true
	puller.client = client

	// Return puller
	return puller, nil
}

// -----------------------------------------------------------------------------

func (c *httpPuller) Pull(ctx context.Context, d dispatcher.Dispatcher) error {
	log.For(ctx).Info("Requesting file usign http puller")

	// Prepare the http query
	bodyReader, responseTime, lastModifiedDate, err := c.sendRequest(ctx)
	if err != nil {
		return xerrors.Errorf("http: unable to query feed: %w", err)
	}
	defer func(closer io.Closer) {
		log.SafeClose(closer, "Unable to close request body")
	}(bodyReader)

	// Prepare response
	result := map[string]interface{}{}

	result["timestamp"] = lastModifiedDate.Unix()
	result["puller"] = "http"
	result["url"] = c.dopts.url

	// Body fingerprinting
	h := blake2b.New512()
	contentReader := io.TeeReader(bodyReader, h)

	// Timers
	if responseTime > 0 {
		result["response_time"] = responseTime
	}

	// Parse as JSON payload
	if c.dopts.json {
		var payload map[string]interface{}

		// Decode JSON
		err := c.readJSON(contentReader, &payload)
		if err != nil {
			return err
		}

		// Extract lines using jmesPath expression
		lines, err := c.dopts.jmesPath.Search(&payload)
		if err != nil {
			return xerrors.Errorf("Unable to extract an array from given JMESPath expression: %w", err)
		}

		result["content_type"] = "lines"
		result["body"] = lines
	} else {
		lines, err := c.readLines(contentReader)
		if err != nil {
			return err
		}
		result["content_type"] = "lines"
		result["line_count"] = len(lines)
		result["body"] = lines
	}

	// Assign body fingerprint
	result["hash"] = fmt.Sprintf("%x", h.Sum(nil))

	// No error
	return d.Dispatch(ctx, result)
}

func (c *httpPuller) SetParameters(parameters map[string]interface{}) {
	var options []Option

	for key, value := range parameters {
		switch key {
		case "url":
			options = append(options, WithURL(value.(string)))
		case "method":
			options = append(options, WithMethod(value.(string)))
		case "parameters":
			options = append(options, WithParameters(value.(map[string]string)))
		default:
		}
	}

	// Assign all options
	for _, opt := range options {
		opt(&c.dopts)
	}
}

// -----------------------------------------------------------------------------

func (c *httpPuller) sendRequest(ctx context.Context) (io.ReadCloser, float64, *time.Time, error) {
	// Prepare URL
	requestURL, err := url.Parse(c.dopts.url)
	if err != nil {
		return nil, -1, nil, fmt.Errorf("Invalid server URL \"%s\"", c.dopts.url)
	}

	// Add query params
	params := requestURL.Query()
	for k, v := range c.dopts.parameters {
		params.Add(k, v)
	}
	requestURL.RawQuery = params.Encode()

	// Create + send request
	req, err := http.NewRequest(c.dopts.method, requestURL.String(), nil)
	if err != nil {
		return nil, -1, nil, err
	}

	// Add custom headers
	for k, v := range c.dopts.headers {
		req.Header.Add(k, v)
	}

	// Add gzip encoding support
	req.Header.Add("Accept-Encoding", "gzip")

	start := time.Now()
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, -1, nil, err
	}

	// Calculate the response time
	responseTime := time.Since(start).Seconds()

	// Process response
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("Response from url \"%s\" has status code %d (%s), expected %d (%s)",
			requestURL.String(),
			resp.StatusCode,
			http.StatusText(resp.StatusCode),
			http.StatusOK,
			http.StatusText(http.StatusOK))
		return nil, responseTime, nil, err
	}

	// Retrieve last-modified header for observable timestamp
	lastModifiedRaw := resp.Header.Get("Last-Modified")
	lastModifiedDate := time.Now().UTC()

	if len(strings.TrimSpace(lastModifiedRaw)) > 0 {
		lastModifiedDate, err = time.Parse(http.TimeFormat, lastModifiedRaw)
		if err != nil {
			lastModifiedDate = time.Now().UTC()
		}
	}

	// Check that the server actually sent compressed data
	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			checkClose(resp.Body, &err)
			return nil, responseTime, &lastModifiedDate, err
		}
	default:
		reader = resp.Body
	}

	// Check that the server actually sent a compressed file
	var bodyReader io.ReadCloser

	switch resp.Header.Get("Content-Type") {
	case "application/x-gzip":
		bodyReader, err = gzip.NewReader(reader)
		if err != nil {
			checkClose(reader, &err)
			return nil, responseTime, &lastModifiedDate, err
		}
	case "application/zip":
		var zipReader *zip.Reader
		b, err := ioutil.ReadAll(reader)
		zipReader, err = zip.NewReader(bytes.NewReader(b), resp.ContentLength)
		if err != nil {
			return nil, responseTime, &lastModifiedDate, err
		}

		if len(zipReader.File) > 1 {
			return nil, responseTime, nil, fmt.Errorf("Unable to handle multiple files archive")
		}

		bodyReader, err = zipReader.File[0].Open()
		if err != nil {
			checkClose(reader, &err)
			return nil, responseTime, nil, err
		}
	default:
		bodyReader = reader
	}

	return bodyReader, responseTime, &lastModifiedDate, err
}

// readLines reads the response into an array of strings.
func (c *httpPuller) readLines(bodyReader io.Reader) (lines []string, err error) {
	reader := bufio.NewReader(bodyReader)
	buffer := bytes.NewBuffer(make([]byte, 0, 128))
	var part []byte
	var prefix bool
	for {
		if part, prefix, err = reader.ReadLine(); err != nil {
			break
		}
		buffer.Write(part)
		if !prefix {
			lines = append(lines, buffer.String())
			buffer.Reset()
		}
	}
	if err == io.EOF {
		err = nil
	}
	return
}

// readJson reads the response into the json type passed in
func (c *httpPuller) readJSON(bodyReader io.Reader, result interface{}) (err error) {
	decoder := json.NewDecoder(bodyReader)
	return decoder.Decode(result)
}

// -----------------------------------------------------------------------------
