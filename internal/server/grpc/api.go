package grpc

import (
	"crud-toy/logger"
	// "crud-toy/internal/model"
	"crud-toy/procProto"
	"net"
	"google.golang.org/grpc"
)


func Start() error {
	listener, err:=net.Listen("tcp",":8000")
	server := grpc.NewServer()

	procProto.RegisterProcServiceServer(server,&procServiceServer{})
	if err!=nil{
		logger.Fatal("grpc server not started")
		return err
	}
	logger.Info("grpc server started on port 8000")
	server.Serve(listener)
	return nil
}
