package connect_db

import (
	"github.com/jinzhu/gorm"
	"os"
	"todo/internal/db"
)

type DB struct {
	IDBRepository db.IDBRepository
}

func (db DB) ConnectDB() *gorm.DB {
	return db.IDBRepository.ConnectDB()
}

func (db DB) CloseDB(conn *gorm.DB) error {
	return db.IDBRepository.CloseDB(conn)
}

func (db DB) Create(model interface{}, value interface{}) error {
	return db.IDBRepository.Create(model, value)
}

var DBConn = DB{IDBRepository: &db.PostgreSQLRepository{DatabaseDialect: "postgres",
	DatabaseURL: os.Getenv("POSTGRES_URL")}}
