package main

import (
	"b30northwindapi/api/handlers"
	"b30northwindapi/config"
	db "b30northwindapi/db/sqlc"
	"b30northwindapi/server"
	"log"
	"os"
)

func main() {
	log.Println("Starting Northwind API")

	log.Println("Initialiazing Configuration")
	config := config.InitConfig(getConfigFileName())

	log.Println("Initializing database")
	dbHandler := server.InitDatabase(config)

	//call store, replace service
	store := db.NewStoreManager(dbHandler)

	handlerCtrl := handlers.NewHandlerManager(store)

	router := server.CreateRouter(handlerCtrl)

	log.Println("Initializig HTTP sever")
	httpServer := server.NewHttpServer(config, store, router)

	httpServer.Start()

}

func getConfigFileName() string {
	env := os.Getenv("ENV")
	if env != "" {
		return "northwind-" + env
	}

	return "northwind"
}
