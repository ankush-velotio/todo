package connect_db

import (
	"github.com/jinzhu/gorm"
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
