package main

import (
	"exchangeapp/config"
	"exchangeapp/router"
)

func main() {
	config.InitConfig()

	r := router.SetupRouter()
	listenAddr := config.AppConfig.App.Host + ":" + config.AppConfig.App.Port
	if config.AppConfig.App.Host == "" || config.AppConfig.App.Port == "" {
		listenAddr = "localhost:8080"
	}
	r.Run(listenAddr)
}
