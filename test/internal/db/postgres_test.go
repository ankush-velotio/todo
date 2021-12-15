package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	testify "github.com/stretchr/testify/assert"
	"gotest.tools/v3/assert"
	"log"
	"os"
	"testing"
	"time"
	"todo/internal/db"
)

var databaseURL = os.Getenv("POSTGRES_URL")
var postgresUser = os.Getenv("POSTGRES_USER")
var postgresPassword = os.Getenv("POSTGRES_PASSWORD")
var postgresDBName = os.Getenv("POSTGRES_DB")
var port = "5432"

var dbConn *gorm.DB

func TestMain(m *testing.M) {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "14",
		Env: []string{
			"POSTGRES_PASSWORD=" + postgresPassword,
			"POSTGRES_USER=" + postgresUser,
			"POSTGRES_DB=" + postgresDBName,
			"listen_addresses = '*'",
		},
		ExposedPorts: []string{"5432"},
		PortBindings: map[docker.Port][]docker.PortBinding{"5432": {
			{HostIP: "0.0.0.0", HostPort: port},
		}},
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	log.Println("Connecting to database on url: ", databaseURL)

	err = resource.Expire(120)
	if err != nil {
		log.Println(err)
	} // Tell docker to hard kill the container in 120 seconds

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	pool.MaxWait = 120 * time.Second
	if err = pool.Retry(func() error {
		dbConn, err = gorm.Open("postgres", databaseURL)
		if err != nil {
			return err
		}
		return dbConn.DB().Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	//Run tests
	code := m.Run()

	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func TestConnectDB(t *testing.T) {
	// Provide valid database dialect and database URL
	dbURL := databaseURL
	postgres := db.PostgreSQLRepository{DatabaseDialect: "postgres", DatabaseURL: dbURL}

	pgConn := postgres.ConnectDB()
	assert.Assert(t, pgConn != nil)

	// Provide invalid database URL
	// database named nodatabase does not exist
	dbURL = "postgres://postgres:pass@localhost:5432/nodatabase?sslmode=disable"
	postgres = db.PostgreSQLRepository{DatabaseDialect: "postgres", DatabaseURL: dbURL}
	// This code will panic as the database URL is not correct
	panics := testify.Panics(t, func() { postgres.ConnectDB() })
	assert.Assert(t, panics)
}

func TestCloseDB(t *testing.T) {
	dbURL := databaseURL
	postgres := db.PostgreSQLRepository{DatabaseDialect: "postgres", DatabaseURL: dbURL}
	pgConn := postgres.ConnectDB()

	// If database is closed successfully then nil error should be returned
	assert.NilError(t, postgres.CloseDB(pgConn))
}
