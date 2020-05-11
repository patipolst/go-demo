package query

type Todo struct {
	ID     int
	Text   string
	Done   bool
	UserID int
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

func (q *todoQuery) GetTodos() []*Todo {
	return q.store.GetAllTodos()
}

func (q *todoQuery) GetTodo(id int) (*Todo, error) {
	return q.store.GetTodo(id)
}
