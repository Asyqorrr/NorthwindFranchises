package server

import (
	"b30northwindapi/api/handlers"
	"b30northwindapi/controller"
	"b30northwindapi/middleware"
	"b30northwindapi/models"

	"github.com/gin-gonic/gin"
)

func CreateRouter(handler *handlers.HandlerManager) *gin.Engine {
	router := gin.Default()

	// set a lower memory limit for multipart form
	router.MaxMultipartMemory = 8 << 20 //8Mib
	router.Static("/static", "./public")

	api := router.Group("/api")

	api.GET("/home", func(ctx *gin.Context) {
		ctx.String(200, "Hello Gin FB")
	})

	categoryRoute := api.Group("/category")
	{

		categoryRoute.GET("", handler.GetListCategory)
		categoryRoute.GET("/", handler.GetListCategory)

	}

	return router
}

func InitRouter(handler *controller.ControllerManager, jwt *models.JWTHandler) *gin.Engine {
	router := gin.Default()

	// set a lower memory limit for multipart form
	router.MaxMultipartMemory = 8 << 20 //8Mib
	router.Static("/static", "./public")

	api := router.Group("/api")

	api.GET("/home", func(ctx *gin.Context) {
		ctx.String(200, "Hello Gin FB")
	})

	categoryRoute := api.Group("/category")
	{
		categoryRoute.Use(middleware.AuthMiddleware(jwt))
		categoryRoute.GET("", handler.GetListCategory)
		categoryRoute.GET("/", handler.GetListCategory)
		categoryRoute.GET("/:id", handler.GetCategoryById)
		categoryRoute.POST("/", handler.CreateCategory)
		categoryRoute.PUT("/:id", handler.UpdateCategory)
		categoryRoute.DELETE("/:id", handler.DeleteCategory)

	}

	productRoute := api.Group("/product")
	{
		productRoute.GET("", handler.FindAllProduct)
		productRoute.GET("/", handler.FindAllProduct)
		productRoute.POST("/", handler.CreateProduct)
		productRoute.GET("/paging", handler.FindAllProductPaging)
		productRoute.POST("/multiUpload", handler.UploadMultipleProductImage)

	}

	cartRoute := api.Group("/cart")
	{
		cartRoute.POST("/", handler.AddToCart)

	}

	orderRoute := api.Group("/order")
	{
		orderRoute.GET("/", handler.FindAllOrder)
		orderRoute.GET("/:id", handler.FindOrderById)
		orderRoute.POST("/", handler.CreateOrder)
	}

	userRoute := api.Group("/user")
	{
		userRoute.POST("/signup", handler.Signup)
		userRoute.POST("/signin", handler.Sigin)
	}

	return router
}
