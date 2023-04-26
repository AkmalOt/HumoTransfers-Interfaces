package Transfers

import (
	"awesomeProject2/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Transfers struct {
	repo       TransfersInterface
	pagination Paginate
}

func New(r TransfersInterface, p Paginate) *Transfers {
	return &Transfers{
		repo:       r,
		pagination: p,
	}
}

type Paginate interface {
	GeneratePaginationFromRequest(ctx *gin.Context) models.Pagination
}

type TransfersInterface interface {
	AddTransfer(transfer *models.Transfers) error
	GetTransfer(pagination *models.Pagination) ([]models.Transfers, error)
	UpdateTransfers(transfer *models.Transfers) error
	DeleteTransfers(transfers *models.Transfers) error
	TotalPageTransfer(limit int64) (int64, error)
}

func (h *Transfers) AddTransfer(ctx *gin.Context) {
	var transfers models.Transfers

	if err := ctx.ShouldBindJSON(&transfers); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err := h.repo.AddTransfer(&transfers)

	if err != nil {
		log.Printf("%s in AddTest(server)", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, "Transfer added!")
}

func (h *Transfers) GetTransfer(ctx *gin.Context) {
	pagination := h.pagination.GeneratePaginationFromRequest(ctx)
	TransferLists, err := h.repo.GetTransfer(&pagination)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	TotalPages, err := h.repo.TotalPageTransfer(int64(pagination.Limit))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	pagination.Records = TransferLists
	pagination.TotalPages = TotalPages
	ctx.JSON(http.StatusOK, pagination)
}

func (h Transfers) UpdateTransfer(ctx *gin.Context) {
	var transfer *models.Transfers

	if err := ctx.ShouldBindJSON(&transfer); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if transfer.EntityId == 0 && transfer.Entity == "" &&
		transfer.LangId == 0 && transfer.KeyField == 0 && transfer.Value == "" {
		err := h.repo.DeleteTransfers(transfer)
		if err != nil {

			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			log.Println(err)
			return
		}

	} else {
		err := h.repo.UpdateTransfers(transfer)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			log.Println(err)
			return
		}
	}
	ctx.JSON(http.StatusOK, " Done!")
}
