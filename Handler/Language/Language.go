package Language

import (
	"awesomeProject2/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Language struct {
	repo       LanguageInterface
	pagination Paginate
}

func New(r LanguageInterface, p Paginate) *Language {
	return &Language{
		repo:       r,
		pagination: p,
	}
}

type Paginate interface {
	GeneratePaginationFromRequest(ctx *gin.Context) models.Pagination
}

type LanguageInterface interface {
	AddLanguage(language *models.Languages) error
	GetLanguages(pagination *models.Pagination) ([]models.Languages, error)
	UpdateLanguage(language *models.Languages) error
	DeleteLanguage(language *models.Languages) error
	LanguageStatus(language *models.Languages) error
	TotalPageLanguage(limit int64) (int64, error)
}

func (h *Language) AddLanguage(ctx *gin.Context) {
	var language models.Languages

	if err := ctx.ShouldBindJSON(&language); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	language.Active = true
	err := h.repo.AddLanguage(&language)
	if err != nil {
		log.Printf("%s in AddLanguage(server)", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, "Language added!")
}

func (h *Language) GetLanguage(ctx *gin.Context) {
	pagination := h.pagination.GeneratePaginationFromRequest(ctx)
	LanguageLists, err := h.repo.GetLanguages(&pagination)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	TotalPages, err := h.repo.TotalPageLanguage(int64(pagination.Limit))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	pagination.Records = LanguageLists
	pagination.TotalPages = TotalPages
	ctx.JSON(http.StatusOK, pagination)
}

func (h Language) UpdateLanguage(ctx *gin.Context) {
	var language *models.Languages

	if err := ctx.ShouldBindJSON(&language); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if language.Name == "" && language.Icon == "" {
		err := h.repo.DeleteLanguage(language)
		if err != nil {
			if err.Error() == "ERROR: update or delete on table \"languages\" violates foreign key constraint \"transfers_lang_id_fkey\" on table \"transfers\" (SQLSTATE 23503)" {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": " ERROR: update or delete on table 'languages' violates foreign key constraint 'transfers_lang_id_fkey' on table 'transfers' (SQLSTATE 23503)"})
				return
			} else {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
				log.Println(err)
				return
			}
		}
	} else {
		err := h.repo.UpdateLanguage(language)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			log.Println(err)
			return
		}
	}
	ctx.JSON(http.StatusOK, " Done!")
}

func (h Language) LanguageStatus(ctx *gin.Context) {
	var language *models.Languages

	if err := ctx.ShouldBindJSON(&language); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err := h.repo.LanguageStatus(language)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, "Done!")
}
