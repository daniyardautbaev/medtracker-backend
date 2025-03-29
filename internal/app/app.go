package app

import (
	"context"
	"log"
	"medtracker/medtracker/internal/app/config"
	"medtracker/medtracker/internal/app/connections"
	"medtracker/medtracker/internal/app/start"
)

func Run(configFile string) {
	ctx := context.Background()

	cfg, err := config.NewConfig(configFile)
	if err != nil {
		log.Fatalf("config error: %v", err)
	}

	conn, err := connections.NewConnections(cfg)
	if err != nil {
		log.Fatalf("connection error: %v", err)
	}
	defer conn.Close()

	// Здесь можно инициализировать репозитории/сервисы

	start.HTTP(ctx, cfg)
}
