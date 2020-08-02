package etcd

import (
	"github.com/stretchr/testify/mock"
	"github.com/coreos/etcd/clientv3"
	"context"
)

type ClientMock struct {
	mock.Mock
}

func (m ClientMock) DeleteKey(ctx context.Context, key string) error {
	args := m.Called(ctx,key)
	return args.Error(0)
}

func (m ClientMock) PutValue(ctx context.Context, key string, value string) (*clientv3.PutResponse, error) {
	args := m.Called(ctx,key,value)
	return args.Get(0).(*clientv3.PutResponse),args.Error(1)
}

func (m ClientMock) GetValue(ctx context.Context, key string) (*clientv3.GetResponse, error) {
	args := m.Called(ctx,key)
	return args.Get(0).(*clientv3.GetResponse),args.Error(1)
}

func (m ClientMock) GetAllValueWithPrefix(ctx context.Context, key string) (*clientv3.GetResponse, error) {
	args := m.Called(ctx,key)
	return args.Get(0).(*clientv3.GetResponse),args.Error(1)
}

func (m ClientMock) GetValueWithRevision(ctx context.Context, key string, pr *clientv3.PutResponse) (*clientv3.GetResponse, error) {
	args := m.Called(ctx,key,pr)
	return args.Get(0).(*clientv3.GetResponse),args.Error(1)
}