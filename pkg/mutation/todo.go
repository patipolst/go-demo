package mutation

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type TodoMutation interface {
	CreateTodo(...NewTodo)
	CreateTodos([]NewTodo)
}

type TodoStore interface {
	CreateTodo(NewTodo) error
}

type todoMutation struct {
	store TodoStore
}

func NewTodoMutation(store TodoStore) TodoMutation {
	return &todoMutation{store}
}

func (s *todoMutation) CreateTodo(todo ...NewTodo) {
	for _, t := range todo {
		_ = s.store.CreateTodo(t)
	}
}

func (s *todoMutation) CreateTodos(todos []NewTodo) {
	for _, t := range todos {
		_ = s.store.CreateTodo(t)
	}
}
