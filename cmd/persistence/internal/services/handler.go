package services

import (
	"encoding/json"
	"io"
	"net/http"

	dto "github.com/FelipeMandelli/ProjetoInterativoV/pkg/DTO"
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

	h.Provider.Log.Sugar().Infof("received info: %+v", *receivedPack)
}

func (h *Handler) NewRegistryReceiveHandler(w http.ResponseWriter, r *http.Request) {
	h.Provider.Log.Debug("Received Attendance request!")
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

	h.Provider.Log.Sugar().Infof("received info: %+v", *receivedRegistry)
}
