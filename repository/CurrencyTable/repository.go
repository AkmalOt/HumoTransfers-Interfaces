package CurrencyTable

import (
	"gorm.io/gorm"
)

type CurrencyTableRepo struct {
	Connection *gorm.DB
}
