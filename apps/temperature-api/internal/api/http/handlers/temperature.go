package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"temperature-api/internal/api/http/dto"
	"temperature-api/internal/core/temperature"
)

// TemperatureHandler обрабатывает запросы показаний датчиков температуры.
type TemperatureHandler struct {
}

// NewTemperatureHandler создает обработчик показаний датчиков температуры.
func NewTemperatureHandler() *TemperatureHandler {
	return &TemperatureHandler{}
}

// GetByLocation отдает показание датчика по названию комнаты, переданному в query-параметре location.
func (h *TemperatureHandler) GetByLocation(w http.ResponseWriter, r *http.Request) {
	location := r.URL.Query().Get("location")
	if location == "" {
		writeJSON(w, http.StatusBadRequest, dto.ErrorResponse{Error: "location is required"})
		return
	}

	reading := temperature.NewReading(location, temperature.ResolveSensorID(location))
	writeJSON(w, http.StatusOK, reading)
}

// GetByID отдает показание датчика по идентификатору, переданному в пути запроса.
func (h *TemperatureHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	sensorID := chi.URLParam(r, "id")
	reading := temperature.NewReading(temperature.ResolveLocation(sensorID), sensorID)
	writeJSON(w, http.StatusOK, reading)
}

// writeJSON записывает тело ответа в формате JSON с указанным статусом.
func writeJSON(w http.ResponseWriter, status int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(body)
}
