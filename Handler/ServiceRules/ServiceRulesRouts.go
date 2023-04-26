package ServiceRules

import (
	"awesomeProject2/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type ServiceRules struct {
	repo       ServiceRulesInterface
	pagination Pagination
}

func New(r ServiceRulesInterface, p Pagination) *ServiceRules {
	return &ServiceRules{
		repo:       r,
		pagination: p,
	}
}

type Pagination interface {
	GeneratePaginationFromRequest(ctx *gin.Context) models.Pagination
}

type ServiceRulesInterface interface {
	AddServiceRules(ServiceRules *models.ServicesRules) error
	GetServiceRules(pagination *models.Pagination) ([]models.ServicesRules, error)
	UpdateServiceRules(serviceRules *models.ServicesRules) error
	DeleteServiceRules(ServiceRules *models.ServicesRules) error
	TotalPageServiceRules(limit int64) (int64, error)
}

func (h *ServiceRules) AddServiceRules(ctx *gin.Context) {
	var ServiceRules models.ServicesRules

	if err := ctx.ShouldBindJSON(&ServiceRules); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err := h.repo.AddServiceRules(&ServiceRules)
	if err != nil {
		log.Printf("%s in AddServiceRules(server)", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, "ServiceRules added!")
}

func (h *ServiceRules) GetServiceRules(ctx *gin.Context) {
	pagination := h.pagination.GeneratePaginationFromRequest(ctx)
	ServiceRulesList, err := h.repo.GetServiceRules(&pagination)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	TotalPages, err := h.repo.TotalPageServiceRules(int64(pagination.Limit))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	pagination.Records = ServiceRulesList
	pagination.TotalPages = TotalPages
	ctx.JSON(http.StatusOK, pagination)
}

func (h ServiceRules) UpdateServiceRules(ctx *gin.Context) {
	var ServiceRules *models.ServicesRules

	if err := ctx.ShouldBindJSON(&ServiceRules); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if ServiceRules.Name == "" && ServiceRules.Type == "" {
		err := h.repo.DeleteServiceRules(ServiceRules)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			log.Println(err)
			return
		}
	} else {
		err := h.repo.UpdateServiceRules(ServiceRules)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			log.Println(err)
			return
		}
	}
	ctx.JSON(http.StatusOK, " Done!")
}
