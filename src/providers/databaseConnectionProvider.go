package providers

import (
	"database/sql"
	"fmt"

	"github.com/devices/src/logger"
	_ "github.com/lib/pq"
)

type DatabaseConfiguration struct {
	User             string
	Password         string
	Name             string
	ConnectionString string
}

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "mysecretpassword"
	DB_NAME     = "devices"
)

func GetDataBaseConnection(logger logger.Logger, databaseConfig DatabaseConfiguration) *sql.DB {
	dbConfig := getConfigs(databaseConfig)

	connectionString := getConnectionString(dbConfig, true)
	db, _ := sql.Open("postgres", connectionString)

	_, err := db.Query("CREATE TABLE IF NOT EXISTS devices (id SERIAL PRIMARY KEY, name VARCHAR, brand VARCHAR, created VARCHAR);")

	if err != nil {
		logger.LogWarning("Creating devices database", "DatabaseConnector")

		connectionString := getConnectionString(databaseConfig, false)

		db, _ := sql.Open("postgres", connectionString)

		_, err = db.Exec("CREATE DATABASE devices;")

		if err != nil {
			logger.LogError(err.Error(), "DatabaseConnector")
			return nil
		}

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

func getConnectionString(databaseConfig DatabaseConfiguration, addDatabaseName bool) string {
	connectionString := ""

	if databaseConfig.ConnectionString == "" {
		connectionString = fmt.Sprintf("user=%s password=%s sslmode=disable", databaseConfig.User, databaseConfig.Password)

		if addDatabaseName {
			connectionString = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", databaseConfig.User, databaseConfig.Password, databaseConfig.Name)
		}
	} else {
		connectionString = fmt.Sprintf("%s?sslmode=disable", databaseConfig.ConnectionString)

		if addDatabaseName {
			connectionString = fmt.Sprintf("%s/%s?sslmode=disable", databaseConfig.ConnectionString, databaseConfig.Name)
		}
	}

	return connectionString
}
