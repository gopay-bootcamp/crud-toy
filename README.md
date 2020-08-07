# crud-toy
CRUD operations using ETCD

## Problem Statement

The task was to create a full fledged application that is able to do CRUD with etcd and to integrate logger and config. Addons are implementation of a websocket and a router that can be used to handle grpc requests as well as http requests.

## Prerequisites

1. golang
2. Docker
3. Setting up etcd container
### installation
    docker run -d --name etcd-server \
    --publish 2379:2379 \
    --publish 2380:2380 \
    --env ALLOW_NONE_AUTHENTICATION=yes \
    --env ETCD_ADVERTISE_CLIENT_URLS=http://etcd-server:2379 \
    bitnami/etcd:latest
4. protobuf
### installation
    brew install protobuf
5. protoc-gen-go
### installation
    go get -u google.golang.org/protobuf/cmd/protoc-gen-go

    go install google.golang.org/protobuf/cmd/protoc-gen-go

6. protoc-gen-go-grpc
### installation
    go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

## Development

1. `git clone https://github.com/gopay-bootcamp/crud-toy.git`
2. `cd crud-toy`
3. `go mod tidy`


## Running 

Build the proto file -:

`protoc --go_out=. --go_opt=paths=source_relative  ./procProto/process.proto`

`protoc --go-grpc_out=. --go-grpc_opt=requireUnimplementedServers=false  ./procProto/process.proto` 

To enable and disable Grpc server-:

* Go to the config file `config.yml`, find the field called `grpc_enabled` and set that to `true` or `false` as needed

To create the binaries -:

`make build`

To run the http server -: 

`_output/bin/server start`

To run the grpc server -:

`_output/bin/server grpcStart`

To run the client -:

`_output/bin/cli <command> <args>`



