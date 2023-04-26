package UserInfoTable

import (
	"gorm.io/gorm"
)

type UserInfoRepo struct {
	Connection *gorm.DB
}
