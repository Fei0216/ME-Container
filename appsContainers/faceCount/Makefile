.PHONY: build

GO=CGO_ENABLED=1 GO111MODULE=on go

GIT_SHA=$(shell git rev-parse HEAD)

build:
	$(GO) build $(GOFLAGS) -o app-service main.go NetConn.go

