.PHONY: all deps test build

all: deps test build

deps:
	@go mod vendor
	@go mod tidy

test:
	@go vet ./{cmd,handler}/...
	@go test -v -race -cover ./{cmd,handler}/...

build-go-files:
	@GOBIN=/build go install -ldflags "-w -s" ./... # needs to be looked into
build-solidity-files:
	@solc --abi contracts/Store.sol -o build --overwrite #generate abi from contract
	@solc --bin contracts/Store.sol -o build --overwrite #complile sol to evm binary format
	@abigen --abi=./build/Store.abi --pkg=contracts --out=contracts/store.go
	@abigen --bin=./build/Store.bin --abi=./build/Store.abi --pkg=contracts --out=contracts/store.go

install: 
	@go install ./cmd/golang-ethereum
