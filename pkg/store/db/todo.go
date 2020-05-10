package db

import (
	"fmt"
	"math/rand"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/patipolst/go-demo/pkg/mutation"
	"github.com/patipolst/go-demo/pkg/query"
)

type Todo struct {
	ID     string
	Text   string
	Done   bool
	UserID string
}

func (t *Todo) ToQueryTodo() *query.Todo {
	return &query.Todo{
		ID:     t.ID,
		Text:   t.Text,
		Done:   t.Done,
		UserID: t.UserID,
	}
}

type TodoStore struct {
	db *gorm.DB
}

func NewTodoStore() (*TodoStore, error) {
	var err error
	s := new(TodoStore)

	s.db, err = gorm.Open("sqlite3", "todos.db")
	if err != nil {
		return nil, err
	}

	s.db.AutoMigrate(&Todo{})

	return s, nil
}

func (s *TodoStore) GetAllTodos() []*query.Todo {
	var ts []Todo
	var todos []*query.Todo
	s.db.Find(&ts)
	for _, t := range ts {
		todos = append(todos, t.ToQueryTodo())
	}
	return todos
}

func (s *TodoStore) GetTodo(id string) (*query.Todo, error) {
	var t Todo
	s.db.First(&t, id)
	return t.ToQueryTodo(), nil
}

func (s *TodoStore) CreateTodo(t mutation.NewTodo) (*query.Todo, error) {
	newTodo := Todo{
		ID:     fmt.Sprintf("T%d", rand.Int()),
		Text:   t.Text,
		Done:   false,
		UserID: t.UserID,
	}
	s.db.Create(&newTodo)
	return newTodo.ToQueryTodo(), nil
}

func (s *TodoStore) DeleteTodo(id string) {
	var t Todo
	s.db.First(&t, id)
	s.db.Delete(&t)
}
