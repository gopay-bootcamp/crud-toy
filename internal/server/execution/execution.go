package execution

import (
	"context"
	"crud-toy/internal/model"
	"crud-toy/internal/server/db/etcd"
	crypto_rand "crypto/rand"
	"encoding/binary"
	math_rand "math/rand"
	"strconv"
	"time"
)

var (
	timeout time.Duration = 2 * time.Second
)

type Execution interface {
	CreateProc(proc *model.Proc) (string, error)
	ReadProcByID(proc *model.Proc) (*model.Proc, error)
	ReadAllProc() ([]model.Proc, error)
	UpdateProc(proc *model.Proc) (string, error)
	DeleteProc(proc *model.Proc) error
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

func (e *execution) CreateProc(proc *model.Proc) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
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

func (e *execution) ReadProcByID(proc *model.Proc) (*model.Proc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	id := proc.ID
	result, err := e.client.GetValue(ctx, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (e *execution) ReadAllProc() ([]model.Proc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	procs, err := e.client.GetAllValues(ctx)
	if err != nil {
		return nil, err
	}
	return procs, nil
}

func (e *execution) UpdateProc(proc *model.Proc) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	result, err := e.client.PutValue(ctx, proc.ID, proc)
	if err != nil {
		return "", err
	}
	return result, nil
}

func (e *execution) DeleteProc(proc *model.Proc) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	id := proc.ID
	err := e.client.DeleteKey(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

