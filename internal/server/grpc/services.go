package grpc

import (
	"context"
	"crud-toy/internal/model"
	"crud-toy/internal/server/execution"
	"crud-toy/procProto"
)

type procServiceServer struct {
	procExec execution.Execution
}

func NewProcServiceServer(exec execution.Execution) procProto.ProcServiceServer {
	return &procServiceServer{
		procExec: exec,
	}
}

func (s *procServiceServer) CreateProc(ctx context.Context, request *procProto.RequestForCreateProc) (*procProto.ProcID, error) {
	var proc model.Proc
	proc.Name = request.Name
	proc.Author = request.Author
	id, err := s.procExec.CreateProc(ctx, &proc)
	if err != nil {
		return nil, err
	}
	resp := &procProto.ProcID{Value: id, Message: "successfully created Proc"}
	return resp, nil
}

func (s *procServiceServer) ReadProcByID(ctx context.Context, request *procProto.RequestForReadByID) (*procProto.Proc, error) {
	id := request.ID
	proc := model.Proc{
		ID: id,
	}
	result, err := s.procExec.ReadProcByID(ctx, &proc)
	if err != nil {
		return nil, err
	}
	resp := &procProto.Proc{ID: result.ID, Name: result.Name, Author: result.Author}
	return resp, nil
}

func (s *procServiceServer) UpdateProcByID(ctx context.Context, request *procProto.RequestForUpdateProcByID) (*procProto.ProcID, error) {
	id := request.ID
	proc := &model.Proc{ID: id, Name: request.Name, Author: request.Author}
	id, err := s.procExec.UpdateProc(ctx, proc)
	if err != nil {
		return nil, err
	}
	resp := &procProto.ProcID{Value: id, Message: "successfully updated Proc"}
	return resp, nil
}

func (s *procServiceServer) DeleteProcByID(ctx context.Context, request *procProto.RequestForDeleteByID) (*procProto.ProcID, error) {
	id := request.ID
	proc := model.Proc{
		ID: id,
	}
	err := s.procExec.DeleteProc(ctx, &proc)
	if err != nil {
		return nil, err
	}
	resp := &procProto.ProcID{Value: id, Message: "successfully deleted Proc"}
	return resp, nil
}

func (s *procServiceServer) ReadAllProcs(ctx context.Context, request *procProto.RequestForReadAllProcs) (*procProto.ProcList, error) {
	procList, err := s.procExec.ReadAllProc(ctx)
	if err != nil {
		return nil, err
	}
	protoProcs := []*procProto.Proc{}
	for _, proc := range procList {
		protoProc := procProto.Proc{ID: proc.ID, Name: proc.Name, Author: proc.Author}
		protoProcs = append(protoProcs, &protoProc)
	}
	return &procProto.ProcList{Procs: protoProcs}, nil
}
