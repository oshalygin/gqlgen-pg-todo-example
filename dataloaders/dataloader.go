package dataloaders

import (
	"net/http"

	"github.com/go-pg/pg/v9"
)

// The following are context keys which will be referenced for various loaders
const UserLoader = "userLoader"

func setLoader(db *pg.DB, dataloader func(db *pg.DB, w http.ResponseWriter, r *http.Request, next http.Handler)) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			dataloader(db, w, r, next)
		})
	}
}

func NewMiddleware(session *pg.DB) []func(handler http.Handler) http.Handler {
	return []func(handler http.Handler) http.Handler{
		setLoader(session, User),
	}
}
