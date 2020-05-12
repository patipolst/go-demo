package query

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UserQuery interface {
	GetUsers() []*User
	GetUser(int) (*User, error)
}

type UserQueryStore interface {
	GetAllUsers() []*User
	GetUser(int) (*User, error)
}

type userQuery struct {
	store UserQueryStore
}

func NewUser(store UserQueryStore) UserQuery {
	return &userQuery{store}
}

func (q *userQuery) GetUsers() []*User {
	return q.store.GetAllUsers()
}

func (q *userQuery) GetUser(id int) (*User, error) {
	return q.store.GetUser(id)
}
