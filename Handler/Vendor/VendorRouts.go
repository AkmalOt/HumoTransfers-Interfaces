package Vendor

import "github.com/gin-gonic/gin"

func (h *Vendor) VendorGenRouting(router *gin.RouterGroup) {
	router.POST("/add_vendor", h.AddVendor)
	router.GET("/get_vendor", h.GetVendor)
	router.PUT("/update_vendor", h.UpdateVendor)
	router.PATCH("/status_vendor", h.VendorStatus)
}
