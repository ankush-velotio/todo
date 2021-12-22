package db

import "github.com/jinzhu/gorm"

type IDBRepository interface {
	ConnectDB() *gorm.DB
	CloseDB(connection *gorm.DB) error
	Create(model, value interface{}) error
	CreateTodo(model, value interface{}) error
	FindTodo(model interface{}) interface{}
	Where(query, model interface{}, args ...interface{}) interface{}
}
