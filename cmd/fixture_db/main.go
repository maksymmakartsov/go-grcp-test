package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-testfixtures/testfixtures/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Country ...
type Country struct {
	gorm.Model
	Name  string
	Likes int32
}

type postgresConf struct {
	user     string
	password string
	dbname   string
}

func (pg *postgresConf) dsn() string {
	return fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", pg.user, pg.password, pg.dbname)
}

func newPostgresConf() *postgresConf {
	// Add your db config
	return &postgresConf{
		user:     os.Getenv("PG_USER"),
		password: os.Getenv("PG_PASS"),
		dbname:   os.Getenv("PG_DB_NAME"),
	}
}

// Before run this file setup newPostgresConf
func main() {
	dsn := newPostgresConf().dsn()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database %v", err)
	}
	if err := db.AutoMigrate(&Country{}); err != nil {
		log.Fatalf("failed to make migration: %v", err)
	}
	addFixtures(db)
}

// This method cleanup database before add fixtures
func addFixtures(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed get generic DB %v", err)
	}
	fixtures, err := testfixtures.New(
		testfixtures.Database(sqlDB),
		testfixtures.Dialect("postgresql"),
		testfixtures.Directory(os.Getenv("PG_FIXTURES")),
		testfixtures.DangerousSkipTestDatabaseCheck(), //Disable check on test database
	)
	if err != nil {
		log.Fatalf("Failed add fixtures %v", err)
	}

	if err := fixtures.Load(); err != nil {
		log.Fatalf("Failed to load fixtures %v", err)
	}
}
