package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/nurtaims/golang-mongodb/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config_path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()
	// Config
	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
	// APIServer
	server := apiserver.New(config)
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
