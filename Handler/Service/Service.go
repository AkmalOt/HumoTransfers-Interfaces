package Service

import (
	"awesomeProject2/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Service struct {
	repo       ServiceInterface
	pagination Paginate
}

func New(r ServiceInterface, p Paginate) *Service {
	return &Service{
		repo:       r,
		pagination: p,
	}
}

type Paginate interface {
	GeneratePaginationFromRequest(ctx *gin.Context) models.Pagination
}

type ServiceInterface interface {
	AddService(account *models.Services) error
	GetService(pagination *models.Pagination) ([]models.Services, error)
	UpdateService(Service *models.Services) error
	DeleteService(Service *models.Services) error
	ServicesStatus(Service *models.Services) error
	TotalPageServices(limit int64) (int64, error)
}

func (h *Service) AddServices(ctx *gin.Context) {
	var service models.Services

	if err := ctx.ShouldBindJSON(&service); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	service.Active = true
	err := h.repo.AddService(&service)
	if err != nil {
		log.Printf("%s in AddService(server)", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, "AddService added!")
}

func (h *Service) GetServices(ctx *gin.Context) {
	pagination := h.pagination.GeneratePaginationFromRequest(ctx)

	ServicesList, err := h.repo.GetService(&pagination)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	TotalPages, err := h.repo.TotalPageServices(int64(pagination.Limit))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	pagination.Records = ServicesList
	pagination.TotalPages = TotalPages
	ctx.JSON(http.StatusOK, pagination)
}

func (h Service) UpdateService(ctx *gin.Context) {
	var service *models.Services

	if err := ctx.ShouldBindJSON(&service); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if service.VendorId == 0 && service.Name == "" &&
		service.Type == "" {
		err := h.repo.DeleteService(service)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
	} else {
		err := h.repo.UpdateService(service)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
	}
	ctx.JSON(http.StatusOK, " Done!")
}

func (h Service) ServiceStatus(ctx *gin.Context) {
	var services *models.Services

	if err := ctx.ShouldBindJSON(&services); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err := h.repo.ServicesStatus(services)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, "Done!")
}
