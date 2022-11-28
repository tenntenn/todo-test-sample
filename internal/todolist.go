package internal

import "context"

type TodoList struct {
	Repo Repository
}

func (list *TodoList) Get(ctx context.Context, id int) (*Todo, error) {
	return list.Repo.GetTodoByID(ctx, id)
}
