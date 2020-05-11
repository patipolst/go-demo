package main

import (
	"github.com/patipolst/go-demo/pkg/http"
	"github.com/patipolst/go-demo/pkg/service"
	"github.com/patipolst/go-demo/pkg/store/database"
)

func main() {
	db := database.New()
	database.Migrate(db)
	defer db.Close()

	todoStore := database.NewTodoStore(db)
	userStore := database.NewUserStore(db)
	todoService := service.NewTodoDBService(todoStore)
	userService := service.NewUserDBService(userStore)
	// graphql.Run(todoService)
	http.Run(todoService, userService)
}
