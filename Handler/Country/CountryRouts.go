package Country

import (
	"github.com/gin-gonic/gin"
)

func (h *Country) CountryGenRouting(router *gin.RouterGroup) {
	router.POST("/add_country", h.addCountry)
	router.GET("/get_country", h.getCountry)
	router.PUT("/update_countries", h.updateCountries)
	router.PATCH("/status_countries", h.countryStatus)
}
