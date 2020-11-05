package countriesserver

import (
	"fmt"
	"go-grcp-test/internal/app/countries/store/sqlstore"
	"go-grcp-test/pkg/api"
	"go-grcp-test/pkg/countries"
	"log"
	"net"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"google.golang.org/grpc"
)

func (c *Config) dsn() string {
	return fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", c.User, c.Password, c.Dbname)
}

// Start ...
func Start(config *Config) error {
	db, err := newDB(config.dsn())
	if err != nil {
		return err
	}
	store := sqlstore.New(db)

	s := grpc.NewServer()
	srv := countries.NewGRPCServer(store)
	api.RegisterCountriesRepositoryServer(s, srv)

	l, err := net.Listen("tcp", config.BindAddr)
	if err != nil {
		log.Fatal(err)
	}

	return s.Serve(l)
}

func newDB(dbURL string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()

	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
