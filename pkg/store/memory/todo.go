package memory

import (
	"errors"

	"github.com/patipolst/go-demo/pkg/mutation"
	"github.com/patipolst/go-demo/pkg/query"
)

type Todo struct {
	ID     int
	Text   string
	Done   bool
	UserID int
}

func (t *Todo) ToQueryTodo() *query.Todo {
	return &query.Todo{
		ID:     t.ID,
		Text:   t.Text,
		Done:   t.Done,
		UserID: t.UserID,
	}
}

type todoStore struct {
	todos map[int]Todo
}

func NewTodoStore() *todoStore {
	return &todoStore{
		make(map[int]Todo),
	}
}

func (s *todoStore) GetAllTodos() []*query.Todo {
	var todos []*query.Todo
	for i := range s.todos {
		todo := query.Todo{
			ID:     s.todos[i].ID,
			Text:   s.todos[i].Text,
			Done:   s.todos[i].Done,
			UserID: s.todos[i].UserID,
		}
		todos = append(todos, &todo)
	}
	return todos
}

func (s *todoStore) GetTodo(id int) (*query.Todo, error) {
	var todo query.Todo

	for i := range s.todos {
		if s.todos[i].ID == id {
			todo.ID = s.todos[i].ID
			todo.Text = s.todos[i].Text
			todo.Done = s.todos[i].Done
			todo.UserID = s.todos[i].UserID

			return &todo, nil
		}
	}
	return &todo, errors.New("todo not found")
}

func (s *todoStore) CreateTodo(t mutation.NewTodo) (*query.Todo, error) {
	id := len(s.todos) + 1
	newTodo := Todo{
		ID:     id,
		Text:   t.Text,
		Done:   false,
		UserID: t.UserID,
	}
	s.todos[id] = newTodo
	return newTodo.ToQueryTodo(), nil
}

func (s *todoStore) DeleteTodo(id int) {
	delete(s.todos, id)
}
