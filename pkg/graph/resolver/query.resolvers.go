package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/patipolst/go-demo/pkg/graph/generated"
	"github.com/patipolst/go-demo/pkg/query"
)

func (r *queryResolver) Todos(ctx context.Context) ([]*query.Todo, error) {
	return r.TodoService.Query.GetTodos(), nil
}

func (r *queryResolver) Todo(ctx context.Context, id int) (*query.Todo, error) {
	return r.TodoService.Query.GetTodo(id)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
