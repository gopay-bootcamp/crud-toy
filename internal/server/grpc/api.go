package grpc

import (
	"crud-toy/internal/server/db/etcd"
	"crud-toy/internal/server/execution"
	"crud-toy/logger"

	// "crud-toy/internal/model"
	"crud-toy/procProto"
	"net"

	"google.golang.org/grpc"
)

func Start() error {
	listener, err := net.Listen("tcp", ":8000")
	server := grpc.NewServer()
	etcdClient := etcd.NewClient()
	defer etcdClient.Close()
	exec := execution.NewExec(etcdClient)

	procGrpcServer := NewProcServiceServer(exec)
	procProto.RegisterProcServiceServer(server, procGrpcServer)
	if err != nil {
		logger.Fatal("grpc server not started")
		return err
	}
	logger.Info("grpc server started on port 8000")
	server.Serve(listener)
	return nil
}
