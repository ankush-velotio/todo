package connect_db

import (
	"os"
	"todo/internal/db"
)

type DB struct {
	IDBRepository db.IDBRepository
}

func (db DB) ConnectDB() interface{} {
	return db.IDBRepository.ConnectDB()
}

func (db DB) CloseDB(conn interface{}) error {
	return db.IDBRepository.CloseDB(conn)
}

func (db DB) Create(model interface{}, value interface{}) error {
	return db.IDBRepository.Create(model, value)
}

func (db DB) FindUser(value interface{}) interface{} {
	return db.IDBRepository.FindUser(value)
}

func (db DB) GetAllTodo() interface{} {
	return db.IDBRepository.GetAllTodo()
}

var DBConn = DB{IDBRepository: &db.NPostgreSQLRepository{DatabaseDialect: "postgres",
	DatabaseURL: os.Getenv("POSTGRES_URL")}}
