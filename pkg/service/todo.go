package service

import (
	"github.com/patipolst/go-demo/pkg/mutation"
	"github.com/patipolst/go-demo/pkg/query"
	"github.com/patipolst/go-demo/pkg/store"
)

type TodoService struct {
	Query    query.TodoQuery
	Mutation mutation.TodoMutation
}

func NewTodoService(s store.TodoStore) *TodoService {
	q := query.NewTodoQuery(s)
	m := mutation.NewTodoMutation(s)
	return &TodoService{q, m}
}
