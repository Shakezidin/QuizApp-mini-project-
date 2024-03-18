package main

import (
	"log"

	"github.com/Shakezidin/config"
	admin "github.com/Shakezidin/pkg/admin/di"
	sr "github.com/Shakezidin/pkg/server"
	user "github.com/Shakezidin/pkg/user/di"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Printf("Error Loading Config Files, error: %v", err)
	}

	engine := sr.Server()
	admin.Init(engine.R, config)
	user.Init(engine.R, config)

	engine.StartServer(config.SERVERPORT)
}
