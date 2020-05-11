package database

import (
	"github.com/jinzhu/gorm"
	"github.com/patipolst/go-demo/pkg/mutation"
	"github.com/patipolst/go-demo/pkg/query"
)

type User struct {
	gorm.Model
	Name string
}

func (u *User) ToQueryUser() *query.User {
	return &query.User{
		ID:   int(u.ID),
		Name: u.Name,
	}
}

type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{db}
}

func (s *UserStore) GetAllUsers() []*query.User {
	var us []User
	var users []*query.User
	s.db.Find(&us)
	for _, u := range us {
		users = append(users, u.ToQueryUser())
	}
	return users
}

func (s *UserStore) GetUser(id int) (*query.User, error) {
	var u User
	s.db.First(&u, id)
	return u.ToQueryUser(), nil
}

func (s *UserStore) CreateUser(u mutation.NewUser) (*query.User, error) {
	newUser := User{
		Name: u.Name,
	}
	s.db.Create(&newUser)
	return newUser.ToQueryUser(), nil
}

func (s *UserStore) DeleteUser(id int) {
	var u User
	s.db.First(&u, id)
	s.db.Delete(&u)
}
