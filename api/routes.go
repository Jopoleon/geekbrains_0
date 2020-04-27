package api

import (
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (a *API) InitRouter() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(20 * time.Second))
	r.Group(func(r chi.Router) {
		r.MethodFunc("GET", "/books", a.GetBooks)
		r.MethodFunc("POST", "/books", a.CreateBook)
		r.MethodFunc("GET", "/authors", a.GetAuthors)
		r.MethodFunc("POST", "/authors", a.CreateAuthor)
	})
	a.Router = r
}
