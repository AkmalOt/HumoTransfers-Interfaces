package SysMessage

import "gorm.io/gorm"

type SysMessageRepo struct {
	Connection *gorm.DB
}
