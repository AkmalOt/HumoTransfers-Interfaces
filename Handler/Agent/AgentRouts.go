package Agent

import "github.com/gin-gonic/gin"

func (h *Agent) AgentGenRouting(router *gin.RouterGroup) {
	router.POST("/add_agents", h.AddAgent)
	router.GET("/get_agents", h.GetAgent)
	router.PUT("/update_agent", h.UpdateAgents)
	router.PATCH("/status_agent", h.AgentStatus)
}
