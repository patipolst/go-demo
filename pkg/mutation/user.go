package mutation

import "github.com/patipolst/go-demo/pkg/query"

type NewUser struct {
	Name string `json:"name"`
}

type UserMutation interface {
	CreateUser(NewUser) *query.User
	CreateUsers([]NewUser) []*query.User
}

type UserStore interface {
	CreateUser(NewUser) (*query.User, error)
}

type userMutation struct {
	store UserStore
}

func NewUserMutation(store UserStore) UserMutation {
	return &userMutation{store}
}

func (m *userMutation) CreateUser(user NewUser) *query.User {
	u, _ := m.store.CreateUser(user)
	return u
}

func (m *userMutation) CreateUsers(users []NewUser) []*query.User {
	var created []*query.User
	for _, user := range users {
		u := m.CreateUser(user)
		created = append(created, u)
	}
	return created
}
