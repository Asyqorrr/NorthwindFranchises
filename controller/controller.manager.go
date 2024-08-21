package controller

import "b30northwindapi/services"

type ControllerManager struct {
	*CategoryController
	*OrderController
	*ProductController
	*UserController
}

func NewControllerManager(store services.Store) *ControllerManager {
	return &ControllerManager{
		CategoryController: NewCategoryController(store),
		OrderController:    NewOrderController(store),
		ProductController:  NewProductController(store),
		UserController:     NewUserController(store),
	}
}
