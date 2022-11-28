package internal_test

import (
	"context"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/keyamin/todo-test-sample/internal"
)

//go:generate moq -stub -pkg internal_test -fmt goimports -out mock_repo_test.go . Repository

func TestTodoList_Get(t *testing.T) {
	t.Parallel()

	const exist, notfound = 100, 200
	database := map[int]*internal.Todo{
		exist: &internal.Todo{ID: exist, Title: "todo-1"},
	}

	tests := map[string]struct {
		id      int
		want    *internal.Todo
		wantErr bool
	}{
		"exist":     {exist, database[exist], false},
		"not found": {notfound, nil, true},
	}

	for name, tt := range tests {
		name, tt := name, tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			repo := &RepositoryMock{
				GetTodoByIDFunc: func(_ context.Context, id int) (*internal.Todo, error) {
					todo, ok := database[id]
					if ok {
						return todo, nil
					}
					return nil, errors.New("not found")
				},
			}

			list := &internal.TodoList{Repo: repo}
			got, err := list.Get(context.Background(), tt.id)

			switch {
			case tt.wantErr && err == nil:
				t.Error("expected error did not occur")
			case !tt.wantErr && err != nil:
				t.Fatal("unexpected error:", err)
			}

			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Error(diff)
			}
		})
	}
}
