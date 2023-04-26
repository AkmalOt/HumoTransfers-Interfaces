package SysMessage

import (
	"awesomeProject2/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type SysMessage struct {
	repo       SysMessageInterface
	pagination Paginate
}

func New(r SysMessageInterface, p Paginate) *SysMessage {
	return &SysMessage{
		repo:       r,
		pagination: p,
	}
}

type Paginate interface {
	GeneratePaginationFromRequest(ctx *gin.Context) models.Pagination
}

type SysMessageInterface interface {
	AddSysMessage(message *models.SysMessage) error
	GetSysMessage(pagination *models.Pagination) ([]models.SysMessage, error)
	UpdateSysMessage(message *models.SysMessage) error
	DeleteSysMessage(message *models.SysMessage) error
	SysMessageStatus(message *models.SysMessage) error
	TotalPageSysMessage(limit int64) (int64, error)
}

func (h *SysMessage) AddSysMessage(ctx *gin.Context) {
	var message models.SysMessage

	if err := ctx.ShouldBindJSON(&message); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	message.Active = true
	err := h.repo.AddSysMessage(&message)
	if err != nil {
		log.Printf("%s in AddSysMessage(server)", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, "SysMessage added!")
}

func (h *SysMessage) GetSysMessage(ctx *gin.Context) {
	pagination := h.pagination.GeneratePaginationFromRequest(ctx)
	SysMessageLists, err := h.repo.GetSysMessage(&pagination)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	TotalPages, err := h.repo.TotalPageSysMessage(int64(pagination.Limit))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	pagination.Records = SysMessageLists
	pagination.TotalPages = TotalPages
	ctx.JSON(http.StatusOK, pagination)
}

func (h SysMessage) UpdateSysMessage(ctx *gin.Context) {
	var message *models.SysMessage

	if err := ctx.ShouldBindJSON(&message); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if message.Name == "" {
		err := h.repo.DeleteSysMessage(message)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			log.Println(err)
			return
		}
	} else {
		err := h.repo.UpdateSysMessage(message)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			log.Println(err)
			return
		}
	}

	ctx.JSON(http.StatusOK, " Done!")
}

func (h SysMessage) SysMessageStatus(ctx *gin.Context) {
	var message *models.SysMessage

	if err := ctx.ShouldBindJSON(&message); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err := h.repo.SysMessageStatus(message)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, "Done!")
}
