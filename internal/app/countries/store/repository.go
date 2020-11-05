package store

import (
	"go-grcp-test/internal/app/countries/model"
)

// CountriesRepository ...
type CountriesRepository interface {
	GetList(model.Pagination) ([]model.Country, model.Pagination, error)
	Like(uint32) (model.Country, error)
	Dislike(uint32) (model.Country, error)
	FindByID(uint32) (model.Country, error)
}
