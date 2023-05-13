package services

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

const (
	healthPath     = "/health"
	newPath        = "/new"
	studentPath    = "/student"
	teacherPath    = "/teacher"
	attendancePath = "/sendAttendance"
)

func CreateRouter(provider *Provider) http.Handler {
	handler := Handler{Provider: provider}

	r := chi.NewRouter()

	r.Get(healthPath, handler.HealthCheckHandler)

	r.Post(attendancePath, handler.AttendanceReceiveHandler)

	return r
}
