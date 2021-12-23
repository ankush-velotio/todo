package db

type IDBRepository interface {
	ConnectDB() interface{}
	CloseDB(connection interface{}) error
	Create(model interface{}, value interface{}) error
	FindUser(value interface{}) interface{}
	GetAllTodo() interface{}
}
