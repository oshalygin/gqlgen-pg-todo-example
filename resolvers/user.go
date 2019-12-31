package resolvers

import (
	"context"

	"github.com/oshalygin/gqlgen-pg-todo-example/models"
)

func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	panic("not implemented")
}
func (r *queryResolver) Users(ctx context.Context, limit *int, offset *int) ([]models.User, error) {
	panic("not implemented")
}

func (r *mutationResolver) UserCreate(ctx context.Context, user models.UserInput) (*models.User, error) {
	panic("not implemented")
}
