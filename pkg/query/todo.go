package query

type Todo struct {
	ID     int    `json:"id"`
	Text   string `json:"text"`
	Done   bool   `json:"done"`
	UserID int    `json:"user"`
}

type TodoQuery interface {
	GetTodos() []*Todo
	GetTodo(int) (*Todo, error)
}

type TodoStore interface {
	GetAllTodos() []*Todo
	GetTodo(int) (*Todo, error)
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

func (s *todoQuery) GetTodo(id int) (*Todo, error) {
	return s.store.GetTodo(id)
}
