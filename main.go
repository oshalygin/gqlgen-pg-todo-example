package main

import (
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"
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

		srv := handler.NewDefaultServer(schema)
		srv.Use(extension.FixedComplexityLimit(300))

		r.Handle("/", srv)
	})

	gqlPlayground := playground.Handler("api-gateway", "/graphql")
	r.Get("/", gqlPlayground)

	startMessage()
	panic(http.ListenAndServe(port, r))
}
