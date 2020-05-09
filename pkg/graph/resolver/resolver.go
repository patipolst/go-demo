package resolver

//go:generate go run github.com/99designs/gqlgen

import (
	"github.com/patipolst/go-demo/pkg/graph/model"
	"github.com/patipolst/go-demo/pkg/mutation"
	"github.com/patipolst/go-demo/pkg/query"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	AllTodos     []*model.Todo // fixme
	TodoQuery    query.TodoQuery
	TodoMutation mutation.TodoMutation
}
