package routing

import (
	"net/http"

	"github.com/go-chi/chi"
)

func CreateRouter() chi.Router {
	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Elemental"))
	})

	router.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Successfully up and running"))
	})

	return router
}
