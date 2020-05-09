package mutation

type Todo struct {
	Text string `json:"text"`
	Done bool   `json:"done"`
}

type TodoMutation interface {
	CreateTodo(...Todo)
	CreateTodos([]Todo)
}

type TodoStore interface {
	CreateTodo(Todo) error
}

type todoMutation struct {
	store TodoStore
}

func NewTodoMutation(store TodoStore) TodoMutation {
	return &todoMutation{store}
}

func (s *todoMutation) CreateTodo(todo ...Todo) {
	for _, t := range todo {
		_ = s.store.CreateTodo(t)
	}
}

func (s *todoMutation) CreateTodos(todos []Todo) {
	for _, t := range todos {
		_ = s.store.CreateTodo(t)
	}
}
