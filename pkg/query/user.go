package query

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UserQuery interface {
	GetUsers() []*User
	GetUser(int) (*User, error)
}

type userQueryStore interface {
	GetAllUsers() []*User
	GetUser(int) (*User, error)
}

type userQuery struct {
	store userQueryStore
}

func NewUser(store userQueryStore) UserQuery {
	return &userQuery{store}
}

func (q *userQuery) GetUsers() []*User {
	return q.store.GetAllUsers()
}

func (q *userQuery) GetUser(id int) (*User, error) {
	return q.store.GetUser(id)
}
