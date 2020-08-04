package etcd

import (
	"context"
	"testing"
	"crud-toy/internal/model"
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
	proc := model.Proc{
		ID:     "1",
		Name:   "New Name",
		Author: "New Author",
	}
	_, err := client.PutValue(ctx, "test_key",&proc)
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
	proc, err := client.GetValue(ctx, "test_key")

	if err == nil {
		t.Error("value still being retirieved")
	}
	if proc != nil {
		t.Error("key not deleted", proc)
	}
	cancel()
}

func TestEtcdClient_GetValue(t *testing.T) {
	client := NewClient()
	defer client.Close()
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	proc := model.Proc{
		ID:     "1",
		Name:   "New Name",
		Author: "New Author",
	}
	_, err := client.PutValue(ctx, proc.ID, &proc)
	if err != nil {
		t.Error("error in get value")
	}
	res, err := client.GetValue(ctx, proc.ID)
	cancel()
	if err != nil {
		t.Error("error in get value", err)
	}
	if res.ID != proc.ID {
		t.Errorf("expected %s, returned %s", "test_value2", res.ID)
	}
}
func TestEtcdClient_GetValueWithRevision(t *testing.T) {
	client := NewClient()
	defer client.Close()
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	proc1 := model.Proc{
		ID:     "1",
		Name:   "New Name",
		Author: "New Author",
	}
	_, err := client.PutValue(ctx, proc1.ID, &proc1)
	if err != nil {
		t.Error("error in put value", err)
	}
	proc2 := model.Proc{
		ID:     "2",
		Name:   "New Name2",
		Author: "New Author2",
	}
	_, err = client.PutValue(ctx, proc2.ID ,&proc2)
	if err != nil {
		t.Error("error in put value", err)
	}
	
	header,err := client.GetProcRevisionById(ctx, proc1.ID)
	if err != nil {
		t.Error("error in getting revision number", err)
	}
	grv, err := client.GetValueWithRevision(ctx, proc1.ID, header)
	if err != nil {
		t.Error("error in get value", err)
	}
	if grv.ID != "1" {
		t.Errorf("expected %s, returned %s", "1", grv.ID)
	}

	header,err = client.GetProcRevisionById(ctx, proc2.ID)
	if err != nil {
		t.Error("error in getting revision number", err)
	}
	grv, err = client.GetValueWithRevision(ctx,  proc2.ID, header)
	if err != nil {
		t.Error("error in get value", err)
	}
	if grv.ID!= "2" {
		t.Errorf("expected %s, returned %s", "2", grv.ID)
	}
	cancel()
}
