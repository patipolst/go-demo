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

type userStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *userStore {
	return &userStore{db}
}

func (s *userStore) GetAllUsers() []*query.User {
	var us []User
	var users []*query.User
	s.db.Find(&us)
	for _, u := range us {
		users = append(users, u.ToQueryUser())
	}
	return users
}

func (s *userStore) GetUser(id int) (*query.User, error) {
	var u User
	s.db.First(&u, id)
	return u.ToQueryUser(), nil
}

func (s *userStore) CreateUser(u mutation.NewUser) (*query.User, error) {
	newUser := User{
		Name: u.Name,
	}
	s.db.Create(&newUser)
	return newUser.ToQueryUser(), nil
}

func (s *userStore) DeleteUser(id int) {
	var u User
	s.db.First(&u, id)
	s.db.Delete(&u)
}
