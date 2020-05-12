package database

import (
	"github.com/jinzhu/gorm"
	"github.com/patipolst/go-demo/pkg/mutation"
	"github.com/patipolst/go-demo/pkg/query"
)

type Todo struct {
	gorm.Model
	Text   string
	Done   bool
	UserID int
}

func (t *Todo) ToQueryTodo() *query.Todo {
	return &query.Todo{
		ID:     int(t.ID),
		Text:   t.Text,
		Done:   t.Done,
		UserID: t.UserID,
	}
}

type todoStore struct {
	db *gorm.DB
}

func NewTodoStore(db *gorm.DB) *todoStore {
	return &todoStore{db}
}

func (s *todoStore) GetAllTodos() []*query.Todo {
	var ts []Todo
	var todos []*query.Todo
	s.db.Find(&ts)
	for _, t := range ts {
		todos = append(todos, t.ToQueryTodo())
	}
	return todos
}

func (s *todoStore) GetTodo(id int) (*query.Todo, error) {
	var t Todo
	s.db.First(&t, id)
	return t.ToQueryTodo(), nil
}

func (s *todoStore) CreateTodo(t mutation.NewTodo) (*query.Todo, error) {
	newTodo := Todo{
		Text:   t.Text,
		Done:   false,
		UserID: t.UserID,
	}
	s.db.Create(&newTodo)
	return newTodo.ToQueryTodo(), nil
}

func (s *todoStore) DeleteTodo(id int) {
	var t Todo
	s.db.First(&t, id)
	s.db.Delete(&t)
}
