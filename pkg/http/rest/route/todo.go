package route

import (
	"github.com/gofiber/fiber"
	"github.com/patipolst/go-demo/pkg/controller"
)

func Todo(app *fiber.App, c *controller.TodoController) {
	v1 := app.Group("/v1")
	todoV1 := v1.Group("/todos")
	todoV1.Get("/", c.GetTodos)
	todoV1.Get("/:id", c.GetTodo)
	todoV1.Post("/", c.CreateTodo)
	todoV1.Delete("/:id", c.DeleteTodo)
}
