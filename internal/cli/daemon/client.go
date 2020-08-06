package daemon

import (
	"crud-toy/procProto"
	"crud-toy/config"
	"crud-toy/internal/cli/client"
	"errors"
	"google.golang.org/grpc"
)

var proctorDClient client.Client
var proctorDClientGrpc client.GrpcClient
var (
	grpcEnabled = config.Config().GrpcEnabled
)

func StartClient() {
	if grpcEnabled == false{
		proctorDClient = client.NewClient()
	} else{
		conn, _ := grpc.Dial("localhost:8000",grpc.WithInsecure())
		grpcclient := procProto.NewProcServiceClient(conn)
		proctorDClientGrpc = client.NewGrpcClient(grpcclient)
	}
}

func FindProcs(args []string) error {
	if len(args) < 1 {
		return errors.New("Invalid argument error")
	}
	id := args[0]
	if grpcEnabled == false{
		return proctorDClient.FindProcs(id)
	} else {
		return proctorDClientGrpc.FindProcsGrpc(id)
	}
}

func CreateProcs(args []string) error {
	if len(args) < 2 {
		return errors.New("Invalid argument error")
	}
	name := args[0]
	author := args[1]
	if grpcEnabled == false{
		return proctorDClient.CreateProcs(name, author)
	} else {
		return proctorDClientGrpc.CreateProcsGrpc(name,author)
	}
}

func UpdateProcs(args []string) error {
	if len(args) < 3 {
		return errors.New("Invalid argument error")
	}
	id := args[0]
	name := args[1]
	author := args[2]

	if grpcEnabled == false{
		return proctorDClient.UpdateProcs(id, name, author)
	} else {
		return proctorDClientGrpc.UpdateProcsGrpc(id,name,author)
	}
}

func DeleteProcs(args []string) error {
	if len(args) < 1 {
		return errors.New("Invalid argument error")
	}
	id := args[0]
	if grpcEnabled == false{
		return proctorDClient.DeleteProcs(id)
	} else {
		return proctorDClientGrpc.DeleteProcsGrpc(id)
	}
}

func ReadAllProcs() error {
	return proctorDClient.ReadAllProcs()
}
