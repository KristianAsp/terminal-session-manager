SHELL:=/bin/bash

build:
	go build -o artifacts/termsesh -ldflags="-X 'main.Version=$(version)'" src/cmd/termsesh.go

test:
	go test ./...

clean:
	rm -rf artifacts/
