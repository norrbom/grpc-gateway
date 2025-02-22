SHELL := /bin/bash
BIN_DIR := ./bin
APP_NAME := go-grpc-gateway

all: main

main: main.go protoc
	go mod tidy
	GOOS=linux GOARCH=amd64 go build -o $(BIN_DIR)/$(APP_NAME) .

.PHONY: protoc
protoc:
	protoc \
	--go_out=. \
	--go_opt=paths=source_relative \
    --go-grpc_out=. \
	--go-grpc_opt=paths=source_relative \
    server/server.proto

.PHONY: test
test:
	go test -v ./...

.PHONY: clean
clean:
	go clean -cache
	rm -rf $(BIN_DIR)

.PHONY: docker
docker:
	docker build --platform=linux/amd64 -t $(APP_NAME) .

