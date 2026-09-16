package dto

import "time"

// Temperature представляет показание датчика температуры.
type Temperature struct {
	Value       float64   `json:"value"`
	Unit        string    `json:"unit"`
	Timestamp   time.Time `json:"timestamp"`
	Location    string    `json:"location"`
	Status      string    `json:"status"`
	SensorID    string    `json:"sensor_id"`
	SensorType  string    `json:"sensor_type"`
	Description string    `json:"description"`
}

// ErrorResponse представляет тело ответа с ошибкой.
type ErrorResponse struct {
	Error string `json:"error"`
}
