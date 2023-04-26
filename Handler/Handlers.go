package Handler

import (
	"awesomeProject2/Handler/AccAgent"
	"awesomeProject2/Handler/Agent"
	"awesomeProject2/Handler/Country"
	"awesomeProject2/Handler/Currency"
	"awesomeProject2/Handler/Language"
	"awesomeProject2/Handler/PaymentType"
	"awesomeProject2/Handler/ServCountry"
	"awesomeProject2/Handler/Service"
	"awesomeProject2/Handler/ServiceRules"
	"awesomeProject2/Handler/SysMessage"
	"awesomeProject2/Handler/Transfers"
	"awesomeProject2/Handler/UserInfo"
	"awesomeProject2/Handler/Vendor"
	"awesomeProject2/repository"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Engine     *gin.Engine
	Repository *repository.Repository
}

func NewHandler(engine *gin.Engine, repository *repository.Repository) *Handler {
	return &Handler{
		Engine:     engine,
		Repository: repository,
	}
}

func (h Handler) Init() {

	apiV1 := h.Engine.Group("/")
	AccAgent.New(h.Repository, h).AccAgentGenRouting(apiV1)
	Agent.New(h.Repository, h).AgentGenRouting(apiV1)
	Country.New(h.Repository, h).CountryGenRouting(apiV1)
	Currency.New(h.Repository, h).CurrencyGenRouting(apiV1)
	Language.New(h.Repository, h).LanguageGenRouting(apiV1)
	PaymentType.New(h.Repository, h).PaymentGenRouting(apiV1)
	ServCountry.New(h.Repository, h).ServCountryGenRouting(apiV1)
	Service.New(h.Repository, h).ServiceGenRouting(apiV1)
	ServiceRules.New(h.Repository, h).ServiceRulesGenRouting(apiV1)
	SysMessage.New(h.Repository, h).SysMessageGenRouting(apiV1)
	Transfers.New(h.Repository, h).TransfersGenRouting(apiV1)
	UserInfo.New(h.Repository, h).UserInfoGenRouting(apiV1)
	Vendor.New(h.Repository, h).VendorGenRouting(apiV1)

	h.Engine.POST("/upload", h.UploadImage)
}
