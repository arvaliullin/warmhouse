package handlers

import "net/http"

// HealthHandler обрабатывает проверку доступности сервиса.
type HealthHandler struct {
}

// NewHealthHandler создает обработчик проверки доступности.
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// ServeHTTP отвечает статусом ok для запроса проверки доступности.
func (h *HealthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"ok"}`))
}
