package countries

import (
	"context"
	"go-grcp-test/pkg/api"
	"io"
	"log"
)

// GRPCServer ...
type GRPCServer struct{}

// Single ...
func (s *GRPCServer) Single(req *api.Request_Country_Single, stream api.Countries_SingleServer) error {
	err := stream.Send(&api.Response_Country_Single{Id: 1, Name: "Ukraine", Likes: 333})
	if err != nil {
		return err
	}
	return nil
}

// List ...
func (s *GRPCServer) List(ctx context.Context, req *api.Request_Country_List) (*api.Response_Country_List, error) {
	countriesList := []*api.Response_Country_Single{&api.Response_Country_Single{Id: 1, Name: "Ukraine", Likes: 333}, &api.Response_Country_Single{Id: 2, Name: "USA", Likes: 1}}
	return &api.Response_Country_List{Countries: countriesList}, nil
}

// LikeDislike ...
func (s *GRPCServer) LikeDislike(stream api.Countries_LikeDislikeServer) error {
	for {
		countryID, err := stream.Recv()
		if err == io.EOF {
			log.Printf("%s", countryID)
			return stream.SendAndClose(&api.Response_Country_Single{Id: 1, Name: "Ukraine", Likes: 333})
		}
		if err != nil {
			return err
		}
	}
}
