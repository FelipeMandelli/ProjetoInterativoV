package services

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

const (
	healthPath     = "/health"
	newPath        = "/new/registry"
	attendancePath = "/sendAttendance"
)

func CreateRouter(provider *Provider) http.Handler {
	handler := Handler{Provider: provider}

	r := chi.NewRouter()

	r.Get(healthPath, handler.HealthCheckHandler)

	r.Post(newPath, handler.NewRegistryReceiveHandler)

	r.Post(attendancePath, handler.AttendanceReceiveHandler)

	return r
}
