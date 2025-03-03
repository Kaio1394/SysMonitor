package routes

import (
	"SysMonitor/internal/handlers"
	"SysMonitor/internal/repository"
	"SysMonitor/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAgentRoutes(r *gin.Engine, db *gorm.DB) {
	repo := repository.NewAgentRepositoryImpl(db)
	service := services.NewAgentServiceImpl(repo)
	handler := handlers.NewAgentHandlerImpl(service)

	r.POST("/agent", handler.CreateAgent)
	r.GET("/agents", handler.GetAgents)
	//r.GET("/agent", handler.GetAgent)
	//r.DELETE("/agent", handler.DeleteAgent)
	//r.PUT("/agent", handler.UpdateAgent)
}
