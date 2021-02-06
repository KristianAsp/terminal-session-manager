SHELL:=/bin/bash

all: clean build test

GOPATH = ${HOME}/go

build:
	go install github.com/go-bindata/go-bindata/...
	$(GOPATH)/bin/go-bindata -o generated/resources.go -pkg embedded resources resources/properties
	go build -o artifacts/termsesh -ldflags="-X 'main.Version=$(version)'" src/cmd/termsesh.go

test:
	$(GOPATH)/bin/go-bindata -o generated/resources.go -pkg embedded resources resources/properties
	go test ./...

clean:
	rm -rf artifacts/
