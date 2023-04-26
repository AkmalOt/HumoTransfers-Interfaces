package PaymentType

import (
	"gorm.io/gorm"
)

type PaymentTypeRepo struct {
	Connection *gorm.DB
}
