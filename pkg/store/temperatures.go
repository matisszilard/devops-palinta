package store

import "github.com/matisszilard/devops-palinta/pkg/model"

// TemperatureStore intarface handle methods with the temperature connected methods
type TemperatureStore interface {
	// Save a temperature struct into the database
	Save(temp model.Temperature)
}
