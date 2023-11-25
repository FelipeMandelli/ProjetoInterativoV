package services

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

const (
	healthPath        = "/health"
	newRegistryPath   = "/new/registry"
	newPath           = "/new"
	studentPath       = "/student"
	teacherPath       = "/professor"
	attendancePath    = "/attendance"
	contentTypeHeader = "Content-Type"
	jsonContentType   = "application/json"
)

func CreateRouter(provider *Provider) http.Handler {
	handler := Handler{Provider: provider}

	r := chi.NewRouter()

	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set(contentTypeHeader, jsonContentType)
			next.ServeHTTP(w, r)
		})
	})

	r.Get(healthPath, handler.HealthCheckHandler)

	r.Post(newRegistryPath, handler.NewRegistryHandler)

	r.Route(attendancePath, func(r chi.Router) {
		r.Post("/", handler.AttendanceHandler)
	})

	return r
}
