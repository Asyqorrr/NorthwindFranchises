package main

import (
	"log"
	"os"

	"b30northwindapi/config"
	"b30northwindapi/controller"
	"b30northwindapi/server"
	"b30northwindapi/services"
)

func main() {
	log.Println("Starting Northwind App")
	log.Println("Initializing configuration")

	config := config.LoadConfig(getConfigFileName(), ".")
	log.Println("Initializing database")
	dbConnection := server.InitDatabase(&config)
	defer server.Close(dbConnection)

	store := services.NewStoreManager(dbConnection)

	handlerCtrl := controller.NewControllerManager(store)

	router := server.CreateRouter(handlerCtrl, "dev")

	log.Println("Initializig HTTP sever")
	httpServer := server.NewHttpServer(&config, store, router)

	httpServer.Start()

}

func getConfigFileName() string {
	env := os.Getenv("ENV")
	if env != "" {
		return "northwind-" + env
	}

	return "northwind"
}
