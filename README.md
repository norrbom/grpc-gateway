# Go gRPC Gateway

Starting point for a gRPC Gateway functioning as a single entry point for clients.

Features:
- Acts as a proxy, receiving grpc calls from clients and routing them to the appropriate backend.
- Implements caching and retry mechanisms; masks backend network errors by retrying from an in-memory queue that spool to file when the queue length or memory limits are reached.

TODO:
- Implement spool to file or local database
- Send to backend concurrently
- Configuration via .env

## Quick Start

```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

Build protobuf and run the server:

```shell
# builds protobuf and binary
make
go run .
```

In another terminal, send a ping message to the server:

```shell
go run . --client
```

Docker:

```shell
make docker
docker run -p 50051:50051 go-grpc-gateway
```