package execution

import (
	"crud-toy/internal/model"
	"crud-toy/internal/server/db/etcd"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateProc(t *testing.T) {
	mockClient := new(etcd.ClientMock)
	testExec := NewExec(mockClient)
	proc := model.Proc{
		ID:     "1",
		Name:   "New Name",
		Author: "New Author",
	}
	mockClient.On("PutValue", &proc).Return("1",nil)

	id,err := testExec.CreateProc(&proc)

	if err != nil{
		assert.Error(t,err)
	}
	mockClient.AssertExpectations(t)
	assert.Equal(t, "1", id)
}

func TestReadAllProcs(t *testing.T) {
	mockClient := new(etcd.ClientMock)
	testExec := NewExec(mockClient)
	proc := model.Proc{
		ID:     "1",
		Name:   "New Name",
		Author: "New Author",
	}
	mockClient.On("GetAllValues").Return([]model.Proc{proc}, nil)

	result, err := testExec.ReadAllProc()
	if err != nil{
		assert.Error(t,err)
	}
	mockClient.AssertExpectations(t)
	assert.NotEqual(t,len(result),0)
}

func TestReadProcByID(t *testing.T) {
	mockClient := new(etcd.ClientMock)
	testExec := NewExec(mockClient)
	proc := model.Proc{
		ID:     "1",
		Name:   "New Name",
		Author: "New Author",
	}
	mockClient.On("GetValue",proc.ID).Return(&proc, nil)

	result, err:= testExec.ReadProcByID(&model.Proc{ID: "1"})
	if err != nil{
		assert.Error(t,err)
	}
	mockClient.AssertExpectations(t)
	assert.Equal(t,result.ID,"1")
}

func TestUpdateProc(t *testing.T) {
	mockClient := new(etcd.ClientMock)
	testExec := NewExec(mockClient)
	proc := model.Proc{
		ID:     "1",
		Name:   "Changed Name",
		Author: "Changed Author",
	}
	mockClient.On("PutValue", &proc).Return(proc.ID, nil)

	_, err := testExec.UpdateProc(&proc)
	if err != nil{
		assert.Error(t,err)
	}
	mockClient.AssertExpectations(t)
}

func TestDeleteProc(t *testing.T) {
	mockClient := new(etcd.ClientMock)
	testExec := NewExec(mockClient)

	mockClient.On("DeleteKey").Return(nil)

	err := testExec.DeleteProc(&model.Proc{ID: "1"})
	mockClient.AssertExpectations(t)
	//data assertions
	assert.Nil(t, err, nil)
}
