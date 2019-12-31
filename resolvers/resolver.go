//go:generate go run github.com/99designs/gqlgen -v
package resolvers

import (
	"github.com/go-pg/pg/v9"
	"github.com/oshalygin/gqlgen-pg-todo-example/graph/generated"
)

type Resolver struct {
	DB *pg.DB
}

func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Todo() generated.TodoResolver {
	return &todoResolver{r}
}

type mutationResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }
