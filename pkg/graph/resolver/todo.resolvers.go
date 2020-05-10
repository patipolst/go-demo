package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/patipolst/go-demo/pkg/graph/generated"
	"github.com/patipolst/go-demo/pkg/query"
)

func (r *todoResolver) User(ctx context.Context, obj *query.Todo) (*query.User, error) {
	return &query.User{ID: obj.UserID, Name: fmt.Sprintf("User %d", obj.UserID)}, nil
}

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type todoResolver struct{ *Resolver }
