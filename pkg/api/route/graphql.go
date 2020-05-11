package route

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/adaptor"
	"github.com/gofiber/fiber"
	"github.com/patipolst/go-demo/pkg/graph/generated"
	"github.com/patipolst/go-demo/pkg/graph/resolver"
)

func GraphQL(app *fiber.App, r *resolver.Resolver) {
	app.Post("/query", graphqlHandler(r))
	app.Get("/", playgroundHandler())
}

func graphqlHandler(r *resolver.Resolver) func(*fiber.Ctx) {
	es := generated.NewExecutableSchema(generated.Config{Resolvers: r})
	h := handler.NewDefaultServer(es)
	return adaptor.HTTPHandler(h)
}

func playgroundHandler() func(*fiber.Ctx) {
	h := playground.Handler("GraphQL", "/query")
	return adaptor.HTTPHandler(h)
}
