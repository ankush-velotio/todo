package db

import (
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"todo/internal/models"
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

func (c *PostgreSQLRepository) Create(model, value interface{}) error {
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
func (c *PostgreSQLRepository) CreateTodo(model, value interface{}) error {
	connection := c.ConnectDB()
	defer func(conn *gorm.DB) {
		err := c.CloseDB(conn)
		if err != nil {
			log.Println("Postgres: cannot close current database")
		}
	}(connection)

	// Todo: If editors is not omitted then it will override the related user in database
	// If editors field is omitted then it will not create todo and editors relationship
	res := connection.Model(model).Omit("Editors").Create(value)
	return res.Error
}

func (c *PostgreSQLRepository) FindTodo(model interface{}) interface{} {
	connection := c.ConnectDB()
	defer func(conn *gorm.DB) {
		err := c.CloseDB(conn)
		if err != nil {
			log.Println("Postgres: cannot close current database")
		}
	}(connection)

	var rec []models.Todo

	connection.Model(model).Preload("Editors").Preload("Owner").Find(&rec)

	return rec
}

func (c *PostgreSQLRepository) Where(query, model interface{}, args ...interface{}) interface{} {
	connection := c.ConnectDB()
	defer func(conn *gorm.DB) {
		err := c.CloseDB(conn)
		if err != nil {
			log.Println("Postgres: cannot close current database")
		}
	}(connection)

	connection.Model(model).Where(query, args...).First(model)
	return model
}
