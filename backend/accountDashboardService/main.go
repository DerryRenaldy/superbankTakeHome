package main

import (
	config "accountDashboardService/configs"
	"accountDashboardService/server"

	"github.com/DerryRenaldy/logger/logger"
)

func main() {
	cfg := config.Cfg
	log := logger.New(cfg.App.AppName, cfg.App.Environment, cfg.App.LogLevel)

	svr := server.NewServer(cfg, log)

	svr.StartServer()
}
