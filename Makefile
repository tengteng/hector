GOPATH := ${PWD}:${PWD}/_vendor
export GOPATH

default: install

install: fmt
	go install ./...

build: fmt vet
	go build ./...

fmt:
	go fmt ./...

vet:
	go vet ./...
