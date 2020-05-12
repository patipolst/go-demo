package mutation

import "github.com/patipolst/go-demo/pkg/query"

type User struct {
	Name string `json:"name"`
}

type UserMutation interface {
	CreateUser(User) *query.User
	CreateUsers([]User) []*query.User
}

type userMutationStore interface {
	CreateUser(User) (*query.User, error)
}

type userMutation struct {
	store userMutationStore
}

func NewUser(store userMutationStore) UserMutation {
	return &userMutation{store}
}

func (m *userMutation) CreateUser(user User) *query.User {
	u, _ := m.store.CreateUser(user)
	return u
}

func (m *userMutation) CreateUsers(users []User) []*query.User {
	var created []*query.User
	for _, user := range users {
		u := m.CreateUser(user)
		created = append(created, u)
	}
	return created
}
