package main

import (
	config "authenticationService/configs"
	"authenticationService/server"
	"fmt"
	"github.com/DerryRenaldy/logger/logger"
)

func main() {
	fmt.Println("Hello World")
	cfg := config.Cfg
	log := logger.New(cfg.App.AppName, cfg.App.Environment, cfg.App.LogLevel)

	svr := server.NewServer(cfg, log)

	svr.StartServer()
}
