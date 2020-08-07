package client

import (
	"context"
	io "crud-toy/internal/cli/printer"
	"crud-toy/internal/model"
	"crud-toy/procProto"
	"fmt"
	"time"
)

type grpcClient struct {
	client                procProto.ProcServiceClient
	printer               io.Printer
	connectionTimeoutSecs time.Duration
}

func NewGrpcClient(client procProto.ProcServiceClient) Client {
	return &grpcClient{
		client:                client,
		printer:               io.PrinterInstance,
		connectionTimeoutSecs: time.Second,
	}
}

func (c *grpcClient) FindProcs(id string) error {
	request := procProto.RequestForReadByID{ID: id}
	proc, err := c.client.ReadProcByID(context.Background(), &request)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Printf("Id: %s, Name: %s, Author: %s", proc.ID, proc.Name, proc.Author)
	return nil
}

func (c *grpcClient) CreateProcs(name string, author string) error {
	request := procProto.RequestForCreateProc{Name: name, Author: author}
	proc, err := c.client.CreateProc(context.Background(), &request)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Printf("Id: %s, Message: %s", proc.Value, proc.Message)
	return nil
}

func (c *grpcClient) DeleteProcs(id string) error {
	request := procProto.RequestForDeleteByID{ID: id}
	proc, err := c.client.DeleteProcByID(context.Background(), &request)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Printf("Id: %s, Message: %s", proc.Value, proc.Message)
	return nil
}

func (c *grpcClient) UpdateProcs(id string, name string, author string) error {
	request := procProto.RequestForUpdateProcByID{ID: id, Name: name, Author: author}
	proc, err := c.client.UpdateProcByID(context.Background(), &request)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Printf("Id: %s, Message: %s", proc.Value, proc.Message)
	return nil
}

func (c *grpcClient) ReadAllProcs() error {
	request := procProto.RequestForReadAllProcs{}
	protoProcList, err := c.client.ReadAllProcs(context.Background(), &request)
	if err != nil {
		return err
	}
	procList := []model.Proc{}
	for _, protoProc := range protoProcList.Procs {
		proc := model.Proc{ID: protoProc.ID, Name: protoProc.Name, Author: protoProc.Author}
		procList = append(procList, proc)
	}
	c.printer.PrintTable(procList)
	return nil
}
