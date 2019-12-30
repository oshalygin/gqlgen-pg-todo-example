<p align="center">
  <img alt="Application Logo" src="docs/logo.png" height="150" width="150" />
  <h3 align="center">GQLGen PG TODO Example</h3>
  <p align="center">A simple, no fuss, example thats updated regularly to stay current with the API landscape</p>
</p>

# Introduction

This project is intended to help newcomers to gqlgen and GraphQL. This isn't meant to be the ONLY way you should organize and build your application. I took extensively liberties here to minimize func calls and writing helper wrappers which do nothing more than confuse newcomers.

# Setup

1. **Install [Go 1.13 or greater](https://dl.google.com/go/go1.13.darwin-amd64.pkg)**.
   - The recommended approach is to use the installer to get started.
2. **Install Postgres**
   - In the interest of ensuring that you can build/run this application w/o a ton of fuss, 

# Updating GraphQL Models + CodeGen

The application uses ([GQLGEN](https://gqlgen.com)) to generate statically typed bindings for all of the grahql models.

1. Generate the Go model that represents the graph model. An example looks like the following:

```go
package graph

import (
	"time"
)

// User Graph datamodel
type User struct {
	ID     string `json:"id"`
	Email  string `json:"email"`

	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`

	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

```

2. Add/Update your model in the `/schema` folder of the application. An example looks like the following. Note that you must reference the graph model package with the `@goModel` directive.

```graphql
type User
  @goModel(
    model: "github.com/oshalygin/gqlgen-pg-todo-example/models/graph.User"
  ) {
  id: ID!
  email: String!

  firstName: String!
  lastName: String!

  createdAt: Time!
  updatedAt: Time!
}
```

3. Run the codegen process

```bash
make gen
```

4. Implement the resolver within the resolvers package.

# Generating Dataloaders

```bash
# Run the following commands to generate dataloaders
# This command will place a codegen dataloader in the graph/generated folder
# Note that the loader argument is required and the value is case sensitive
make dataloader loader=User
```

You must still create the appropriate dataloader in the codebase to support
the generated implementation.

```go

// Create the loader that resembles other loaders in the dataloaders package
// This is merely an example

func User(session *mgo.Session, w http.ResponseWriter, r *http.Request, next http.Handler) {
	loader := generated.NewUserLoader(generated.UserLoaderConfig{
		MaxBatch: 100,
		Wait:     1 * time.Millisecond,
		Fetch: func(keys []string) ([]*graph.User, []error) {
			s := session.Copy()
			dal := db.UserDAL{}
			users, err := dal.FindAllByID(s, keys)

			if err != nil {
				return []*graph.User{}, []error{err}
			}

			graphModels := make([]*graph.User, len(users))

			for i, user := range users {
				graphModel := user.ToGraph()
				graphModels[i] = &graphModel
			}

			return graphModels, nil
		},
	})

	ctx := context.WithValue(r.Context(), UserLoader, loader)
	next.ServeHTTP(w, r.WithContext(ctx))
}


// Add a new entry in the NewMiddleware function

func NewMiddleware(session *mgo.Session) []func(handler http.Handler) http.Handler {
	return []func(handler http.Handler) http.Handler{
		setLoader(session, User),
		setLoader(session, Building),
    // setLoader(session, YourLoader)
	}
}

```

# Limitations

Will list if I encounter any

# Dependencies

| **Tech**                                      | **Description**                                                         |
| --------------------------------------------- | ----------------------------------------------------------------------- |
| [gqlgen](https://github.com/99designs/gqlgen) | `gqlgen` is a Go library for building GraphQL servers without any fuss. |

# License

MIT
