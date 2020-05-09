package mutation

import "github.com/patipolst/go-demo/pkg/query"

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type TodoMutation interface {
	CreateTodo(NewTodo) *query.Todo
	// CreateTodo(...NewTodo) // fixme
	CreateTodos([]NewTodo)
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

// func (s *todoMutation) CreateTodo(todo ...NewTodo) {
// 	for _, t := range todo {
// 		_ = s.store.CreateTodo(t)
// 	}
// }

func (s *todoMutation) CreateTodos(todos []NewTodo) {
	for _, t := range todos {
		_, _ = s.store.CreateTodo(t)
	}
}
