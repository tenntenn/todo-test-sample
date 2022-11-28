package samplerepo

import (
	"context"

	"github.com/keyamin/todo-test-sample/internal"
)

type SampleRepo struct{}

var _ internal.Repository = SampleRepo{}

func (repo SampleRepo) GetTodoByID(ctx context.Context, id int) (*internal.Todo, error) {
	return &internal.Todo{
		ID:        id,
		Title:     "title",
		Completed: false,
	}, nil
}
