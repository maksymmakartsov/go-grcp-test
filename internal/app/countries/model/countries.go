package model

import (
	"go-grcp-test/pkg/api"

	"gorm.io/gorm"
)

// Country ...
type Country struct {
	gorm.Model
	Name  string
	Likes int32
}

// Pagination ...
type Pagination struct {
	Offset int
	Limit  int
}

// ToResponseCountry ...
func (c *Country) ToResponseCountry() *api.CountriesRepositoryResponse_Country_Single {
	return &api.CountriesRepositoryResponse_Country_Single{Id: uint32(c.ID), Name: c.Name, Likes: c.Likes}
}
