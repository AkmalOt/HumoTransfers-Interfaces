package Currency

import (
	"awesomeProject2/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Currency struct {
	repo       CurrencyInterface
	pagination Paginate
}

func New(r CurrencyInterface, p Paginate) *Currency {
	return &Currency{
		repo:       r,
		pagination: p,
	}
}

type Paginate interface {
	GeneratePaginationFromRequest(ctx *gin.Context) models.Pagination
}

type CurrencyInterface interface {
	AddCurrency(currency *models.Currency) error
	GetCurrency(pagination *models.Pagination) ([]models.Currency, error)
	UpdateCurrency(currency *models.Currency) error
	DeleteCurrency(currency *models.Currency) error
	TotalPageCurrency(limit int64) (int64, error)
}

func (h *Currency) AddCurrency(ctx *gin.Context) {
	var currency models.Currency

	if err := ctx.ShouldBindJSON(&currency); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err := h.repo.AddCurrency(&currency)
	if err != nil {
		log.Printf("%s in AddCurrency(server)", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, "Currency added!")
}

func (h *Currency) GetCurrency(ctx *gin.Context) {
	pagination := h.pagination.GeneratePaginationFromRequest(ctx)
	CurrencyLists, err := h.repo.GetCurrency(&pagination)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	TotalPages, err := h.repo.TotalPageCurrency(int64(pagination.Limit))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	pagination.Records = CurrencyLists
	pagination.TotalPages = TotalPages
	ctx.JSON(http.StatusOK, pagination)
}

func (h Currency) UpdateCurrency(ctx *gin.Context) {
	var currency *models.Currency

	if err := ctx.ShouldBindJSON(&currency); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if currency.Name == "" && currency.Icon == "" {
		err := h.repo.DeleteCurrency(currency)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			log.Println(err)
			return
		}
	} else {
		err := h.repo.UpdateCurrency(currency)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			log.Println(err)
			return
		}
	}
	ctx.JSON(http.StatusOK, " Done!")
}
