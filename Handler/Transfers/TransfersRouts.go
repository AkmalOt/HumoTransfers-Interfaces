package Transfers

import "github.com/gin-gonic/gin"

func (h Transfers) TransfersGenRouting(router *gin.RouterGroup) {
	router.POST("/add_transfer", h.AddTransfer)
	router.GET("/get_transfer", h.GetTransfer)
	router.PUT("/update_transfer", h.UpdateTransfer)
}
