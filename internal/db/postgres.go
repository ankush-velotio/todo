package db

import (
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

type PostgreSQLRepository struct {
	DatabaseDialect string
	DatabaseURL     string
}

func (c *PostgreSQLRepository) ConnectDB() *gorm.DB {
	connection, err := gorm.Open(c.DatabaseDialect, c.DatabaseURL)

	if err != nil {
		panic("failed to connect database")
	}

	if err = connection.DB().Ping(); err != nil {
		log.Fatalln("failed to ping database")
	}

	log.Println("connected to database")
	return connection
}

func (c *PostgreSQLRepository) CloseDB(connection *gorm.DB) error {
	if err := connection.Close(); err != nil {
		return errors.New("cannot close current database")
	}
	log.Println("database closed successfully")
	return nil
}

func (c *PostgreSQLRepository) Create(model interface{}, value interface{}) error {
	connection := c.ConnectDB()
	defer func(conn *gorm.DB) {
		err := c.CloseDB(conn)
		if err != nil {
			log.Println("Postgres: cannot close current database")
		}
	}(connection)

	res := connection.Model(model).Create(value)
	return res.Error
}
