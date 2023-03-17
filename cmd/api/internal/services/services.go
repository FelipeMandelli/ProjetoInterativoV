package services

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

const (
	sudentEndpoint = "/student"
)

func CreateRouter() http.Handler {

	r := chi.NewRouter()

	r.Route(sudentEndpoint, func(r chi.Router) {
		r.Get("/", handler)

		r.Route("/teste", func(r chi.Router) {
			r.Get("/", handler)
		})
	})

	r.Put("/testes", handler)

	return r
}

func handler(w http.ResponseWriter, r *http.Request) {
	resp, _ := json.Marshal("OK!")

	w.Write(resp)
}
