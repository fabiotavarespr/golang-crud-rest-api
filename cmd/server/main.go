package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fabiotavarespr/golang-crud-rest-api/internal/config"
	"github.com/fabiotavarespr/golang-crud-rest-api/internal/database"
	"github.com/fabiotavarespr/golang-crud-rest-api/internal/logger"
	"github.com/fabiotavarespr/golang-crud-rest-api/internal/routes"
	"github.com/gorilla/mux"
)

func main() {
	// Load Configurations from config.json using Viper
	config.LoadAppConfig()
	// Initialize Database
	database.Connect(config.AppConfig.ConnectionString)
	database.Migrate()

	// Initialize the router
	router := mux.NewRouter().StrictSlash(true)
	// Register Routes
	routes.RegisterProductRoutes(router)
	// Start the server
	logger.Info(fmt.Sprintf("Starting Server on port %s", config.AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", config.AppConfig.Port), router))
}
