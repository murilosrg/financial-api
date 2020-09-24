package main

import (
	"flag"

	"github.com/gin-gonic/gin"
	"github.com/murilosrg/financial-api/config"
	"github.com/murilosrg/financial-api/internal/routes"
)

func main() {
	load()
	listen()
}

func load() {
	var shouldInit = flag.Bool("init", false, "initialize all")
	flag.Parse()

	if *shouldInit {
		initAll()
	}
}

func listen() {
	gin.SetMode(config.Load().Mode)
	server := gin.Default()
	routes.SetupApiRouter(server)
	server.Run(config.Load().Address)
}
