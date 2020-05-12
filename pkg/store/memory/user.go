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

type userStore struct {
	users map[int]User
}

func NewUser() *userStore {
	return &userStore{
		make(map[int]User),
	}
}

func (s *userStore) GetAllUsers() []*query.User {
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

func (s *userStore) GetUser(id int) (*query.User, error) {
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

func (s *userStore) CreateUser(u mutation.user) (*query.User, error) {
	id := len(s.users) + 1
	newUser := User{
		ID:   id,
		Name: u.Name,
	}
	s.users[id] = newUser
	return newUser.ToQueryUser(), nil
}

func (s *userStore) DeleteUser(id int) {
	delete(s.users, id)
}
