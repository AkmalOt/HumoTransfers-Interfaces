package AccAgent

import (
	"github.com/gin-gonic/gin"
)

func (h *AccAgent) AccAgentGenRouting(router *gin.RouterGroup) {
	router.POST("/add_acc_agents", h.AddAccount)
	router.GET("/get_acc_agents", h.GetAccountAgent)
	router.PUT("/update_account", h.UpdateAccountAgent)
	router.PATCH("/status_account", h.AccountAgentStatus)

}
