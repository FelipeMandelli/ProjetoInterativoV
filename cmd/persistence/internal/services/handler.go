package services

import (
	"encoding/json"
	"io"
	"net/http"

	dto "github.com/FelipeMandelli/ProjetoInterativoV/pkg/DTO"
	"github.com/FelipeMandelli/ProjetoInterativoV/pkg/entities"
)

type Handler struct {
	Provider *Provider
}

func (h *Handler) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal("Persistence is Healthy!")
	if err != nil {
		h.Provider.Log.Sugar().Errorf("error marshalling response: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	h.Provider.Log.Sugar().Infof("Health Checked by IP %s", r.RemoteAddr)
	w.Write(resp)
}

func (h *Handler) AttendanceReceiveHandler(w http.ResponseWriter, r *http.Request) {
	h.Provider.Log.Debug("Received Attendance request!")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.Provider.Log.Sugar().Error("error reading request body: ", err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	receivedPack := new(dto.PackagerDTO)

	if err := json.Unmarshal(body, &receivedPack); err != nil {
		h.Provider.Log.Sugar().Error("error unmarshalling received pack: ", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	if h.Provider.DbIsON {
		err := PersistAtendance(h.Provider, receivedPack.TeacherID, receivedPack.AttendanceIDs, receivedPack.SendingTime)
		if err != nil {
			h.Provider.Log.Sugar().Error("error persisting received pack: ", err)
		}
	}

	h.Provider.Log.Sugar().Infof("received info: %+v", *receivedPack)
}

func (h *Handler) NewRegistryReceiveHandler(w http.ResponseWriter, r *http.Request) {
	h.Provider.Log.Debug("Received Registry request!")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.Provider.Log.Sugar().Errorf("error reading request body: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	receivedRegistry := new(dto.RegistryDTO)

	if err := json.Unmarshal(body, receivedRegistry); err != nil {
		h.Provider.Log.Sugar().Errorf("error unmarshalling received registry: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	if h.Provider.DbIsON {
		if receivedRegistry.Registry.Role == string(entities.StudentRole) {
			h.Provider.Log.Sugar().Infof("received student registry [%+v]", receivedRegistry)
			err := PersistStudentRegistry(h.Provider, receivedRegistry.Registry)
			if err != nil {
				h.Provider.Log.Sugar().Error("error persisting received student registry: ", err)
			}
		}

		if receivedRegistry.Registry.Role == string(entities.ProfessorRole) {
			h.Provider.Log.Sugar().Infof("received professor registry [%+v]", receivedRegistry)
			err := PersistProfessorRegistry(h.Provider, receivedRegistry.Registry)
			if err != nil {
				h.Provider.Log.Sugar().Error("error persisting received teacher registry: ", err)
			}
		}

	}

	h.Provider.Log.Sugar().Infof("received info: %+v", *receivedRegistry)
}

func (h *Handler) NewSubjectRegistryReceiveHandler(w http.ResponseWriter, r *http.Request) {
	h.Provider.Log.Debug("Received Subject Registry request!")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.Provider.Log.Sugar().Errorf("error reading request body: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	receivedSubReg := new(dto.SubjectRegistryDTO)

	if err := json.Unmarshal(body, receivedSubReg); err != nil {
		h.Provider.Log.Sugar().Errorf("error unmarshalling received registry: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	if h.Provider.DbIsON {
		err = PersistSubjectRegistry(h.Provider, receivedSubReg.Registry)
		if err != nil {
			h.Provider.Log.Sugar().Error("error persisting received subject registry: ", err)
		}
	}

	h.Provider.Log.Sugar().Infof("received info: %+v", *receivedSubReg)
}
