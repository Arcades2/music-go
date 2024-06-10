.DEFAULT_GOAL := build
APP_NAME := main
CMD_DIR := cmd/$(APP_NAME)

fmt:
	@echo "Running go fmt"
	go fmt ./...
.PHONY:fmt

lint: fmt
	@echo "Running golangci-lint"
	golangci-lint run
.PHONY:lint

vet: lint
	@echo "Running go vet"
	go vet ./...
.PHONY:vet

test: vet
	@echo "Running go test"
	go test -v ./...
.PHONY:test

build: vet
	@echo "Building binary"
	go build -o $(APP_NAME) $(CMD_DIR)/main.go
.PHONY:build
