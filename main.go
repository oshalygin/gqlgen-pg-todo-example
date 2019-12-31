package main

import (
	"net/http"

	"github.com/99designs/gqlgen/example/starwars/generated"
	"github.com/99designs/gqlgen/handler"
	"github.com/go-chi/chi"
)

const (
	port = ":8080"
)

func main() {
	r := chi.NewRouter()

	// The base path that users would use is POST /graphql which is fairly
	// idiomatic.
	r.Route("/graphql", func(r chi.Router) {
		schema := generated.NewExecutableSchema(generated.Config{
			//Resolvers:  &resolvers.Resolver{},
			Directives: generated.DirectiveRoot{},
			Complexity: generated.ComplexityRoot{},
		})

		h := handler.GraphQL(
			schema,
			handler.ComplexityLimit(100),
		)
		// API Layer
		r.Post("/", h)

		// This is the UI playground that can be used to interact with the API layer
		// http://localhost:8080/graphql
		playground := handler.Playground("Interactive Playground", "/graphql")
		r.Get("/", playground)

	})

	panic(http.ListenAndServe(port, r))
}
