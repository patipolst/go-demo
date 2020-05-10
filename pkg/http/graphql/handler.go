package graphql

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/patipolst/go-demo/pkg/graph/generated"
	"github.com/patipolst/go-demo/pkg/graph/resolver"
	"github.com/patipolst/go-demo/pkg/service"
)

const defaultPort = "8080"

// Run starts the graphql playground
func Run(s *service.TodoService) {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{s}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
