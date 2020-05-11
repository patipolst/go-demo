package resolver

//go:generate go run github.com/99designs/gqlgen

import (
	"github.com/patipolst/go-demo/pkg/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	TodoService *service.TodoService
	UserService *service.UserService
}
