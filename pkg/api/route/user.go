package route

import (
	"github.com/gofiber/fiber"
	"github.com/patipolst/go-demo/pkg/api/controller"
)

func User(app *fiber.App, c *controller.User) {
	v1 := app.Group("/v1")
	userV1 := v1.Group("/users")
	userV1.Get("/", c.GetUsers)
	userV1.Get("/:id", c.GetUser)
	userV1.Post("/", c.CreateUser)
	userV1.Delete("/:id", c.DeleteUser)
}
