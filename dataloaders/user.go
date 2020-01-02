package dataloaders

import (
	"context"
	"net/http"
	"time"

	"github.com/go-pg/pg/v9"
	"github.com/oshalygin/gqlgen-pg-todo-example/graph/generated"
	"github.com/oshalygin/gqlgen-pg-todo-example/models"
)

func User(db *pg.DB, w http.ResponseWriter, r *http.Request, next http.Handler) {
	loader := generated.NewUserLoader(generated.UserLoaderConfig{
		MaxBatch: 100,
		Wait:     1 * time.Millisecond,
		Fetch: func(keys []int) ([]*models.User, []error) {

			var users []*models.User
			err := db.Model(&users).WhereIn("id IN (?)", keys).Select()

			if err != nil {
				return []*models.User{}, []error{err}
			}

			return users, []error{err}

		},
	})

	ctx := context.WithValue(r.Context(), UserLoader, loader)
	next.ServeHTTP(w, r.WithContext(ctx))
}
