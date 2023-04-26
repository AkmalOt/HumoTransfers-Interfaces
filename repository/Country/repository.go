package Country

import (
	"gorm.io/gorm"
)

type CountryRepo struct {
	Connection *gorm.DB
}
