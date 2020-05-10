package main

import (
	"github.com/patipolst/go-demo/pkg/http/rest"
	"github.com/patipolst/go-demo/pkg/mutation"
	"github.com/patipolst/go-demo/pkg/query"
	"github.com/patipolst/go-demo/pkg/store/db"
)

func main() {
	todoStore, _ := db.NewTodoStore()
	// todoStore := memory.NewTodoStore()
	todoQuery := query.NewTodoQuery(todoStore)
	todoMutation := mutation.NewTodoMutation(todoStore)
	// graphql.Run(todoQuery, todoMutation)
	rest.Run(todoQuery, todoMutation)
}
