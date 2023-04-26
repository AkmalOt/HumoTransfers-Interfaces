package ServCountry

import (
	"gorm.io/gorm"
)

type ServCountryRepo struct {
	Connection *gorm.DB
}
