package api

import (
	"log"

	"github.com/gofiber/fiber"
	"github.com/patipolst/go-demo/pkg/api/controller"
	"github.com/patipolst/go-demo/pkg/api/route"
	"github.com/patipolst/go-demo/pkg/graph/resolver"
	"github.com/patipolst/go-demo/pkg/service"
)

func Run(ts *service.Todo, us *service.User) {
	app := fiber.New()

	tc := controller.NewTodo(ts)
	uc := controller.NewUser(us)
	r := resolver.New(ts, us)

	route.Todo(app, tc)
	route.User(app, uc)
	route.GraphQL(app, r)

	const port = 3000
	log.Printf("Connect to http://localhost:%d/ for api", port)
	app.Listen(port)
}
