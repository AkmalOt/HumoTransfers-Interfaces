package Language

import "github.com/gin-gonic/gin"

func (h *Language) LanguageGenRouting(router *gin.RouterGroup) {
	router.POST("/add_language", h.AddLanguage)
	router.GET("/get_language", h.GetLanguage)
	router.PUT("/update_language", h.UpdateLanguage)
	router.PATCH("/status_language", h.LanguageStatus)
}
