package temperature

import (
	"fmt"
	"math/rand"
	"time"

	"temperature-api/internal/api/http/dto"
)

const (
	minValue              = 18.0
	maxValue              = 28.0
	unitCelsius           = "C"
	statusOK              = "ok"
	sensorTypeTemperature = "temperature"
	unknownLocation       = "Unknown"
	unknownSensorID       = "0"
)

var sensorIDByLocation = map[string]string{
	"Living Room": "1",
	"Bedroom":     "2",
	"Kitchen":     "3",
}

var locationBySensorID = map[string]string{
	"1": "Living Room",
	"2": "Bedroom",
	"3": "Kitchen",
}

// ResolveLocation возвращает название комнаты по идентификатору датчика либо Unknown, если идентификатор не распознан.
func ResolveLocation(sensorID string) string {
	if location, ok := locationBySensorID[sensorID]; ok {
		return location
	}
	return unknownLocation
}

// ResolveSensorID возвращает идентификатор датчика по названию комнаты либо 0, если комната не распознана.
func ResolveSensorID(location string) string {
	if sensorID, ok := sensorIDByLocation[location]; ok {
		return sensorID
	}
	return unknownSensorID
}

// NewReading формирует показание датчика со случайным значением температуры для указанной комнаты и идентификатора.
func NewReading(location, sensorID string) dto.Temperature {
	return dto.Temperature{
		Value:       minValue + rand.Float64()*(maxValue-minValue),
		Unit:        unitCelsius,
		Timestamp:   time.Now(),
		Location:    location,
		Status:      statusOK,
		SensorID:    sensorID,
		SensorType:  sensorTypeTemperature,
		Description: describe(location),
	}
}

// describe формирует описание показания по названию комнаты.
func describe(location string) string {
	return fmt.Sprintf("Temperature sensor in %s", location)
}
