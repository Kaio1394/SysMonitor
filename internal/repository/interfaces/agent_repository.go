package interfaces

import (
	"SysMonitor/internal/models"
	"context"
)

type IAgentRepository interface {
	CreateAgent(ctx context.Context, agent models.Agent) (models.Agent, error)
	GetAgents(ctx context.Context) ([]models.Agent, error)
}
