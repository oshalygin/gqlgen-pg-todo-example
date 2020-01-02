package resolvers

import (
	"context"

	"github.com/oshalygin/gqlgen-pg-todo-example/dataloaders"
	"github.com/oshalygin/gqlgen-pg-todo-example/graph/generated"
	"github.com/oshalygin/gqlgen-pg-todo-example/models"
)

type todoResolver struct{ *Resolver }

func (r *queryResolver) Todo(ctx context.Context, id int) (*models.Todo, error) {
	todo := models.Todo{ID: id}

	if err := r.DB.Select(&todo); err != nil {
		return nil, err
	}

	return &todo, nil
}
func (r *queryResolver) Todos(ctx context.Context, limit *int, offset *int) ([]models.Todo, error) {
	var todos []models.Todo

	if err := r.DB.Model(&todos).Select(); err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *todoResolver) CreatedBy(ctx context.Context, obj *models.Todo) (*models.User, error) {
	return ctx.Value(dataloaders.UserLoader).(*generated.UserLoader).Load(obj.CreatedBy)
}
func (r *todoResolver) UpdatedBy(ctx context.Context, obj *models.Todo) (*models.User, error) {
	return ctx.Value(dataloaders.UserLoader).(*generated.UserLoader).Load(obj.UpdatedBy)
}
