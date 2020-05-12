package store

import (
	"github.com/patipolst/go-demo/pkg/mutation"
	"github.com/patipolst/go-demo/pkg/query"
)

type TodoStore interface {
	GetAllTodos() []*query.Todo
	GetTodo(id int) (*query.Todo, error)
	CreateTodo(t mutation.NewTodo) (*query.Todo, error)
	DeleteTodo(id int)
}
