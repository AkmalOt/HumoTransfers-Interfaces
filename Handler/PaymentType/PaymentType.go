package PaymentType

import (
	"awesomeProject2/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type PaymentType struct {
	repo       PaymentTypeInterface
	pagination Paginate
}

func New(r PaymentTypeInterface, p Paginate) *PaymentType {
	return &PaymentType{
		repo:       r,
		pagination: p,
	}
}

type Paginate interface {
	GeneratePaginationFromRequest(ctx *gin.Context) models.Pagination
}

type PaymentTypeInterface interface {
	AddPaymentType(paymentType *models.PaymentType) error
	GetPaymentType(pagination *models.Pagination) ([]models.PaymentType, error)
	UpdatePaymentType(PaymentType *models.PaymentType) error
	DeletePaymentType(PaymentType *models.PaymentType) error
	TotalPagePaymentType(limit int64) (int64, error)
}

func (h *PaymentType) AddPaymentType(ctx *gin.Context) {
	var PaymentType models.PaymentType

	if err := ctx.ShouldBindJSON(&PaymentType); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err := h.repo.AddPaymentType(&PaymentType)
	if err != nil {
		log.Printf("%s in AddPaymentType(server)", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, "PaymentType added!")
}

func (h *PaymentType) GetPaymentType(ctx *gin.Context) {
	pagination := h.pagination.GeneratePaginationFromRequest(ctx)
	PAymentTypeList, err := h.repo.GetPaymentType(&pagination)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	TotalPages, err := h.repo.TotalPagePaymentType(int64(pagination.Limit))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	pagination.Records = PAymentTypeList
	pagination.TotalPages = TotalPages
	ctx.JSON(http.StatusOK, pagination)
}

func (h PaymentType) UpdatePaymentType(ctx *gin.Context) {
	var paymentType *models.PaymentType

	if err := ctx.ShouldBindJSON(&paymentType); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if paymentType.Name == "" {
		err := h.repo.DeletePaymentType(paymentType)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			log.Println(err)
			return
		}
	} else {
		err := h.repo.UpdatePaymentType(paymentType)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			log.Println(err)
			return
		}
	}
	ctx.JSON(http.StatusOK, " Done!")
}
