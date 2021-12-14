package migration

import (
	"github.com/jinzhu/gorm"
	"log"
	db "todo/cmd/connect_db"
	"todo/internal/auth/user"
	"todo/internal/todos"
)

// Migrate function is used to create database table from the model
// AutoMigrate function does not work by providing the reference of struct
// Hence, Migrating the model to database using AutoMigrate function concretely
func Migrate(postgres db.PostgreSQL) {
	conn := postgres.ConnectDB()
	defer func(conn *gorm.DB) {
		err := postgres.CloseDB(conn)
		if err != nil {
			log.Println("Migrate DB: cannot close current database")
		}
	}(conn)
	conn.AutoMigrate(user.User{})
	conn.AutoMigrate(todos.Todo{})
	log.Println("Migrate DB: database migrated successfully")
}
