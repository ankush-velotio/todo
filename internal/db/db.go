package db

import "github.com/jinzhu/gorm"

type IDBRepository interface {
	ConnectDB() *gorm.DB
	CloseDB(connection *gorm.DB) error
	Create(model interface{}, value interface{}) error
}
