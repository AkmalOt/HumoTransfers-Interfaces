package ServiceTable

import (
	"gorm.io/gorm"
)

type ServiceTableRepo struct {
	Connection *gorm.DB
}
