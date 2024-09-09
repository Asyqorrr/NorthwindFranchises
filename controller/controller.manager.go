package controller

import "b30northwindapi/services"

type ControllerManager struct {
	*CategoryController
	*OrderFranchisesController
	// *ProductController
	*UserController
	*FranchisesController
	*CartController
}

func NewControllerManager(store services.Store) *ControllerManager {
	return &ControllerManager{
		CategoryController: NewCategoryController(store),
		OrderFranchisesController:    NewOrderFranchisesController(store),
		// ProductController:  NewProductController(store),
		UserController:     NewUserController(store),
		FranchisesController:	NewFranchisesController(store),
		CartController: NewCartController(store),
	}
}
