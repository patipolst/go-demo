package main

import (
	"github.com/patipolst/go-demo/pkg/http/graphql"
	"github.com/patipolst/go-demo/pkg/mutation"
	"github.com/patipolst/go-demo/pkg/query"
	"github.com/patipolst/go-demo/pkg/service"
	"github.com/patipolst/go-demo/pkg/store/db"
)

func main() {
	todoStore, _ := db.NewTodoStore()
	// todoStore := memory.NewTodoStore()
	todoQuery := query.NewTodoQuery(todoStore)
	todoMutation := mutation.NewTodoMutation(todoStore)
	todoService := service.NewTodoService(todoQuery, todoMutation)
	graphql.Run(todoService)
	// rest.Run(todoService)
}
