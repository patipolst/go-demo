package route

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/adaptor"
	"github.com/gofiber/fiber"
	"github.com/patipolst/go-demo/pkg/graph/generated"
	"github.com/patipolst/go-demo/pkg/graph/resolver"
	"github.com/patipolst/go-demo/pkg/service"
)

func GraphQL(app *fiber.App, ts *service.TodoService, us *service.UserService) {
	app.Post("/query", graphqlHandler(ts, us))
	app.Get("/", playgroundHandler())
}

func graphqlHandler(ts *service.TodoService, us *service.UserService) func(*fiber.Ctx) {
	r := resolver.New(ts, us)
	es := generated.NewExecutableSchema(generated.Config{Resolvers: r})
	h := handler.NewDefaultServer(es)
	return adaptor.HTTPHandler(h)
}

func playgroundHandler() func(*fiber.Ctx) {
	h := playground.Handler("GraphQL", "/query")
	return adaptor.HTTPHandler(h)
}
