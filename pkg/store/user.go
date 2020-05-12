package store

import (
	"github.com/patipolst/go-demo/pkg/mutation"
	"github.com/patipolst/go-demo/pkg/query"
)

type UserStore interface {
	GetAllUsers() []*query.User
	GetUser(id int) (*query.User, error)
	CreateUser(u mutation.User) (*query.User, error)
	DeleteUser(id int)
}
