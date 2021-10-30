package main

import (
	"os"

	"github.com/devices/src/logger"
	"github.com/devices/src/providers"
)

func main() {
	// read env
	databaseUser := os.Getenv("DB_USER")
	databasePassword := os.Getenv("DB_PASSWORD")
	databaseName := os.Getenv("DB_NAME")

	// instances
	loggerInstance := logger.NewLogger(os.Stdout, logger.Info)

	// connections
	database := providers.GetDataBaseConnection(loggerInstance, providers.DatabaseConfiguration{User: databaseUser, Password: databasePassword, Name: databaseName})
	database.Ping()
}
