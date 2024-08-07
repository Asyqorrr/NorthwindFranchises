package services

import (
	"b30northwindapi/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ServiceManager struct {
	*CategoryService
	*ProductService
	*CartService
	*OrderService
	*UserService
}

func NewServiceManager(dbConn *pgxpool.Conn, jwt *models.JWTHandler) *ServiceManager {
	return &ServiceManager{
		CategoryService: NewCategoryService(dbConn),
		ProductService:  NewProductService(dbConn),
		CartService:     NewCartService(dbConn),
		OrderService:    NewOrderService(dbConn),
		UserService:     NewUserService(dbConn, jwt),
	}
}
