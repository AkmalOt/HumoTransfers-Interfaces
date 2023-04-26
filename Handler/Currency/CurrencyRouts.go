package Currency

import "github.com/gin-gonic/gin"

func (h *Currency) CurrencyGenRouting(router *gin.RouterGroup) {
	router.POST("/add_currency", h.AddCurrency)
	router.GET("/get_currency", h.GetCurrency)
	router.PUT("/update_currency", h.UpdateCurrency)
}
