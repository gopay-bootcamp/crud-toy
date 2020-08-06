package grpc

import (
	"crud-toy/internal/server/db/etcd"
	"crud-toy/internal/model"
	"context"
	"crud-toy/procProto"
	crypto_rand "crypto/rand"
	"encoding/binary"
	math_rand "math/rand"
	"strconv"
)

type procServiceServer struct {
}


func (s *procServiceServer) CreateProc(ctx context.Context,request *procProto.RequestForCreateProc) (*procProto.ProcID,error){
	var proc model.Proc
	proc.Name=request.Name
	proc.Author=request.Author
	var b [8]byte
	crypto_rand.Read(b[:])
	math_rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
	proc.ID = strconv.FormatInt(math_rand.Int63(), 10)
	etcdClient := etcd.NewClient()
	defer etcdClient.Close()
	id,err:= etcdClient.PutValue(ctx,proc.ID,&proc)
	if err !=nil {
		return nil,err
	}
	resp := &procProto.ProcID{Value:id,Message:"successfully created Proc"}
	return resp,nil
}

func (s *procServiceServer) ReadProcByID(ctx context.Context,request *procProto.RequestForReadByID) (*procProto.Proc,error){
	id := request.ID
	etcdClient := etcd.NewClient()
	defer etcdClient.Close()
	proc,err := etcdClient.GetValue(ctx,id)
	if err !=nil {
		return nil,err
	}
	resp := &procProto.Proc{ID:proc.ID,Name:proc.Name,Author:proc.Author}
	return resp,nil
}

func (s *procServiceServer) UpdateProcByID(ctx context.Context,request *procProto.RequestForUpdateProcByID) (*procProto.ProcID,error){
	id := request.ID
	proc := &model.Proc{ID:id,Name:request.Name,Author:request.Author}
	etcdClient := etcd.NewClient()
	defer etcdClient.Close()
	id,err:= etcdClient.PutValue(ctx,id,proc)
	if err !=nil {
		return nil,err
	}
	resp := &procProto.ProcID{Value:id,Message:"successfully updated Proc"}
	return resp,nil
}

func (s *procServiceServer) DeleteProcByID(ctx context.Context,request *procProto.RequestForDeleteByID) (*procProto.ProcID,error){
	id := request.ID
	etcdClient := etcd.NewClient()
	defer etcdClient.Close()
	err := etcdClient.DeleteKey(ctx,id)
	if err !=nil {
		return nil,err
	}
	resp := &procProto.ProcID{Value:id,Message:"successfully deleted Proc"}
	return resp,nil
}

