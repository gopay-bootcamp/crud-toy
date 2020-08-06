# crud-toy
CRUD operations using ETCD

## Problem Statement

The task was to create a full fledged application that is able to do CRUD with etcd and to integrate logger and config. Addons are implementation of a websocket and a router that can be used to handle grpc requests as well as http requests.

## Prerequisites

1. golang
2. protoc-gen-go
### installation
    `go get -u google.golang.org/protobuf/cmd/protoc-gen-go`
    `go install google.golang.org/protobuf/cmd/protoc-gen-go`

3. protoc-gen-go-grpc
### installation
    `go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc`
    `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc`

## Development

1. `git clone https://github.com/gopay-bootcamp/crud-toy.git`
2. `cd crud-toy`
3. `go mod tidy`
4. `make build`


## Running 

Build the proto file -:

`protoc --go_out=. --go_opt=paths=source_relative  ./procProto/process.proto`
`protoc --go-grpc_out=. --go-grpc_opt=requireUnimplementedServers=false  ./procProto/process.proto` 

To run the server -: 

`_output/bin/server start`

To run the client -:

`_output/bin/client <command> <args>`



