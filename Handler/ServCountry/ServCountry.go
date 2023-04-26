package ServCountry

import (
	"awesomeProject2/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type ServCountry struct {
	repo       ServCountryInterface
	pagination Paginate
}

func New(r ServCountryInterface, p Paginate) *ServCountry {
	return &ServCountry{
		repo:       r,
		pagination: p,
	}
}

type Paginate interface {
	GeneratePaginationFromRequest(ctx *gin.Context) models.Pagination
}

type ServCountryInterface interface {
	AddServCountry(servCountry *models.ServicesCountry) error
	GetServCountry(pagination *models.Pagination) ([]models.ServicesCountry, error)
	DeleteServCountry(servCountry *models.ServicesCountry) error
	ServCountryStatus(servCountry *models.ServicesCountry) error
	TotalPageServCountry(limit int64) (int64, error)
}

func (h *ServCountry) AddServCountry(ctx *gin.Context) {
	var servicesCountry models.ServicesCountry

	if err := ctx.ShouldBindJSON(&servicesCountry); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	servicesCountry.Active = true
	err := h.repo.AddServCountry(&servicesCountry)
	if err != nil {
		log.Printf("%s in AddServCountry(server)", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, "AddServCountry added!")
}

func (h *ServCountry) GetServCountry(ctx *gin.Context) {
	pagination := h.pagination.GeneratePaginationFromRequest(ctx)

	ServCountryLists, err := h.repo.GetServCountry(&pagination)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	TotalPages, err := h.repo.TotalPageServCountry(int64(pagination.Limit))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	pagination.Records = ServCountryLists
	pagination.TotalPages = TotalPages
	ctx.JSON(http.StatusOK, pagination)
}

func (h ServCountry) DeleteServCountry(ctx *gin.Context) {
	var servCountry *models.ServicesCountry

	if err := ctx.ShouldBindJSON(&servCountry); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err := h.repo.DeleteServCountry(servCountry)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, " Done!")
}

func (h ServCountry) ServCountryStatus(ctx *gin.Context) {
	var servicesCountry *models.ServicesCountry

	if err := ctx.ShouldBindJSON(&servicesCountry); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err := h.repo.ServCountryStatus(servicesCountry)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, "Done!")
}
