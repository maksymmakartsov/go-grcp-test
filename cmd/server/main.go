package main

import (
	"go-grcp-test/pkg/api"
	"go-grcp-test/pkg/countries"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	srv := &countries.GRPCServer{}
	api.RegisterCountriesServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
