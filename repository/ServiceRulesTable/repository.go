package ServiceRules

import (
	"gorm.io/gorm"
)

type ServiceRulesRepo struct {
	Connection *gorm.DB
}
