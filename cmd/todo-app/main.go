package main

import (
	"github.com/patipolst/go-demo/pkg/http/rest"
	"github.com/patipolst/go-demo/pkg/service"
	"github.com/patipolst/go-demo/pkg/store/memory"
)

func main() {
	// todoStore, _ := db.NewTodoStore()
	todoStore := memory.NewTodoStore()
	// todoService := service.NewTodoDBService(todoStore)
	todoService := service.NewTodoMemoryService(todoStore)
	// graphql.Run(todoService)
	rest.Run(todoService)
}
