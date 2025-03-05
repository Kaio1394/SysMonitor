package interfaces

import "github.com/gin-gonic/gin"

type AgentHandler interface {
	CreateAgent(c *gin.Context)
	GetAgents(c *gin.Context)
}
