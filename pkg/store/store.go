package store

// Store interface for all of the supported model Stores
type Store interface {
	Temperatures() TemperatureStore
}

type store struct {
	name         string
	temperatures TemperatureStore
}

func (s *store) Temperatures() TemperatureStore {
	return s.temperatures
}

// New creates a new Store
func New(name string, temperatures TemperatureStore) Store {
	return &store{name, temperatures}
}
