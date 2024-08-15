package server

import (
	"b30northwindapi/controller"
	db "b30northwindapi/db/sqlc"
	"b30northwindapi/models"
	"b30northwindapi/services"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
)

type HttpServer struct {
	Config *viper.Viper
	Router *gin.Engine
	Store  db.Store
}

// new design
func NewHttpServer(config *viper.Viper, store db.Store, router *gin.Engine) *HttpServer {
	return &HttpServer{
		Config: config,
		Router: router,
		Store:  store,
	}
}

// old design pattern
func InitHttpServer(config *viper.Viper, dbHandler *pgxpool.Conn) *HttpServer {

	//call jwthandler
	jwtHandler := models.NewJWTHandler(config)

	//call handler & service manager
	serviceManager := services.NewServiceManager(dbHandler, jwtHandler)
	controllerManager := controller.NewControllerManager(serviceManager)

	router := InitRouter(controllerManager, jwtHandler)

	return &HttpServer{
		Config: config,
		Router: router,
	}
}

func (hs HttpServer) Start() {
	err := hs.Router.Run(hs.Config.GetString("http.server_address"))
	if err != nil {
		log.Fatalf("Error while starting HTTP server: %v", err)
	}
}
