package providers

import (
	"database/sql"
	"fmt"

	"github.com/devices/src/logger"
	_ "github.com/lib/pq"
)

type DatabaseConfiguration struct {
	User     string
	Password string
	Name     string
}

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "mysecretpassword"
	DB_NAME     = "devices"
)

func GetDataBaseConnection(logger logger.Logger, databaseConfig DatabaseConfiguration) *sql.DB {
	dbConfig := getConfigs(databaseConfig)

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbConfig.User, dbConfig.Password, dbConfig.Name)
	db, _ := sql.Open("postgres", dbinfo)

	_, err := db.Query("CREATE TABLE IF NOT EXISTS devices (id SERIAL PRIMARY KEY, name VARCHAR, brand VARCHAR, created VARCHAR);")

	if err != nil {
		logger.LogWarning("Creating devices database", "DatabaseConnector")

		dbinfo := fmt.Sprintf("user=%s password=%s sslmode=disable", dbConfig.User, dbConfig.Password)
		db, _ := sql.Open("postgres", dbinfo)

		db.Exec("CREATE DATABASE devices;")
		db.Close()
		GetDataBaseConnection(logger, databaseConfig)
	}

	logger.LogInfo("Connection to database established", "DatabaseConnector")

	return db
}

func getConfigs(databaseConfig DatabaseConfiguration) DatabaseConfiguration {
	if databaseConfig.User == "" {
		databaseConfig.User = DB_USER
	}

	if databaseConfig.Password == "" {
		databaseConfig.Password = DB_PASSWORD
	}

	if databaseConfig.Name == "" {
		databaseConfig.Name = DB_NAME
	}

	return databaseConfig
}
