package query

type Todo struct {
	ID     string `json:"id"`
	Text   string `json:"text"`
	Done   bool   `json:"done"`
	UserID string `json:"user"`
}

type TodoQuery interface {
	GetTodos() []*Todo
	GetTodo(string) (*Todo, error)
}

type TodoStore interface {
	GetAllTodos() []*Todo
	GetTodo(string) (*Todo, error)
}

type todoQuery struct {
	store TodoStore
}

func NewTodoQuery(store TodoStore) TodoQuery {
	return &todoQuery{store}
}

func (s *todoQuery) GetTodos() []*Todo {
	return s.store.GetAllTodos()
}

func (s *todoQuery) GetTodo(id string) (*Todo, error) {
	return s.store.GetTodo(id)
}
