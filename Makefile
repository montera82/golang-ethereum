.PHONY: all deps test build

all: deps test build

deps:
	@go mod vendor
	@go mod tidy

test:
	@go vet ./{cmd,handler}/...
	@go test -v -race -cover ./{cmd,handler}/...

build:
	@GOBIN=/build go install -ldflags "-w -s" ./... # needs to be looked into

install: 
	@go install ./cmd/golang-ethereum
