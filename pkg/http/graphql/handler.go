package graphql

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/patipolst/go-demo/pkg/graph/generated"
	"github.com/patipolst/go-demo/pkg/graph/model"
	"github.com/patipolst/go-demo/pkg/graph/resolver"
	"github.com/patipolst/go-demo/pkg/mutation"
	"github.com/patipolst/go-demo/pkg/query"
	"github.com/patipolst/go-demo/pkg/store/memory"
)

const defaultPort = "8080"

// Run starts the graphql playground
func Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// mockTodos := make([]*model.Todo, 0)
	var mockTodos []*model.Todo
	todoStore := memory.NewTodoStore()
	todoQuery := query.NewTodoQuery(todoStore)
	todoMutation := mutation.NewTodoMutation(todoStore)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{mockTodos, todoQuery, todoMutation}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
