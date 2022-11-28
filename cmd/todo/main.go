package main

import (
	"context"
	"fmt"

	"github.com/keyamin/todo-test-sample/internal"
	"github.com/keyamin/todo-test-sample/internal/samplerepo"
)

func main() {
	list := internal.TodoList{
		Repo: samplerepo.SampleRepo{},
	}
	todo, _ := list.Get(context.Background(), 1)

	fmt.Println(todo.ID)
	fmt.Println(todo.Title)
	fmt.Println(todo.Completed)
}
