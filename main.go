package main

import (
	"awesomeProject2/Handler"
	"awesomeProject2/db"
	"awesomeProject2/repository"
	"awesomeProject2/repository/AccAgent"
	"awesomeProject2/repository/AgentsTable"
	"awesomeProject2/repository/Country"
	"awesomeProject2/repository/CurrencyTable"
	"awesomeProject2/repository/Language"
	PaymentType "awesomeProject2/repository/PaymentTypeTable"
	ServCountry "awesomeProject2/repository/ServCountryTable"
	ServiceRules "awesomeProject2/repository/ServiceRulesTable"
	"awesomeProject2/repository/ServiceTable"
	SysMessage "awesomeProject2/repository/SysMessageTable"
	TransferTable "awesomeProject2/repository/TransfersTable"
	"awesomeProject2/repository/UserInfoTable"
	"awesomeProject2/repository/VendorTable"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	route := gin.Default()
	DB := db.SetupGorm()
	if DB.Error != nil {
		return
	}

	repoLayer := &repository.Repository{
		AccAgentRepo:      AccAgent.AccAgentRepo{},
		AgentsTableRepo:   AgentsTable.AgentsTableRepo{},
		CountryRepo:       Country.CountryRepo{},
		CurrencyTableRepo: CurrencyTable.CurrencyTableRepo{},
		LanguageRepo:      Language.LanguageRepo{},
		PaymentTypeRepo:   PaymentType.PaymentTypeRepo{},
		ServCountryRepo:   ServCountry.ServCountryRepo{},
		ServiceRulesRepo:  ServiceRules.ServiceRulesRepo{},
		ServiceTableRepo:  ServiceTable.ServiceTableRepo{},
		SysMessageRepo:    SysMessage.SysMessageRepo{},
		TransferTableRepo: TransferTable.TransferTableRepo{},
		UserInfoRepo:      UserInfoTable.UserInfoRepo{},
		VendorRepo:        VendorTable.VendorRepo{},
	}
	newHandler := Handler.NewHandler(route, repoLayer)
	//newHandler := routs.NewHandler(route, handler)
	newHandler.Init()

	err := route.Run()
	if err != nil {
		log.Println(err)
		return
	}
}
