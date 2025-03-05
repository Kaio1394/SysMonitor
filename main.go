package main

import (
	"SysMonitor/config"
	"SysMonitor/database"
	"SysMonitor/internal/routes"
	"SysMonitor/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	logger.Log.Info("Starting SysMonitor")
	logger.Log.Info("Connecting to database...")

	db, err := database.ConnectDatabase()
	if err != nil {
		logger.Log.Error(err.Error())
		return
	}
	logger.Log.Info("Connect to database with success!")

	configs, err := config.ConfigSet()
	if err != nil {
		logger.Log.Error(err.Error())
		return
	}
	r := gin.Default()

	routes.RegisterAgentRoutes(r, db)
	logger.Log.Info("Running to port: " + configs.Server.Port)
	_ = r.Run(":" + configs.Server.Port)
}
