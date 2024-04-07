package main

import (
	"log"

	"github.com/Shakezidin/config"
	"github.com/Shakezidin/pkg/server"
	"github.com/Shakezidin/pkg/user/di"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Printf("Error loading config files: %v", err)
		return
	}

	engine := server.Server()
	di.Init(engine.R, *cfg)

	engine.StartServer(cfg.SERVERPORT)
}
