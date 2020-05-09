package resolver

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/patipolst/go-demo/pkg/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		Text:   input.Text,
		ID:     fmt.Sprintf("T%d", rand.Int()),
		UserID: input.UserID,
	}
	r.AllTodos = append(r.AllTodos, todo)
	return todo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.AllTodos, nil
}

// func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
// 	todo := &model.Todo{
// 		Text:   input.Text,
// 		ID:     fmt.Sprintf("T%d", rand.Int()),
// 		UserID: input.UserID,
// 	}
// 	// r.todos = append(r.todos, todo)
// 	// r.TodoMutation.CreateTodo()
// 	return todo, nil
// }

// func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
// 	todos := r.TodoQuery.GetTodos()
// 	d := make([]*model.Todo, 0) // fixme
// 	return d, nil
// }
