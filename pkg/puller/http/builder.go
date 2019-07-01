package http

import (
	"github.com/iocplatform/agent/pkg/puller"
	"github.com/iocplatform/agent/pkg/puller/api"
)

// Build a collector from a parameter map
func Build(parameters map[string]interface{}) (api.Puller, error) {
	var options []Option

	for key, value := range parameters {
		switch key {
		case "url":
			options = append(options, WithURL(value.(string)))
		case "method":
			options = append(options, WithMethod(value.(string)))
		case "parameters":
			parameters := map[string]string{}
			for k, v := range value.(map[interface{}]interface{}) {
				parameters[k.(string)] = v.(string)
			}
			options = append(options, WithParameters(parameters))
		case "headers":
			headers := map[string]string{}
			for k, v := range value.(map[interface{}]interface{}) {
				headers[k.(string)] = v.(string)
			}
			options = append(options, WithHeaders(headers))
		case "skip_tls_verify":
			options = append(options, WithTLSSkipVerify(value.(bool)))
		case "br2nl":
			options = append(options, WithBR2NLPreprocessor(value.(bool)))
		default:
		}
	}

	return New(options...)
}

func init() {
	puller.Add("http", func(parameters map[string]interface{}) (api.Puller, error) {
		return Build(parameters)
	})
}
