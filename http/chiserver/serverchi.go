package chiserver

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Setup() (*chi.Mux) {
	mux := chi.NewRouter()
	mux.Use(middleware.Logger)

	mux.Route("/chi", func(r chi.Router) {
		r.Route("/xml", func(r chi.Router) {
			addXPath(r)
		})
		r.Route("/cookie", func(r chi.Router) {
			addCookie(r)
		})
	})
	return mux
}
