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

	todoStore := database.NewTodoStore(db)
	userStore := database.NewUserStore(db)
	// todoStore := memory.NewTodoStore()
	// userStore := memory.NewUserStore()
	todoService := service.NewTodoService(todoStore)
	userService := service.NewUserService(userStore)
	api.Run(todoService, userService)
}
