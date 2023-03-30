package services

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

const (
	sudentEndpoint = "/student"
	newEndpoint    = "/new"
)

type Handler struct {
	*zap.Logger
}

func CreateRouter(logger *zap.Logger) http.Handler {
	handler := Handler{Logger: logger}

	r := chi.NewRouter()

	r.Route(sudentEndpoint, func(r chi.Router) {
		r.Get("/", handlerF)

		r.Route("/teste", func(r chi.Router) {
			r.Get("/", handlerF)
		})
	})

	r.Put(newEndpoint, handler.newUser)

	return r
}

func handlerF(w http.ResponseWriter, r *http.Request) {
	resp, _ := json.Marshal("OK!")

	w.Write(resp)
}

func (h *Handler) newUser(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	query.Get("")

}
