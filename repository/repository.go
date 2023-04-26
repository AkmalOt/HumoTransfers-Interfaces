package repository

import (
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
)

type Repository struct {
	AccAgent.AccAgentRepo
	AgentsTable.AgentsTableRepo
	Country.CountryRepo
	CurrencyTable.CurrencyTableRepo
	Language.LanguageRepo
	PaymentType.PaymentTypeRepo
	ServCountry.ServCountryRepo
	ServiceRules.ServiceRulesRepo
	ServiceTable.ServiceTableRepo
	SysMessage.SysMessageRepo
	TransferTable.TransferTableRepo
	UserInfoTable.UserInfoRepo
	VendorTable.VendorRepo
}
