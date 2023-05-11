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
	attendancePath = "/attendance"
)

func CreateRouter(provider *Provider) http.Handler {
	handler := Handler{Provider: provider}

	r := chi.NewRouter()

	r.Get(healthPath, handler.HealthCheckHandler)

	r.Route(attendancePath, func(r chi.Router) {
		r.Post("/", handler.AttendanceHandler)
	})

	r.Route(newPath, func(r chi.Router) {
		r.Post(studentPath, handler.NewStudentHandler)
		r.Post(teacherPath, handler.NewTeacherHandler)
	})

	return r
}
