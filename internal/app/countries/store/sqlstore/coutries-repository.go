package sqlstore

import (
	smath "go-grcp-test/internal/app/countries/math"
	"go-grcp-test/internal/app/countries/model"

	"gorm.io/gorm"
)

// CountriesRepository ...
type CountriesRepository struct {
	store *Store
}

// GetList ...
func (cr *CountriesRepository) GetList(pagination model.Pagination) ([]model.Country, model.Pagination, error) {
	var countries []model.Country
	result := cr.store.db.Offset(pagination.Offset).Limit(pagination.Limit).Find(&countries)
	if err := result.Error; err != nil {
		return nil, model.Pagination{}, err
	}
	length := len(countries)
	pagination.Offset += length
	if length < pagination.Limit {
		pagination.Limit = length
	}
	return countries, pagination, nil
}

// Like ...
func (cr *CountriesRepository) Like(id uint32) (model.Country, error) {
	var country model.Country
	err := cr.store.db.Transaction(func(tx *gorm.DB) error {
		if err := cr.store.db.First(&country, id).Error; err != nil {
			return err
		}
		country.Likes++
		if err := cr.store.db.Updates(&country).Error; err != nil {
			return err
		}

		return nil
	})
	return country, err
}

// Dislike ...
func (cr *CountriesRepository) Dislike(id uint32) (model.Country, error) {
	var country model.Country
	err := cr.store.db.Transaction(func(tx *gorm.DB) error {
		if err := cr.store.db.First(&country, id).Error; err != nil {
			return err
		}
		country.Likes = smath.Max(0, country.Likes-1)
		if err := cr.store.db.Updates(&country).Error; err != nil {
			return err
		}

		return nil
	})
	return country, err
}

// FindByID ...
func (cr *CountriesRepository) FindByID(id uint32) (model.Country, error) {
	var country model.Country
	result := cr.store.db.First(&country, id)
	return country, result.Error
}
