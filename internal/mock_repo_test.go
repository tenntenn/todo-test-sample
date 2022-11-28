// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package internal_test

import (
	"context"
	"sync"

	"github.com/keyamin/todo-test-sample/internal"
)

// Ensure, that RepositoryMock does implement internal.Repository.
// If this is not the case, regenerate this file with moq.
var _ internal.Repository = &RepositoryMock{}

// RepositoryMock is a mock implementation of internal.Repository.
//
//	func TestSomethingThatUsesRepository(t *testing.T) {
//
//		// make and configure a mocked internal.Repository
//		mockedRepository := &RepositoryMock{
//			GetTodoByIDFunc: func(ctx context.Context, id int) (*internal.Todo, error) {
//				panic("mock out the GetTodoByID method")
//			},
//		}
//
//		// use mockedRepository in code that requires internal.Repository
//		// and then make assertions.
//
//	}
type RepositoryMock struct {
	// GetTodoByIDFunc mocks the GetTodoByID method.
	GetTodoByIDFunc func(ctx context.Context, id int) (*internal.Todo, error)

	// calls tracks calls to the methods.
	calls struct {
		// GetTodoByID holds details about calls to the GetTodoByID method.
		GetTodoByID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID int
		}
	}
	lockGetTodoByID sync.RWMutex
}

// GetTodoByID calls GetTodoByIDFunc.
func (mock *RepositoryMock) GetTodoByID(ctx context.Context, id int) (*internal.Todo, error) {
	callInfo := struct {
		Ctx context.Context
		ID  int
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockGetTodoByID.Lock()
	mock.calls.GetTodoByID = append(mock.calls.GetTodoByID, callInfo)
	mock.lockGetTodoByID.Unlock()
	if mock.GetTodoByIDFunc == nil {
		var (
			todoOut *internal.Todo
			errOut  error
		)
		return todoOut, errOut
	}
	return mock.GetTodoByIDFunc(ctx, id)
}

// GetTodoByIDCalls gets all the calls that were made to GetTodoByID.
// Check the length with:
//
//	len(mockedRepository.GetTodoByIDCalls())
func (mock *RepositoryMock) GetTodoByIDCalls() []struct {
	Ctx context.Context
	ID  int
} {
	var calls []struct {
		Ctx context.Context
		ID  int
	}
	mock.lockGetTodoByID.RLock()
	calls = mock.calls.GetTodoByID
	mock.lockGetTodoByID.RUnlock()
	return calls
}
