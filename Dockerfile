FROM docker.io/golang:1.23.4-alpine as builder

WORKDIR /usr/src/app

RUN apk update
RUN apk add --no-cache make bash protobuf
RUN make --version
RUN protoc --version
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

COPY go.mod go.sum ./
RUN go mod download
RUN go mod verify

COPY *.go .
COPY server/server.proto ./server/server.proto
COPY Makefile .

RUN make

FROM docker.io/golang:1.23.4-alpine
COPY --from=builder /usr/src/app/bin/go-grpc-gateway /usr/local/bin/

WORKDIR /usr/local/bin

COPY .env .

RUN adduser \
  --disabled-password \
  --gecos "" \
  --home "/nonexistent" \
  --shell "/sbin/nologin" \
  --no-create-home \
  --uid 65532 \
  user

USER user:user

CMD ["go-grpc-gateway"]