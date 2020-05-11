package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
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

func NewUserStore() (*UserStore, error) {
	var err error
	s := new(UserStore)

	s.db, err = gorm.Open("sqlite3", "users.db")
	if err != nil {
		return nil, err
	}

	s.db.AutoMigrate(&User{})

	return s, nil
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
