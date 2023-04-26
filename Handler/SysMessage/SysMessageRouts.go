package SysMessage

import "github.com/gin-gonic/gin"

func (h *SysMessage) SysMessageGenRouting(router *gin.RouterGroup) {
	router.POST("/add_sys_message", h.AddSysMessage)
	router.GET("/get_sys_message", h.GetSysMessage)
	router.PUT("/update_sys_message", h.UpdateSysMessage)
	router.PATCH("/status_sys_message", h.SysMessageStatus)
}
