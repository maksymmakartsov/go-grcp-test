package store

// Store ...
type Store interface {
	Countries() CountriesRepository
}
