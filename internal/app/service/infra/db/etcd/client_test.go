package etcd

import (
	"fmt"
	//"github.com/stretchr/testify/assert"
	"context"
	"testing"
)

func TestNewClient(t *testing.T) {
	client := NewClient()
	defer client.Close()
	if client == nil {
		t.Fatal("client returned nil")
	}
}

func TestEtcdClient_PutValue(t *testing.T) {
	client := NewClient()
	defer client.Close()
	if client == nil {
		t.Fatal("client returned nil")
	}
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)

	_, err := client.PutValue(ctx, "test_key", "test_value")
	cancel()
	if err != nil {
		t.Error("Put value returned error", err)
	}
}

func TestEtcdClient_DeleteKey(t *testing.T) {
	client := NewClient()
	defer client.Close()
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	_, err := client.GetValue(ctx, "test_key")

	err = client.DeleteKey(ctx, "test_key")
	gr, err := client.GetValue(ctx, "test_key")
	fmt.Println(gr)
	if err != nil {
		t.Error("error in deleting key", err)
	}
	if gr.Kvs != nil {
		t.Error("key not deleted", gr)
	}
	cancel()
}

func TestEtcdClient_GetValue(t *testing.T) {
	client := NewClient()
	defer client.Close()
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	_, err := client.PutValue(ctx, "test_key2", "test_value2")
	if err != nil {
		t.Error("error in get value")
	}
	gr, err := client.GetValue(ctx, "test_key2")
	cancel()
	if err != nil {
		t.Error("error in get value", err)
	}
	if string(gr.Kvs[0].Value) != "test_value2" {
		t.Errorf("expected %s, returned %s", "test_value2", string(gr.Kvs[0].Value))
	}
}
func TestEtcdClient_GetValueWithRevision(t *testing.T) {
	client := NewClient()
	defer client.Close()
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)

	resp1, err := client.PutValue(ctx, "test_key2", "test_value2")
	if err != nil {
		t.Error("error in put value", err)
	}
	resp2, err := client.PutValue(ctx, "test_key2", "test_value3")
	if err != nil {
		t.Error("error in put value", err)
	}

	grv, err := client.GetValueWithRevision(ctx, "test_key2", resp1)
	if err != nil {
		t.Error("error in get value", err)
	}
	if string(grv.Kvs[0].Value) != "test_value2" {
		t.Errorf("expected %s, returned %s", "test_value2", string(grv.Kvs[0].Value))
	}

	grv, err = client.GetValueWithRevision(ctx, "test_key2", resp2)
	if err != nil {
		t.Error("error in get value", err)
	}
	if string(grv.Kvs[0].Value) != "test_value3" {
		t.Errorf("expected %s, returned %s", "test_value3", string(grv.Kvs[0].Value))
	}
	cancel()
}
