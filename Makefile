.DEFAULT_GOAL := build

fmt:
	go fmt
.PHONY: fmt

lint: fmt
	staticcheck
.PHONY: lint

vet: lint
	go vet
.PHONY: vet

build: vet
	go mod tidy
	CGO_ENABLED=0 go build
.PHONY: build
