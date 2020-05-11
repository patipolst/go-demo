package service

import (
	"github.com/patipolst/go-demo/pkg/mutation"
	"github.com/patipolst/go-demo/pkg/query"
	"github.com/patipolst/go-demo/pkg/store/db"
	"github.com/patipolst/go-demo/pkg/store/memory"
)

type UserService struct {
	Query    query.UserQuery
	Mutation mutation.UserMutation
}

func NewUserDBService(store *db.UserStore) *UserService {
	q := query.NewUserQuery(store)
	m := mutation.NewUserMutation(store)
	return &UserService{q, m}
}

func NewUserMemoryService(store *memory.UserStore) *UserService {
	q := query.NewUserQuery(store)
	m := mutation.NewUserMutation(store)
	return &UserService{q, m}
}
