package main

import (
	"github.com/IamVladlen/nmap-service/config"
	"github.com/IamVladlen/nmap-service/internal/app"
	"github.com/IamVladlen/nmap-service/pkg/logger"
)

func main() {
	cfg, err := config.New()

	log := logger.New(cfg.App.LogLevel)

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load configuration")
	}

	app.Run(cfg, log)
}
