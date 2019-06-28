# Arguments
ARG BUILD_DATE
ARG VERSION
ARG VCS_REF

## -------------------------------------------------------------------------------------------------

FROM golang:1.12 as builder

RUN set -eux; \
    apt-get update -y && \
    apt-get install -y apt-utils upx zip unzip;

# Create a non-root privilege account to build
RUN adduser --disabled-password --gecos "" -u 1000 golang && \
    mkdir -p $GOPATH/src/workspace && \
    chown -R golang:golang $GOPATH/src/workspace;

# Force go modules
ENV GO111MODULE=on
ENV GOPROXY=https://proxy.golang.org/

WORKDIR $GOPATH/src/workspace

# Prepare an unprivilegied user for run
RUN set -eux; \
    echo 'nobody:x:65534:65534:nobody:/:' > /tmp/passwd && \
    echo 'nobody:x:65534:' > /tmp/group && \
    mkdir /tmp/.config && \
    chown 65534:65534 /tmp/.config

# Drop privileges to build
USER golang
COPY --chown=golang:golang mage.go .
COPY --chown=golang:golang tools tools/

# Install tools
RUN set -eux; \
    go run mage.go -d tools 

# Copy project go module
COPY --chown=golang:golang magefile.go go.mod ./

# Install dependencies
RUN set -eux; \
    go run mage.go go:deps

# Copy all sources
COPY --chown=golang:golang . .

# Build final target
RUN set -eux; \
    go run mage.go

# Compress binaries
RUN set -eux; \
    upx -9 bin/* && \
    chmod +x bin/*

## -------------------------------------------------------------------------------------------------

FROM gcr.io/distroless/static:latest

# Arguments
ARG BUILD_DATE
ARG VERSION
ARG VCS_REF

# Metadata
LABEL \
    org.label-schema.build-date=$BUILD_DATE \
    org.label-schema.name="IOC Platform Agent" \
    org.label-schema.description="Agent is a feed converter." \
    org.label-schema.url="https://github.com/iocplatform/agent" \
    org.label-schema.vcs-url="https://github.com/iocplatform/agent.git" \
    org.label-schema.vcs-ref=$VCS_REF \
    org.label-schema.vendor="Thibault NORMAND" \
    org.label-schema.version=$VERSION \
    org.label-schema.schema-version="1.0" \
    org.zenithar.licence="MIT"

COPY --from=builder /go/src/workspace/bin/agent /usr/bin/agent
COPY --from=builder /tmp/group /tmp/passwd /etc/
COPY --from=builder --chown=65534:65534 /tmp/.config /

USER nobody:nobody
WORKDIR /

ENTRYPOINT [ "/usr/bin/agent" ]
CMD ["--help"]