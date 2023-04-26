package UserInfo

import "github.com/gin-gonic/gin"

func (h *UserInfo) UserInfoGenRouting(router *gin.RouterGroup) {
	router.POST("/add_user", h.AddUserInfo)
	router.GET("/get_user", h.GetUserInfo)
	router.PUT("/update_user", h.UpdateUserInfo)
	router.PATCH("/status_user", h.UserInfoStatus)
}
