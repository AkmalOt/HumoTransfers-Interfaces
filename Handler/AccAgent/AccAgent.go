package AccAgent

import (
	"awesomeProject2/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type AccAgent struct {
	repo       AccAgentInterface
	pagination Paginate
}

func New(r AccAgentInterface, p Paginate) *AccAgent {
	return &AccAgent{
		repo:       r,
		pagination: p,
	}
}

type Paginate interface {
	GeneratePaginationFromRequest(ctx *gin.Context) models.Pagination
}

type AccAgentInterface interface {
	AddAccount(account *models.AccountAgent) error
	GetAccountAgent(pagination *models.Pagination) ([]models.AccountAgent, error)
	UpdateAccountAgent(AccAgent *models.AccountAgent) error
	DeleteAccountAgent(AccAgent *models.AccountAgent) error
	UpdateAccountDefault(AccAgent *models.AccountAgent) error
	AccountAgentStatus(AccAgent *models.AccountAgent) error
	TotalPageAccount(limit int64) (int64, error)
}

func (h *AccAgent) AddAccount(ctx *gin.Context) {
	var account models.AccountAgent

	if err := ctx.ShouldBindJSON(&account); err != nil {
		fmt.Printf("error in AddAccount- ShouldBindJSON - %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	account.Active = true
	err := h.repo.AddAccount(&account)
	if err != nil {
		log.Printf("%s in AddAccount(server)", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, "AccountAgent added!")
}

func (h *AccAgent) GetAccountAgent(ctx *gin.Context) {
	pagination := h.pagination.GeneratePaginationFromRequest(ctx)

	AccLists, err := h.repo.GetAccountAgent(&pagination)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	TotalPages, err := h.repo.TotalPageAccount(int64(pagination.Limit))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	pagination.Records = AccLists
	pagination.TotalPages = TotalPages
	ctx.JSON(http.StatusOK, pagination)
}

func (h AccAgent) UpdateAccountAgent(ctx *gin.Context) {
	var account *models.AccountAgent

	if err := ctx.ShouldBindJSON(&account); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if account.AgentId == 0 && account.CurrencyId == 0 &&
		account.Type == 0 {
		err := h.repo.DeleteAccountAgent(account)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
	} else {
		err := h.repo.UpdateAccountAgent(account)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
	}
	ctx.JSON(http.StatusOK, " Done!")
}

func (h AccAgent) UpdateAccountDefault(ctx *gin.Context) {
	var account *models.AccountAgent

	if err := ctx.ShouldBindJSON(&account); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	err := h.repo.UpdateAccountDefault(account)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, " Done!")
}

func (h AccAgent) AccountAgentStatus(ctx *gin.Context) {
	var account *models.AccountAgent

	if err := ctx.ShouldBindJSON(&account); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err := h.repo.AccountAgentStatus(account)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, "Done!")
}
