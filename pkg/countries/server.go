package countries

import (
	"context"
	"go-grcp-test/internal/app/countries/model"
	"go-grcp-test/internal/app/countries/store"
	"go-grcp-test/pkg/api"
)

// GRPCServer ...
type GRPCServer struct {
	store store.Store
}

// NewGRPCServer ...
func NewGRPCServer(store store.Store) *GRPCServer {
	return &GRPCServer{
		store: store,
	}
}

// Single ...
func (s *GRPCServer) Single(
	ctx context.Context,
	request *api.CountriesRepositoryRequest_Country_Single) (*api.CountriesRepositoryResponse_Country_Single, error) {
	country, err := s.store.Countries().FindByID(request.Id)
	if err != nil {
		return nil, err
	}
	return country.ToResponseCountry(), nil
}

// List ...
func (s *GRPCServer) List(
	ctx context.Context,
	request *api.CountriesRepositoryRequest_Country_List) (*api.CountriesRepositoryResponse_Country_List, error) {
	pagination := model.Pagination{Offset: int(request.Offset), Limit: int(request.Limit)}
	countries, newPagination, err := s.store.Countries().GetList(pagination)
	if err != nil {
		return nil, err
	}
	meta := &api.CountriesRepositoryResponse_Country_List_Meta{Offset: uint32(newPagination.Offset), Limit: uint32(newPagination.Limit)}
	return &api.CountriesRepositoryResponse_Country_List{Countries: fmap(countries), Meta: meta}, nil
}

// Like ...
func (s *GRPCServer) Like(
	ctx context.Context,
	request *api.CountriesRepositoryRequest_Country_Like) (*api.CountriesRepositoryResponse_Country_Single, error) {
	country, err := s.store.Countries().Like(request.Id)
	if err != nil {
		return nil, err
	}
	return country.ToResponseCountry(), nil
}

// Dislike ...
func (s *GRPCServer) Dislike(
	ctx context.Context,
	request *api.CountriesRepositoryRequest_Country_Dislike) (*api.CountriesRepositoryResponse_Country_Single, error) {
	country, err := s.store.Countries().Dislike(request.Id)
	if err != nil {
		return nil, err
	}
	return country.ToResponseCountry(), nil
}

func fmap(c []model.Country) []*api.CountriesRepositoryResponse_Country_Single {
	countries := make([]*api.CountriesRepositoryResponse_Country_Single, len(c))
	for i, v := range c {
		countries[i] = v.ToResponseCountry()
	}
	return countries
}
