package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"temperature-api/internal/api/http/handlers"
)

// New собирает HTTP роутер сервиса temperature-api.
func New() http.Handler {
	r := chi.NewRouter()

	healthHandler := handlers.NewHealthHandler()
	r.Get("/health", healthHandler.ServeHTTP)

	temperatureHandler := handlers.NewTemperatureHandler()
	r.Get("/temperature", temperatureHandler.GetByLocation)
	r.Get("/temperature/{id}", temperatureHandler.GetByID)

	return r
}
