package services

import (
	"encoding/json"
	"io"
	"net/http"

	domain "github.com/FelipeMandelli/ProjetoInterativoV/cmd/api/internal/domain/rest"
	entitys "github.com/FelipeMandelli/ProjetoInterativoV/pkg/Entitys"
)

type Handler struct {
	Provider *Provider
}

func (h *Handler) AttendanceHandler(w http.ResponseWriter, r *http.Request) {
	h.Provider.Log.Debug("Received attendance request!")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.Provider.Log.Sugar().Error("error reading request body")
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	receivedBody := new(domain.AttendanceRequest)

	if err := json.Unmarshal(body, &receivedBody); err != nil {
		h.Provider.Log.Sugar().Error("error unmarshalling request body")
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	h.Provider.RequestBodyChan <- *receivedBody

	h.Provider.Log.Sugar().Infof("received tag: %s", receivedBody.Tag)
}

func (h *Handler) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal("API is Healthy!")
	if err != nil {
		h.Provider.Log.Sugar().Errorf("error marshalling response: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	h.Provider.Log.Sugar().Infof("Health Checked by IP %s", r.RemoteAddr)
	w.Write(resp)
}

func (h *Handler) NewRegistryHandler(w http.ResponseWriter, r *http.Request) {
	h.Provider.Log.Debug("Received New Registry request!")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.Provider.Log.Sugar().Error("error reading request body")
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	receivedBody := new(entitys.Resgistry)

	if err := json.Unmarshal(body, &receivedBody); err != nil {
		h.Provider.Log.Sugar().Error("error unmarshalling request body")
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	if !receivedBody.IsValidRole() {
		h.Provider.Log.Sugar().Warnf("invalid received [role] for registration: %+v", receivedBody.Role)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	switch receivedBody.Role {
	case entitys.StudentRole:
		if !receivedBody.IsValidCourse() {
			h.Provider.Log.Sugar().Warnf("invalid received [course] for registration: %+v", receivedBody.Course)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	case entitys.TeacherRole:
		receivedBody.Course = ""
	}

	h.Provider.Log.Sugar().Infof("received info: %+v", *receivedBody)
}

// func (h *Handler) NewStudentHandler(w http.ResponseWriter, r *http.Request) {
// 	h.Provider.Log.Debug("Received New Student request!")
// 	body, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		h.Provider.Log.Sugar().Error("error reading request body")
// 		w.WriteHeader(http.StatusInternalServerError)

// 		return
// 	}

// 	receivedBody := new(entitys.Student)

// 	if err := json.Unmarshal(body, &receivedBody); err != nil {
// 		h.Provider.Log.Sugar().Error("error unmarshalling request body")
// 		w.WriteHeader(http.StatusBadRequest)

// 		return
// 	}

// 	h.Provider.Log.Sugar().Infof("received info: %+v", receivedBody)
// }

// func (h *Handler) NewTeacherHandler(w http.ResponseWriter, r *http.Request) {
// 	h.Provider.Log.Debug("Received New Teacher request!")
// 	body, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		h.Provider.Log.Sugar().Error("error reading request body")
// 		w.WriteHeader(http.StatusInternalServerError)

// 		return
// 	}

// 	receivedBody := new(entitys.Teacher)

// 	if err := json.Unmarshal(body, &receivedBody); err != nil {
// 		h.Provider.Log.Sugar().Error("error unmarshalling request body")
// 		w.WriteHeader(http.StatusBadRequest)

// 		return
// 	}

// 	h.Provider.Log.Sugar().Infof("received info: %+v", receivedBody)
// }
