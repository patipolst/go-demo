package mutation

import "github.com/patipolst/go-demo/pkg/query"

type NewTodo struct {
	Text   string
	UserID int
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

func (m *todoMutation) CreateTodo(todo NewTodo) *query.Todo {
	t, _ := m.store.CreateTodo(todo)
	return t
}

func (m *todoMutation) CreateTodos(todos []NewTodo) []*query.Todo {
	var created []*query.Todo
	for _, todo := range todos {
		t := m.CreateTodo(todo)
		created = append(created, t)
	}
	return created
}
