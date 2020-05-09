package memory

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/patipolst/go-demo/pkg/mutation"
	"github.com/patipolst/go-demo/pkg/query"
)

type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

type TodoStore struct {
	todos map[string]Todo
}

func NewTodoStore() *TodoStore {
	return &TodoStore{
		make(map[string]Todo),
	}
}

func (s *TodoStore) GetAllTodos() []query.Todo {
	var todos []query.Todo
	for i := range s.todos {
		todo := query.Todo{
			ID:   s.todos[i].ID,
			Text: s.todos[i].Text,
			Done: s.todos[i].Done,
		}
		todos = append(todos, todo)
	}
	return todos
}

func (s *TodoStore) GetTodo(id string) (query.Todo, error) {
	var todo query.Todo

	for i := range s.todos {
		if s.todos[i].ID == id {
			todo.ID = s.todos[i].ID
			todo.Text = s.todos[i].Text
			todo.Done = s.todos[i].Done

			return todo, nil
		}
	}
	return todo, errors.New("todo not found")
}

func (s *TodoStore) CreateTodo(t mutation.NewTodo) error {
	id := fmt.Sprintf("T%d", rand.Int())
	newTodo := Todo{
		ID:   id,
		Text: t.Text,
		Done: false,
	}
	s.todos[id] = newTodo
	return nil
}

func (s *TodoStore) DeleteTodo(id string) {
	delete(s.todos, id)
}
