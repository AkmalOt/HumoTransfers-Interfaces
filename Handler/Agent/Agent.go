package Agent

import (
	"awesomeProject2/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Agent struct {
	repo       AgentInterface
	pagination Paginate
}

func New(r AgentInterface, p Paginate) *Agent {
	return &Agent{
		repo:       r,
		pagination: p,
	}
}

type Paginate interface {
	GeneratePaginationFromRequest(ctx *gin.Context) models.Pagination
}

type AgentInterface interface {
	AddAgent(agents *models.Agents) error
	GetAgent(pagination *models.Pagination) ([]models.Agents, error)
	UpdateAgents(agent *models.Agents) error
	DeleteAgents(agent *models.Agents) error
	AgentStatus(agent *models.Agents) error
	TotalPageAgents(limit int64) (int64, error)
}

func (h *Agent) AddAgent(ctx *gin.Context) {
	var agents models.Agents

	if err := ctx.ShouldBindJSON(&agents); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	agents.Active = true
	err := h.repo.AddAgent(&agents)
	if err != nil {
		log.Printf("%s in AddAgent(server)", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, "Agent added!")
}

func (h *Agent) GetAgent(ctx *gin.Context) {

	pagination := h.pagination.GeneratePaginationFromRequest(ctx)
	AgentLists, err := h.repo.GetAgent(&pagination)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	TotalPages, err := h.repo.TotalPageAgents(int64(pagination.Limit))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	pagination.Records = AgentLists
	pagination.TotalPages = TotalPages
	ctx.JSON(http.StatusOK, pagination)
}

func (h Agent) UpdateAgents(ctx *gin.Context) {
	var agents *models.Agents
	ctx.Param("id")
	if err := ctx.ShouldBindJSON(&agents); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if agents.Name == "" && agents.LegalName == "" {
		err := h.repo.DeleteAgents(agents)
		if err != nil {
			if err.Error() == "ERROR: update or delete on table \"agents\" violates foreign key constraint \"account_agents_agent_id_fkey\" on table \"account_agents\" (SQLSTATE 23503)" {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": " ERROR: update or delete on table 'agents' violates foreign key constraint 'account_agents_agent_id_fkey' on table 'account_agents' (SQLSTATE 23503)"})
				return
			} else {
				//if errors.Is(err, gorm.ErrDryRunModeUnsupported)
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
				log.Println(err)
				return
			}
		}
	} else {
		err := h.repo.UpdateAgents(agents)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			log.Println(err)
			return
		}
	}
	ctx.JSON(http.StatusOK, " Done!")
}

func (h Agent) AgentStatus(ctx *gin.Context) {
	var agent *models.Agents

	if err := ctx.ShouldBindJSON(&agent); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err := h.repo.AgentStatus(agent)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, "Done!")
}
