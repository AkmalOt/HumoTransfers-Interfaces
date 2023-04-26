package TransferTable

import (
	"gorm.io/gorm"
)

type TransferTableRepo struct {
	Connection *gorm.DB
}
