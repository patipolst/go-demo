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

type TodoMutationStore interface {
	CreateTodo(NewTodo) (*query.Todo, error)
}

type todoMutation struct {
	store TodoMutationStore
}

func NewTodoMutation(store TodoMutationStore) TodoMutation {
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
