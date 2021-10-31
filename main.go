package main

import (
	"fmt"
	"os"

	"github.com/devices/src/logger"
	"github.com/devices/src/models"
	"github.com/devices/src/providers"
	"github.com/devices/src/services"
)

func main() {
	// read env
	databaseUser := os.Getenv("DB_USER")
	databasePassword := os.Getenv("DB_PASSWORD")
	databaseName := os.Getenv("DB_NAME")

	// logger
	loggerInstance := logger.NewLogger(os.Stdout, logger.Info)

	// connections
	database := providers.GetDataBaseConnection(loggerInstance, providers.DatabaseConfiguration{User: databaseUser, Password: databasePassword, Name: databaseName})
	database.Ping()

	// services
	deviceService := services.NewDeviceService(database, loggerInstance)

	id := deviceService.Add(models.Device{Name: "test", Brand: "test", Created: "created"})

	device := deviceService.GetById(id)

	fmt.Println(device)
}
