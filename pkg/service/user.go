package service

import (
	"github.com/patipolst/go-demo/pkg/mutation"
	"github.com/patipolst/go-demo/pkg/query"
	"github.com/patipolst/go-demo/pkg/store"
)

type User struct {
	Query    query.UserQuery
	Mutation mutation.UserMutation
}

func NewUser(s store.User) *User {
	q := query.NewUser(s)
	m := mutation.NewUser(s)
	return &User{q, m}
}
