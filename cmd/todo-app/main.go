package main

import (
	"github.com/patipolst/go-demo/pkg/api"
	"github.com/patipolst/go-demo/pkg/service"
	"github.com/patipolst/go-demo/pkg/store/database"
)

func main() {
	db := database.New()
	database.Migrate(db)
	defer db.Close()

	todoStore := database.NewTodo(db)
	userStore := database.NewUser(db)
	todoService := service.NewTodo(todoStore)
	userService := service.NewUser(userStore)
	api.Run(todoService, userService)
}
