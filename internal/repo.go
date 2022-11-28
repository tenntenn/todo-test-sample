package internal

import "context"

type Repository interface {
	GetTodoByID(ctx context.Context, id int) (*Todo, error)
}
