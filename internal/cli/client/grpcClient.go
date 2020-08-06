package client

import (
	"context"
	"crud-toy/procProto"
	"crud-toy/internal/cli/printer"
	"fmt"
	"time"
)

type GrpcClient interface {
	FindProcsGrpc(string) error
	CreateProcsGrpc(string, string) error
	DeleteProcsGrpc(string) error
	UpdateProcsGrpc(string, string, string) error
}

type grpcClient struct {
	client procProto.ProcServiceClient
	printer               io.Printer
	connectionTimeoutSecs time.Duration
}


func NewGrpcClient(client procProto.ProcServiceClient) GrpcClient {
	return &grpcClient{
		client : client,
		printer:               io.PrinterInstance,
		connectionTimeoutSecs: time.Second,
	}
}

func (c *grpcClient) FindProcsGrpc(id string) error {
	request:=procProto.RequestForReadByID{ID:id}
	proc, err := c.client.ReadProcByID(context.Background(),&request)
 	if err != nil {
		 fmt.Println(err.Error())
		 return err
	}
	fmt.Printf("Id: %s, Name: %s, Author: %s",proc.ID,proc.Name,proc.Author)
	return nil
}

func (c *grpcClient) CreateProcsGrpc(name string, author string) error {
	request:=procProto.RequestForCreateProc{Name:name,Author:author}
	proc, err := c.client.CreateProc(context.Background(),&request)
 	if err != nil {
		 fmt.Println(err.Error())
		 return err
	}
	fmt.Printf("Id: %s, Message: %s",proc.Value,proc.Message)
	return nil
}

func (c *grpcClient) DeleteProcsGrpc(id string) error {
	request:=procProto.RequestForDeleteByID{ID:id}
	proc, err := c.client.DeleteProcByID(context.Background(),&request)
 	if err != nil {
		 fmt.Println(err.Error())
		 return err
	}
	fmt.Printf("Id: %s, Message: %s",proc.Value,proc.Message)
	return nil
}

func (c *grpcClient) UpdateProcsGrpc(id string, name string, author string) error {
	request:=procProto.RequestForUpdateProcByID{ID:id,Name:name,Author:author}
	proc, err := c.client.UpdateProcByID(context.Background(),&request)
 	if err != nil {
		 fmt.Println(err.Error())
		 return err
	}
	fmt.Printf("Id: %s, Message: %s",proc.Value,proc.Message)
	return nil
}