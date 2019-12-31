package resolvers

import (
	"context"

	"github.com/oshalygin/gqlgen-pg-todo-example/models"
)

type todoResolver struct{ *Resolver }

func (r *queryResolver) Todo(ctx context.Context, id string) (*models.Todo, error) {
	panic("not implemented")
}
func (r *queryResolver) Todos(ctx context.Context, limit *int, offset *int) ([]models.Todo, error) {
	panic("not implemented")
}

func (r *todoResolver) CreatedBy(ctx context.Context, obj *models.Todo) (*models.User, error) {
	panic("not implemented")
}
func (r *todoResolver) UpdatedBy(ctx context.Context, obj *models.Todo) (*models.User, error) {
	panic("not implemented")
}
