package UserInfo

import (
	"awesomeProject2/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UserInfo struct {
	repo       UserInfoInterface
	pagination Paginate
}

func New(r UserInfoInterface, p Paginate) *UserInfo {
	return &UserInfo{
		repo:       r,
		pagination: p,
	}
}

type Paginate interface {
	GeneratePaginationFromRequest(ctx *gin.Context) models.Pagination
}

type UserInfoInterface interface {
	AddUserInfo(user *models.UserInfo) error
	GetUserInfo(pagination *models.Pagination) ([]models.UserInfo, error)
	UpdateUserInfo(userInfo *models.UserInfo) error
	DeleteUserInfo(UserInfo *models.UserInfo) error
	UserInfoStatus(userInfo *models.UserInfo) error
	TotalPageUserInfo(limit int64) (int64, error)
}

func (h *UserInfo) AddUserInfo(ctx *gin.Context) {
	var UserInfo models.UserInfo

	if err := ctx.ShouldBindJSON(&UserInfo); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	UserInfo.Active = true
	err := h.repo.AddUserInfo(&UserInfo)
	if err != nil {
		log.Printf("%s in AddCountry", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, "Country added!")
}

func (h *UserInfo) GetUserInfo(ctx *gin.Context) {
	pagination := h.pagination.GeneratePaginationFromRequest(ctx)
	UserLists, err := h.repo.GetUserInfo(&pagination)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	TotalPages, err := h.repo.TotalPageUserInfo(int64(pagination.Limit))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	pagination.Records = UserLists
	pagination.TotalPages = TotalPages
	ctx.JSON(http.StatusOK, pagination)
}

func (h UserInfo) UpdateUserInfo(ctx *gin.Context) {
	var UserInfo *models.UserInfo

	if err := ctx.ShouldBindJSON(&UserInfo); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if UserInfo.Icon == "" && UserInfo.Name == "" && UserInfo.Sort == 0 {
		err := h.repo.DeleteUserInfo(UserInfo)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			log.Println(err)
			return
		}
	} else {
		err := h.repo.UpdateUserInfo(UserInfo)
		log.Println("work&")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			log.Println(err)
			return
		}
	}

	ctx.JSON(http.StatusOK, " Done!")
}

func (h UserInfo) UserInfoStatus(ctx *gin.Context) {
	var UserInfo *models.UserInfo

	if err := ctx.ShouldBindJSON(&UserInfo); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err := h.repo.UserInfoStatus(UserInfo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		log.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, "Done!")
}
