package ServiceRules

import "github.com/gin-gonic/gin"

func (h *ServiceRules) ServiceRulesGenRouting(router *gin.RouterGroup) {
	router.POST("/add_serv_rules", h.AddServiceRules)
	router.GET("/get_serv_rules", h.GetServiceRules)
	router.PUT("/update_serv_rules", h.UpdateServiceRules)
}
