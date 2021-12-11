package db

import (
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"strconv"
)

func ConnectDB(port int, dbName, dbUser, dbPassword, host string) *gorm.DB {
	databaseDialect := "postgres"
	databaseURL := "postgres://" + dbUser + ":" + dbPassword + "@" + host + ":" + strconv.Itoa(port) + "/" + dbName + "?sslmode=disable"
	connection, err := gorm.Open(databaseDialect, databaseURL)

	if err != nil {
		panic("failed to connect database")
	}

	log.Println("connected to database")
	return connection
}

func CloseDB(connection *gorm.DB) error {
	if err := connection.Close(); err != nil {
		return errors.New("cannot close current database")
	}
	log.Println("database closed successfully")
	return nil
}
