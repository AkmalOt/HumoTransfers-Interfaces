package AccAgent

import (
	"gorm.io/gorm"
)

type AccAgentRepo struct {
	Connection *gorm.DB
}
