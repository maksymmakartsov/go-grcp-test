package sqlstore

import (
	"go-grcp-test/internal/app/countries/store"

	"gorm.io/gorm"
)

// Store ...
type Store struct {
	db                  *gorm.DB
	countriesRepository *CountriesRepository
}

// New ...
func New(db *gorm.DB) *Store {
	return &Store{
		db: db,
	}
}

// Countries ...
func (s *Store) Countries() store.CountriesRepository {
	if s.countriesRepository != nil {
		return s.countriesRepository
	}

	s.countriesRepository = &CountriesRepository{
		store: s,
	}

	return s.countriesRepository
}
