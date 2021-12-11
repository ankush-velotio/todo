package db

import "github.com/jinzhu/gorm"

type IConnectDBStrategy interface {
	ConnectDB() *gorm.DB
	CloseDB(connection *gorm.DB) error
}
