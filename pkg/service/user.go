package service

import (
	"github.com/patipolst/go-demo/pkg/mutation"
	"github.com/patipolst/go-demo/pkg/query"
	"github.com/patipolst/go-demo/pkg/store"
)

type UserService struct {
	Query    query.UserQuery
	Mutation mutation.UserMutation
}

func NewUserService(s store.UserStore) *UserService {
	q := query.NewUserQuery(s)
	m := mutation.NewUserMutation(s)
	return &UserService{q, m}
}
