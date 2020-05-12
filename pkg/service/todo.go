package service

import (
	"github.com/patipolst/go-demo/pkg/mutation"
	"github.com/patipolst/go-demo/pkg/query"
	"github.com/patipolst/go-demo/pkg/store"
)

type Todo struct {
	Query    query.TodoQuery
	Mutation mutation.TodoMutation
}

func NewTodo(s store.Todo) *Todo {
	q := query.NewTodo(s)
	m := mutation.NewTodo(s)
	return &Todo{q, m}
}
