package service

import (
	"github.com/patipolst/go-demo/pkg/mutation"
	"github.com/patipolst/go-demo/pkg/query"
)

type TodoService struct {
	Query query.TodoQuery
	Mutation mutation.TodoMutation
}

func NewTodoService(q query.TodoQuery, m mutation.TodoMutation) *TodoService {
	return &TodoService{q, m}
}