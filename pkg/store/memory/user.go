package memory

import (
	"errors"

	"github.com/patipolst/go-demo/pkg/mutation"
	"github.com/patipolst/go-demo/pkg/query"
)

type User struct {
	ID   int
	Name string
}

func (u *User) ToQueryUser() *query.User {
	return &query.User{
		ID:   u.ID,
		Name: u.Name,
	}
}

type UserStore struct {
	users map[int]User
}

func NewUserStore() *UserStore {
	return &UserStore{
		make(map[int]User),
	}
}

func (s *UserStore) GetAllUsers() []*query.User {
	var users []*query.User
	for i := range s.users {
		user := query.User{
			ID:   s.users[i].ID,
			Name: s.users[i].Name,
		}
		users = append(users, &user)
	}
	return users
}

func (s *UserStore) GetUser(id int) (*query.User, error) {
	var user query.User

	for i := range s.users {
		if s.users[i].ID == id {
			user.ID = s.users[i].ID
			user.Name = s.users[i].Name

			return &user, nil
		}
	}
	return &user, errors.New("user not found")
}

func (s *UserStore) CreateUser(u mutation.NewUser) (*query.User, error) {
	id := len(s.users) + 1
	newUser := User{
		ID:   id,
		Name: u.Name,
	}
	s.users[id] = newUser
	return newUser.ToQueryUser(), nil
}

func (s *UserStore) DeleteUser(id int) {
	delete(s.users, id)
}
