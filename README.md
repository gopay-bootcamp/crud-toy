# crud-toy
CRUD operations using ETCD

## Problem Statement

The task was to create a full fledged application that is able to do CRUD with etcd and to integrate logger and config. Addons are implementation of a websocket and a router that can be used to handle grpc requests as well as http requests.

## Development

1. `git clone https://github.com/gopay-bootcamp/crud-toy.git`
2. `cd crud-toy`
3. `make build`


## Running 

Build the proto file -:

`protoc --go_out=. --go_opt=paths=source_relative  ./procProto/process.proto`

To run the server -: 

`_output/bin/server start`

To run the client -:

`_output/bin/client <command> <args>`



