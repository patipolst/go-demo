package mutation

import "github.com/patipolst/go-demo/pkg/query"

type NewTodo struct {
	Text   string `json:"text"`
	UserID int    `json:"userId"`
}

type TodoMutation interface {
	CreateTodo(NewTodo) *query.Todo
	CreateTodos([]NewTodo) []*query.Todo
}

type TodoStore interface {
	CreateTodo(NewTodo) (*query.Todo, error)
}

type todoMutation struct {
	store TodoStore
}

func NewTodoMutation(store TodoStore) TodoMutation {
	return &todoMutation{store}
}

func (s *todoMutation) CreateTodo(todo NewTodo) *query.Todo {
	t, _ := s.store.CreateTodo(todo)
	return t
}

func (s *todoMutation) CreateTodos(todos []NewTodo) []*query.Todo {
	var created []*query.Todo
	for _, todo := range todos {
		t := s.CreateTodo(todo)
		created = append(created, t)
	}
	return created
}
