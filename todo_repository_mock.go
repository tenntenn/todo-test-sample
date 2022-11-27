// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package main

import (
	"sync"
)

// Ensure, that TodoRepositoryMock does implement TodoRepository.
// If this is not the case, regenerate this file with moq.
var _ TodoRepository = &TodoRepositoryMock{}

// TodoRepositoryMock is a mock implementation of TodoRepository.
//
//	func TestSomethingThatUsesTodoRepository(t *testing.T) {
//
//		// make and configure a mocked TodoRepository
//		mockedTodoRepository := &TodoRepositoryMock{
//			FindByIDFunc: func(id int) (*Todo, error) {
//				panic("mock out the FindByID method")
//			},
//		}
//
//		// use mockedTodoRepository in code that requires TodoRepository
//		// and then make assertions.
//
//	}
type TodoRepositoryMock struct {
	// FindByIDFunc mocks the FindByID method.
	FindByIDFunc func(id int) (*Todo, error)

	// calls tracks calls to the methods.
	calls struct {
		// FindByID holds details about calls to the FindByID method.
		FindByID []struct {
			// ID is the id argument value.
			ID int
		}
	}
	lockFindByID sync.RWMutex
}

// FindByID calls FindByIDFunc.
func (mock *TodoRepositoryMock) FindByID(id int) (*Todo, error) {
	if mock.FindByIDFunc == nil {
		panic("TodoRepositoryMock.FindByIDFunc: method is nil but TodoRepository.FindByID was just called")
	}
	callInfo := struct {
		ID int
	}{
		ID: id,
	}
	mock.lockFindByID.Lock()
	mock.calls.FindByID = append(mock.calls.FindByID, callInfo)
	mock.lockFindByID.Unlock()
	return mock.FindByIDFunc(id)
}

// FindByIDCalls gets all the calls that were made to FindByID.
// Check the length with:
//
//	len(mockedTodoRepository.FindByIDCalls())
func (mock *TodoRepositoryMock) FindByIDCalls() []struct {
	ID int
} {
	var calls []struct {
		ID int
	}
	mock.lockFindByID.RLock()
	calls = mock.calls.FindByID
	mock.lockFindByID.RUnlock()
	return calls
}
