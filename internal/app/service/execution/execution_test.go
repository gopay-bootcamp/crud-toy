package execution

import (
	"context"
	"crud-toy/internal/app/model"
	"crud-toy/internal/app/service/infra/db/etcd"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateProc(t *testing.T) {
	ctx := context.Background()
	mockClient := new(etcd.ClientMock)
	testExec := NewExec(mockClient)
	proc := model.Proc{
		ID:     "1",
		Name:   "New Name",
		Author: "New Author",
	}
	mockClient.On("PutValue").Return(ctx, &proc, nil)

	result, _ := testExec.CreateProc(&proc)

	mockClient.AssertExpectations(t)
	//data assertions
	assert.NotNil(t, result.ID)
	assert.Equal(t, "New Name", result.Name)
	assert.Equal(t, "New Author", result.Author)
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

	result, _ := testExec.ReadAllProc()

	mockClient.AssertExpectations(t)
	//data assertions
	assert.NotNil(t, result[0].ID)
	assert.Equal(t, "New Name", result[0].Name)
	assert.Equal(t, "New Author", result[0].Author)
}

func TestReadProcByID(t *testing.T) {
	mockClient := new(etcd.ClientMock)
	testExec := NewExec(mockClient)
	proc := model.Proc{
		ID:     "1",
		Name:   "New Name",
		Author: "New Author",
	}
	mockClient.On("GetValue").Return(&proc, nil)

	result, _ := testExec.ReadProcByID(&model.Proc{ID: "1"})
	mockClient.AssertExpectations(t)
	//data assertions
	assert.NotNil(t, result.ID)
	assert.Equal(t, "New Name", result.Name)
	assert.Equal(t, "New Author", result.Author)
}

func TestUpdateProc(t *testing.T) {
	ctx := context.Background()
	mockClient := new(etcd.ClientMock)
	testExec := NewExec(mockClient)
	proc := model.Proc{
		ID:     "1",
		Name:   "Changed Name",
		Author: "Changed Author",
	}
	mockClient.On("PutValue").Return(ctx, &proc, nil)

	result, _ := testExec.UpdateProc(&model.Proc{
		ID:     "1",
		Name:   "Name",
		Author: "Author",
	})

	mockClient.AssertExpectations(t)
	//data assertions
	assert.NotNil(t, result.ID)
	assert.Equal(t, "Changed Name", result.Name)
	assert.Equal(t, "Changed Author", result.Author)
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
