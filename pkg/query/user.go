package query

type User struct {
	ID   int
	Name string
}

type UserQuery interface {
	GetUsers() []*User
	GetUser(int) (*User, error)
}

type UserStore interface {
	GetAllUsers() []*User
	GetUser(int) (*User, error)
}

type userQuery struct {
	store UserStore
}

func NewUserQuery(store UserStore) UserQuery {
	return &userQuery{store}
}

func (q *userQuery) GetUsers() []*User {
	return q.store.GetAllUsers()
}

func (q *userQuery) GetUser(id int) (*User, error) {
	return q.store.GetUser(id)
}
