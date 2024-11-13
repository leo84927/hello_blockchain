package connection

import (
	"gorm.io/gorm"
)

type GormInterface interface {
	Client() *gorm.DB
	Close()
}

func GetGormClient() *gorm.DB {
	return _gormInterface.Client()
}
