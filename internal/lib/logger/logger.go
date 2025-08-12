package logger

import (
	"CachingWebServer/internal/config"
	"fmt"
)

const (
	DEBUG = "debug"
	LOCAL = "local"
	PROD  = "prod"
)

func Level(cfg config.Config) {
	switch cfg.Logger.Level {
	case DEBUG:
		fmt.Printf("level : %s\n", DEBUG)
	case LOCAL:
		fmt.Printf("level : %s\n", LOCAL)
	case PROD:
		fmt.Printf("level : %s\n", PROD)
	}
}
