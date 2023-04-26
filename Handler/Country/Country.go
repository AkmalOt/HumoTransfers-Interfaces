package Country

import (
	"awesomeProject2/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Country struct {
	repo       CountryInterface
	pagination Paginate
}

func New(r CountryInterface, p Paginate) *Country {
	return &Country{
		repo:       r,
		pagination: p,
	}
}

type Paginate interface {
	GeneratePaginationFromRequest(ctx *gin.Context) models.Pagination
}

type CountryInterface interface {
	AddCountry(country *models.Countries) error
	GetCountries(pagination *models.Pagination) ([]models.Countries, error)
	UpdateCountries(country *models.Countries) error
	DeleteCountries(country *models.Countries) error
	CountryStatus(country *models.Countries) error
	TotalPageCountry(limit int64) (int64, error)
}

func (h *Country) addCountry(ctx *gin.Context) {
	var country models.Countries

	if err := ctx.ShouldBindJSON(&country); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	country.Active = true

	err := h.repo.AddCountry(&country)
	if err != nil {
		log.Printf("%s in AddCountry", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, "Country added!")
}

func (h *Country) getCountry(ctx *gin.Context) {

	pagination := h.pagination.GeneratePaginationFromRequest(ctx)
	CountryLists, err := h.repo.GetCountries(&pagination)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	TotalPages, err := h.repo.TotalPageCountry(int64(pagination.Limit))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	pagination.Records = CountryLists
	pagination.TotalPages = TotalPages
	ctx.JSON(http.StatusOK, pagination)
}

func (h *Country) updateCountries(ctx *gin.Context) {
	var countries *models.Countries

	if err := ctx.ShouldBindJSON(&countries); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if countries.Icon == "" && countries.Name == "" {
		err := h.repo.DeleteCountries(countries)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			log.Println(err)
			return
		}
	} else {
		err := h.repo.UpdateCountries(countries)
		log.Println("work&")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			log.Println(err)
			return
		}
	}

	ctx.JSON(http.StatusOK, " Done!")
}

func (h Country) countryStatus(ctx *gin.Context) {
	var country *models.Countries

	if err := ctx.ShouldBindJSON(&country); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err := h.repo.CountryStatus(country)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		log.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, "Done!")
}
