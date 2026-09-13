# api/Makefile
VERSION ?= $(shell git rev-parse --short HEAD)
.PHONY: build test run lint
build:
	go build -ldflags "-X main.version=$(VERSION)" -o bin/api ./cmd/api
test:
	go test -cover ./...
run: build
	./bin/api
lint:
	go vet ./...
