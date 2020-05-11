package api

import (
	"log"

	"github.com/gofiber/fiber"
	"github.com/patipolst/go-demo/pkg/api/route"
	"github.com/patipolst/go-demo/pkg/api/controller"
	"github.com/patipolst/go-demo/pkg/service"
)

func Run(ts *service.TodoService, us *service.UserService) {
	app := fiber.New()

	tc := controller.NewTodoController(ts)
	uc := controller.NewUserController(us)

	route.Todo(app, tc)
	route.User(app, uc)
	route.GraphQL(app, ts, us)

	const port = 3000
	log.Printf("Connect to http://localhost:%d/ for api", port)
	app.Listen(port)
}
