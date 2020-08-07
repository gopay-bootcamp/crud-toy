package execution

import (
	"context"
	"crud-toy/internal/model"
	"crud-toy/internal/server/db/etcd"
	crypto_rand "crypto/rand"
	"encoding/binary"
	math_rand "math/rand"
	"strconv"
)

type Execution interface {
	CreateProc(ctx context.Context, proc *model.Proc) (string, error)
	ReadProcByID(ctx context.Context, proc *model.Proc) (*model.Proc, error)
	ReadAllProc(ctx context.Context) ([]model.Proc, error)
	UpdateProc(ctx context.Context, proc *model.Proc) (string, error)
	DeleteProc(ctx context.Context, proc *model.Proc) error
}

type execution struct {
	client etcd.EtcdClient
	ctx    context.Context
	cancel context.CancelFunc
}

func NewExec(dbClient etcd.EtcdClient) Execution {
	return &execution{
		client: dbClient,
	}
}

func (e *execution) CreateProc(ctx context.Context, proc *model.Proc) (string, error) {
	var b [8]byte
	crypto_rand.Read(b[:])
	math_rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
	proc.ID = strconv.FormatInt(math_rand.Int63(), 10)
	result, err := e.client.PutValue(ctx, proc.ID, proc)
	if err != nil {
		return "", err
	}
	return result, nil
}

func (e *execution) ReadProcByID(ctx context.Context, proc *model.Proc) (*model.Proc, error) {
	id := proc.ID
	result, err := e.client.GetValue(ctx, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (e *execution) ReadAllProc(ctx context.Context) ([]model.Proc, error) {
	procs, err := e.client.GetAllValues(ctx)
	if err != nil {
		return nil, err
	}
	return procs, nil
}

func (e *execution) UpdateProc(ctx context.Context, proc *model.Proc) (string, error) {
	result, err := e.client.PutValue(ctx, proc.ID, proc)
	if err != nil {
		return "", err
	}
	return result, nil
}

func (e *execution) DeleteProc(ctx context.Context, proc *model.Proc) error {
	id := proc.ID
	err := e.client.DeleteKey(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
