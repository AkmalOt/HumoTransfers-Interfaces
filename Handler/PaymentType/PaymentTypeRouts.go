package PaymentType

import "github.com/gin-gonic/gin"

func (h *PaymentType) PaymentGenRouting(router *gin.RouterGroup) {
	router.POST("/add_payment_type", h.AddPaymentType)
	router.GET("/get_payment_type", h.GetPaymentType)
	router.PUT("/update_payment_type", h.UpdatePaymentType)
}
