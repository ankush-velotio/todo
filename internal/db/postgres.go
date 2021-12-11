package db

import (
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"os"
)

var databaseURL = os.Getenv("POSTGRES_URL")

type ConnectPostgresSQLStrategy struct {
	DatabaseDialect string
}

func (c *ConnectPostgresSQLStrategy) ConnectDB() *gorm.DB {
	connection, err := gorm.Open(c.DatabaseDialect, databaseURL)

	if err != nil {
		panic("failed to connect database")
	}

	if err = connection.DB().Ping(); err != nil {
		log.Fatalln("failed to ping database")
	}

	log.Println("connected to database")
	return connection
}

func (c *ConnectPostgresSQLStrategy) CloseDB(connection *gorm.DB) error {
	if err := connection.Close(); err != nil {
		return errors.New("cannot close current database")
	}
	log.Println("database closed successfully")
	return nil
}
