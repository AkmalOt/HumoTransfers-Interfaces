package VendorTable

import (
	"gorm.io/gorm"
)

type VendorRepo struct {
	Connection *gorm.DB
}
