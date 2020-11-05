package countriesserver

import "os"

// Config ...
type Config struct {
	BindAddr string `toml:"bind_addr"`
	User     string
	Password string
	Dbname   string
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		User:     os.Getenv("PG_USER"),
		Password: os.Getenv("PG_PASS"),
		Dbname:   os.Getenv("PG_DB_NAME"),
	}
}
