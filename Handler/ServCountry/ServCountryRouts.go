package ServCountry

import "github.com/gin-gonic/gin"

func (h ServCountry) ServCountryGenRouting(router *gin.RouterGroup) {
	router.POST("/add_serv_country", h.AddServCountry)
	router.GET("/get_serv_country", h.GetServCountry)
	router.PUT("/delete_serv_country", h.DeleteServCountry)
	router.PATCH("/status_serv_country", h.ServCountryStatus)
}
