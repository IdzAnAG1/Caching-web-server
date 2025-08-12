package app

import (
	"CachingWebServer/internal/config"
	"CachingWebServer/internal/lib/logger"
	"flag"
	"fmt"
)

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (a *App) Launch() error {
	flag.Parse()

	cfg := config.MustLoadConfig()

	fmt.Println("Config was loaded")

	log, err := logger.IdentifyLogger(cfg)
	if err != nil {

	}

	log.Info("Ура")
	return nil
}
