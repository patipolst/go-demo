package rest

import (
	"log"

	"github.com/gofiber/fiber"
	"github.com/patipolst/go-demo/pkg/controller"
	"github.com/patipolst/go-demo/pkg/service"
)

func Run(s *service.TodoService) {
	app := fiber.New()
	c := controller.NewTodoController(s)

	v1 := app.Group("/v1")
	v1.Get("/todos", c.GetTodos)
	v1.Get("/todos/:id", c.GetTodo)
	v1.Post("/todos", c.CreateTodo)
	v1.Delete("/todos/:id", c.DeleteTodo)

	const port = 3000
	log.Printf("connect to http://localhost:%d/ for REST api", port)
	app.Listen(port)
}
