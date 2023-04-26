package Service

import "github.com/gin-gonic/gin"

func (h Service) ServiceGenRouting(router *gin.RouterGroup) {
	router.POST("/add_service", h.AddServices)
	router.GET("/get_services", h.GetServices)
	router.PUT("/update_service", h.UpdateService)
	router.PATCH("/status_service", h.ServiceStatus)

}