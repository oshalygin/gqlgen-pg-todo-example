package main

import (
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/handler"
	"github.com/fatih/color"
	"github.com/go-chi/chi"
	"github.com/go-pg/pg/v9"
	"github.com/oshalygin/gqlgen-pg-todo-example/dataloaders"
	database "github.com/oshalygin/gqlgen-pg-todo-example/db"
	"github.com/oshalygin/gqlgen-pg-todo-example/graph/generated"
	"github.com/oshalygin/gqlgen-pg-todo-example/resolvers"
)

const (
	port = ":8080"
)

func lineSeparator() {
	fmt.Println("========")
}

func startMessage() {
	lineSeparator()
	color.Green("Listening on localhost%s\n", port)
	color.Green("Visit `http://localhost%s/graphql` in your browser\n", port)
	lineSeparator()
}

func main() {
	lineSeparator()
	// Create the database `todos` manually within postgres
	db := pg.Connect(&pg.Options{
		Database: "todos",
	})
	defer db.Close()

	err := database.Seed(db)
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()

	// The base path that users would use is POST /graphql which is fairly
	// idiomatic.
	r.Route("/graphql", func(r chi.Router) {
		// Initialize the dataloaders as middleware into our route
		r.Use(dataloaders.NewMiddleware(db)...)

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

	startMessage()
	panic(http.ListenAndServe(port, r))
}
