package Vendor

import (
	"awesomeProject2/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Vendor struct {
	repo       VendorInterface
	pagination Paginate
}

func New(r VendorInterface, p Paginate) *Vendor {
	return &Vendor{
		repo:       r,
		pagination: p,
	}
}

type Paginate interface {
	GeneratePaginationFromRequest(ctx *gin.Context) models.Pagination
}

type VendorInterface interface {
	AddVendor(account *models.Vendor) error
	GetVendor(pagination *models.Pagination) ([]models.Vendor, error)
	UpdateVendor(Vendor *models.Vendor) error
	DeleteVendor(Vendor *models.Vendor) error
	VendorStatus(Vendor *models.Vendor) error
	TotalPageVendor(limit int64) (int64, error)
}

func (h *Vendor) AddVendor(ctx *gin.Context) {
	var Vendor models.Vendor

	if err := ctx.ShouldBindJSON(&Vendor); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	Vendor.Active = true
	err := h.repo.AddVendor(&Vendor)
	if err != nil {
		log.Printf("%s in AddVendor(server)", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, "Vendor added!")
}

func (h *Vendor) GetVendor(ctx *gin.Context) {
	pagination := h.pagination.GeneratePaginationFromRequest(ctx)
	VendorList, err := h.repo.GetVendor(&pagination)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	TotalPages, err := h.repo.TotalPageVendor(int64(pagination.Limit))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	pagination.Records = VendorList
	pagination.TotalPages = TotalPages
	ctx.JSON(http.StatusOK, pagination)
}

func (h Vendor) UpdateVendor(ctx *gin.Context) {
	var vendor *models.Vendor

	if err := ctx.ShouldBindJSON(&vendor); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if vendor.Name == "" {
		err := h.repo.DeleteVendor(vendor)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			log.Println(err)
			return
		}
	} else {
		err := h.repo.UpdateVendor(vendor)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			log.Println(err)
			return
		}
	}
	ctx.JSON(http.StatusOK, " Done!")
}

func (h Vendor) VendorStatus(ctx *gin.Context) {
	var vendor *models.Vendor

	if err := ctx.ShouldBindJSON(&vendor); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err := h.repo.VendorStatus(vendor)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		log.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, "Done!")
}
