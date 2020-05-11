package service

import (
	"github.com/patipolst/go-demo/pkg/mutation"
	"github.com/patipolst/go-demo/pkg/query"
	"github.com/patipolst/go-demo/pkg/store/database"
	"github.com/patipolst/go-demo/pkg/store/memory"
)

type TodoService struct {
	Query    query.TodoQuery
	Mutation mutation.TodoMutation
}

func NewTodoDBService(store *database.TodoStore) *TodoService {
	q := query.NewTodoQuery(store)
	m := mutation.NewTodoMutation(store)
	return &TodoService{q, m}
}

func NewTodoMemoryService(store *memory.TodoStore) *TodoService {
	q := query.NewTodoQuery(store)
	m := mutation.NewTodoMutation(store)
	return &TodoService{q, m}
}
