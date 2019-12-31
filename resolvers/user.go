package resolvers

import (
	"context"
	"time"

	"github.com/oshalygin/gqlgen-pg-todo-example/models"
)

func (r *queryResolver) User(ctx context.Context, id int) (*models.User, error) {
	user := models.User{ID: id}

	if err := r.DB.Select(&user); err != nil {
		return nil, err
	}

	return &user, nil
}
func (r *queryResolver) Users(ctx context.Context, limit *int, offset *int) ([]models.User, error) {
	var users []models.User

	if err := r.DB.Model(&users).Select(); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *mutationResolver) UserCreate(ctx context.Context, user models.UserInput) (*models.User, error) {
	usr := models.User{
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := r.DB.Insert(&usr); err != nil {
		return nil, err
	}

	return &usr, nil
}
