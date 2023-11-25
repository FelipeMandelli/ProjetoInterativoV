package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	domain "github.com/FelipeMandelli/ProjetoInterativoV/cmd/api/internal/domain/rest"
	dto "github.com/FelipeMandelli/ProjetoInterativoV/pkg/DTO"
	"github.com/FelipeMandelli/ProjetoInterativoV/pkg/entities"
)

type Handler struct {
	Provider *Provider
}

func (h *Handler) AttendanceHandler(w http.ResponseWriter, r *http.Request) {
	h.Provider.Log.Debug("Received attendance request!")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.Provider.Log.Sugar().Error("error reading request body: ", err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	receivedBody := new(domain.AttendanceRequest)

	if err := json.Unmarshal(body, &receivedBody); err != nil {
		h.Provider.Log.Sugar().Error("error unmarshalling request body: ", err)
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
		h.Provider.Log.Sugar().Error("error reading request body", err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	receivedBody := new(entities.Registry)

	if err := json.Unmarshal(body, &receivedBody); err != nil {
		h.Provider.Log.Sugar().Error("error unmarshalling request body", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	if !receivedBody.IsValidRole() {
		h.Provider.Log.Sugar().Warnf("invalid received [role] for registration: %+v", receivedBody.Role)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	switch receivedBody.Role {
	case string(entities.StudentRole):
		if !receivedBody.IsValidCourse() {
			h.Provider.Log.Sugar().Warnf("invalid received [course] for registration: %+v", receivedBody.Course)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	case string(entities.ProfessorRole):
		receivedBody.Course = ""
	}

	dto := dto.RegistryDTO{
		Registry: *receivedBody,
	}

	h.Provider.RegChan <- dto

	h.Provider.Log.Sugar().Infof("received info: %+v", *receivedBody)
}

func (h *Handler) NewSubjectHandler(w http.ResponseWriter, r *http.Request) {
	h.Provider.Log.Debug("Received New Subject request!")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.Provider.Log.Sugar().Error("error reading request body", err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	receivedBody := new(entities.SubjectRegistry)

	if err := json.Unmarshal(body, &receivedBody); err != nil {
		h.Provider.Log.Sugar().Error("error unmarshalling request body", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	if receivedBody.Name == "" ||
		receivedBody.ProfessorID == "" ||
		receivedBody.Schedule == 0 ||
		receivedBody.Semester == 0 ||
		receivedBody.StudentsEnrolled == "" ||
		receivedBody.WeekDay == 0 ||
		receivedBody.Year == "" ||
		receivedBody.ProfessorName == "" {

		str := fmt.Sprintf("invalid received request for registration, fill all information: %+v\ntemplate: %+v", receivedBody, entities.SubjectRegistry{})

		h.Provider.Log.Sugar().Warn(str)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(str))

		return
	}

	h.Provider.SubChan <- dto.SubjectRegistryDTO{
		Registry: *receivedBody,
	}

	h.Provider.Log.Sugar().Infof("received info: %+v", *receivedBody)
}
