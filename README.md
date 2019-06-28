# Agent

[![CircleCI](https://circleci.com/gh/iocplatform/agent.svg?style=svg)](https://circleci.com/gh/iocplatform/agent)
[![Go Report Card](https://goreportcard.com/badge/github.com/iocplatform/agent)](https://goreportcard.com/report/github.com/iocplatform/agent)
[![LICENSE](https://img.shields.io/github/license/iocplatform/agent.svg)](https://github.com/iocplatform/agent/blob/master/LICENSE)

Agent is a short-lived job to be scheduled to process a remote feed declared in 
dedicated descriptor.

## Build

### Local 

```sh
> go run mage.go
```

### Docker

```sh
> export DOCKER_BUILDKIT=1
> go run mage.go docker:build
```
