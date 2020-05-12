package mutation

import "github.com/patipolst/go-demo/pkg/query"

type Todo struct {
	Text   string `json:"text"`
	UserID int    `json:"userId"`
}

type TodoMutation interface {
	CreateTodo(Todo) *query.Todo
	CreateTodos([]Todo) []*query.Todo
}

type TodoMutationStore interface {
	CreateTodo(Todo) (*query.Todo, error)
}

type todoMutation struct {
	store TodoMutationStore
}

func NewTodo(store TodoMutationStore) TodoMutation {
	return &todoMutation{store}
}

func (m *todoMutation) CreateTodo(todo Todo) *query.Todo {
	t, _ := m.store.CreateTodo(todo)
	return t
}

func (m *todoMutation) CreateTodos(todos []Todo) []*query.Todo {
	var created []*query.Todo
	for _, todo := range todos {
		t := m.CreateTodo(todo)
		created = append(created, t)
	}
	return created
}
