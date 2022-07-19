package main

import (
	"flag"
	"log"

	config "github.com/honyshyota/l0-wb-test/configs"
	"github.com/honyshyota/l0-wb-test/internal/app"
	"github.com/sirupsen/logrus"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "configs/config.toml", "path to config file")
}

func main() {
	// Init configuration
	flag.Parse()

	config, err := config.NewConfig()
	if err != nil {
		logrus.Fatalln("Unable to load config")
	}

	if err = app.Run(config); err != nil {
		log.Fatal("Unable to run app", err)
	}
}
