package Language

import (
	"gorm.io/gorm"
)

type LanguageRepo struct {
	Connection *gorm.DB
}
