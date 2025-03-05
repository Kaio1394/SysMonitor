package handlers

import (
	"SysMonitor/internal/models"
	"SysMonitor/internal/services"
	"SysMonitor/logger"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AgentHandlerImpl struct {
	s *services.AgentServiceImpl
}

func NewAgentHandlerImpl(s *services.AgentServiceImpl) *AgentHandlerImpl {
	return &AgentHandlerImpl{s: s}
}

func (h *AgentHandlerImpl) CreateAgent(c *gin.Context) {
	var a models.Agent
	if err := c.ShouldBindJSON(&a); err != nil {
		logger.Log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	agent, err := h.s.CreateAgent(context.Background(), a)
	if err != nil {
		logger.Log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"agent": agent,
	})
}

func (h *AgentHandlerImpl) GetAgents(c *gin.Context) {
	logger.Log.Info("Request received")
	logger.Log.Info("Method GET")
	agents, err := h.s.GetAgents(context.Background())
	if err != nil {
		logger.Log.Error("error: " + err.Error())

		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"agents": agents,
	})
	agentsJSON, err := json.Marshal(agents)
	if err != nil {
		logger.Log.Error("error: " + err.Error())
		return
	}
	logger.Log.Info("Response agents: " + string(agentsJSON))
}
