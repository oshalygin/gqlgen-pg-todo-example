package main

import (
	"net/http"

	"github.com/99designs/gqlgen/handler"
	"github.com/go-chi/chi"
	"github.com/go-pg/pg/v9"
	"github.com/oshalygin/gqlgen-pg-todo-example/graph/generated"
	"github.com/oshalygin/gqlgen-pg-todo-example/resolvers"
)

const (
	port = ":8080"
)

func main() {
	db := pg.Connect(&pg.Options{
		User: "postgres",
	})
	defer db.Close()

	r := chi.NewRouter()

	// The base path that users would use is POST /graphql which is fairly
	// idiomatic.
	r.Route("/graphql", func(r chi.Router) {
		schema := generated.NewExecutableSchema(generated.Config{
			Resolvers: &resolvers.Resolver{
				DB: db,
			},
			Directives: generated.DirectiveRoot{},
			Complexity: generated.ComplexityRoot{},
		})

		h := handler.GraphQL(
			schema,
			handler.ComplexityLimit(100),
		)
		r.Post("/", h)

		// This is the UI playground that can be used to interact with the API layer
		// http://localhost:8080/graphql
		playground := handler.Playground("Interactive Playground", "/graphql")
		r.Get("/", playground)

	})

	panic(http.ListenAndServe(port, r))
}
