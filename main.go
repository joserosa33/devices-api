package main

import (
	"log"
	"net/http"
	"os"

	"github.com/devices/src/controllers"
	"github.com/devices/src/handlers"
	"github.com/devices/src/logger"
	"github.com/devices/src/providers"
	"github.com/devices/src/routers"
	"github.com/devices/src/services"
)

func main() {
	// read env
	databaseUser := os.Getenv("DB_USER")
	databasePassword := os.Getenv("DB_PASSWORD")
	databaseName := os.Getenv("DB_NAME")
	envPort := os.Getenv("PORT")
	connectionString := os.Getenv("CONNECTION_STRING")

	// logger
	logger := logger.NewLogger(os.Stdout, logger.Info)

	// connections
	database := providers.GetDataBaseConnection(logger, providers.DatabaseConfiguration{User: databaseUser, Password: databasePassword, Name: databaseName, ConnectionString: connectionString})

	if database == nil {
		return
	}

	//handlers
	errorHandler := handlers.NewErrorHandler(logger)

	// services
	deviceService := services.NewDeviceService(database, logger, errorHandler)

	// controllers
	deviceController := controllers.NewDeviceController(deviceService, errorHandler)

	//routers
	router := routers.GetDeviceRouter(deviceController)

	//main
	var serverPort string

	if envPort == "" {
		serverPort = ":8080"
	} else {
		serverPort = ":" + envPort
	}

	logger.LogInfo("Server running at Port "+serverPort, "server")
	log.Fatal(http.ListenAndServe(serverPort, router))
}
