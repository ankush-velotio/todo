package connect_db

import (
	"github.com/jinzhu/gorm"
	"os"
	"todo/internal/db"
)

type PostgreSQL struct {
	IDBRepository db.IDBRepository
}

func (db PostgreSQL) ConnectDB() *gorm.DB {
	return db.IDBRepository.ConnectDB()
}

func (db PostgreSQL) CloseDB(conn *gorm.DB) error {
	return db.IDBRepository.CloseDB(conn)
}

var PostgresConn = PostgreSQL{IDBRepository: &db.PostgreSQLRepository{DatabaseDialect: "postgres",
	DatabaseURL: os.Getenv("POSTGRES_URL")}}
