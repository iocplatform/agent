# Agent

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
