package main

import (
	"flag"
	"go-grcp-test/internal/app/countries/countriesserver"
	"log"

	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/countries-server.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := countriesserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	if err := countriesserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
