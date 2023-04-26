package AgentsTable

import (
	"gorm.io/gorm"
)

type AgentsTableRepo struct {
	Connection *gorm.DB
}
