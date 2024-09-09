package controller

import (
	db "b30northwindapi/db/sqlc"
	"b30northwindapi/models"
	"b30northwindapi/services"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type OrderFranchisesController struct {
	storedb services.Store
}


func NewOrderFranchisesController(store services.Store) *OrderFranchisesController{
	return &OrderFranchisesController{
		storedb: store,
	}
}

func (OrderFr *OrderFranchisesController) FindOrderFranchisesById (c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
	}

	order, err := OrderFr.storedb.FindOrderFranchisesById(c, int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError(err))
	}

	if order == nil {
		c.JSON(http.StatusNotFound, models.NewError(models.ErrOrderFranchisesNotFound))
	}

	c.JSON(http.StatusOK, order)

}

func (Orderfr *OrderFranchisesController) FindAllOrderFranchises (c *gin.Context){
	orders, err := Orderfr.storedb.FindAllOrderFranchises(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError(err))
	}

	c.JSON(http.StatusOK, orders)
}

func (Orderfr *OrderFranchisesController) CreateOrderFranchises(c *gin.Context){
	var payload models.CreateOrderFranchisesReq
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.NewValidationError(err))
		return
	}

	// Set timestamps to now
	timestamps := time.Now()
	ts := pgtype.Date{
		Time:   timestamps,
		Valid:  true, // Indicates that the timestamp is valid (not NULL)
	}

	args := &db.CreateOrderFranchisesParams{
		OrfiPurchaseNo: payload.OrfiPurchaseNo,
		OrfiTax: payload.OrfiTax,
		OrfiSubtotal: payload.OrfiSubtotal,
		OrfiPatrxNo: payload.OrfiPatrxNo,
		OrfiType: payload.OrfiType,
		OrfiModified: ts,
		OrfiUserID: payload.OrfiUserID,
	}

	newOrder, err := Orderfr.storedb.CreateOrderTx(c, *args)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}

	c.JSON(http.StatusCreated, newOrder)
}

// func (Orderfr *OrderFranchisesController) CreateOrderFranchises (c *gin.Context){
// 	var payload *models.CreateOrderFranchisesReq
	
// 	err := c.BindJSON(&payload)
// 	if err != nil {
// 		c.JSON(http.StatusUnprocessableEntity, models.NewValidationError(err))
// 		return
// 	}

// 	// Set timestamps to now
// 	timestamps := time.Now()
// 	ts := pgtype.Date{
// 		Time:   timestamps,
// 		Valid:  true, // Indicates that the timestamp is valid (not NULL)
// 	}

// 	args := &db.CreateOrderFranchisesParams{
// 		OrfiPurchaseNo: payload.OrfiPurchaseNo,
// 		OrfiTax: payload.OrfiTax,
// 		OrfiSubtotal: payload.OrfiSubtotal,
// 		OrfiPatrxNo: payload.OrfiPatrxNo,
// 		OrfiType: payload.OrfiType,
// 		OrfiModified: ts,
// 		OrfiUserID: payload.OrfiUserID,
// 	}
// }

// import (
// 	db "b30northwindapi/db/sqlc"
// 	"b30northwindapi/models"
// 	"b30northwindapi/services"
// 	"net/http"
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// )

// type OrderController struct {
// 	storedb services.Store
// }

// func NewOrderController(store services.Store) *OrderController {
// 	return &OrderController{
// 		storedb: store,
// 	}
// }

// func (handler *OrderController) AddToCart(c *gin.Context) {
// 	var payload models.CreateCartParams
// 	if err := c.ShouldBindJSON(&payload); err != nil {
// 		c.JSON(http.StatusUnprocessableEntity, models.NewValidationError(err))
// 		return
// 	}

// 	args := &db.FindCartByCustomerAndProductParams{
// 		CustomerID: payload.CustomerID,
// 		ProductID:  payload.ProductID,
// 	}

// 	product, _ := handler.storedb.FindCartByCustomerAndProduct(c, *args)

// 	var response = &models.CartResponse{}
// 	var cart = &db.Cart{}
// 	var err error

// 	if product == nil || product.CartID == 0 {
// 		argsAddCart := &db.CreateCartParams{
// 			CustomerID: payload.CustomerID,
// 			ProductID:  payload.ProductID,
// 			UnitPrice:  payload.UnitPrice,
// 			Qty:        payload.Qty,
// 		}
// 		// create new cart & return
// 		cart, err = handler.storedb.CreateCart(c, *argsAddCart)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, models.NewError(err))
// 			return
// 		}

// 	} else {
// 		argsUpdateCart := &db.UpdateCartQtyParams{
// 			CartID: product.CartID,
// 			Qty:    payload.Qty,
// 		}
// 		// update cart quantity
// 		cart, err = handler.storedb.UpdateCartQty(c, *argsUpdateCart)

// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, models.NewError(err))
// 			return
// 		}
// 	}

// 	//fetch all list product in carts
// 	carts, err := handler.storedb.FindCartByCustomerId(c, cart.CustomerID)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, models.ErrDataNotFound)
// 		return
// 	}

// 	response.CartId = int(carts[0].CartID)
// 	response.CustomerID = carts[0].CustomerID
// 	response.CompanyName = carts[0].CompanyName

// 	//fill carts data to dto response
// 	for _, v := range carts {
// 		product := &models.CartProductResponse{
// 			ProductID:   &v.ProductID,
// 			ProductName: &v.ProductName,
// 			UnitPrice:   v.UnitPrice,
// 			Qty:         v.Qty,
// 			Price:       v.Price,
// 		}
// 		response.Products = append(response.Products, product)
// 	}
// 	c.JSON(http.StatusCreated, response)
// }

// func (handler *OrderController) DeleteCart(c *gin.Context) error {
// 	panic("not implemented") // TODO: Implement
// }

// func (handler *OrderController) FindCartByCustomerAndProduct(c *gin.Context) {
// 	panic("not implemented") // TODO: Implement
// }

// func (handler *OrderController) FindCartByCustomerId(c *gin.Context) {
// 	id := c.Param("id")
// 	carts, err := handler.storedb.FindCartByCustomerId(c, id)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, models.NewError(err))
// 		return
// 	}
// 	var response = &models.CartResponse{}
// 	response.CartId = int(carts[0].CartID)
// 	response.CustomerID = carts[0].CustomerID
// 	response.CompanyName = carts[0].CompanyName

// 	//fill carts data to dto response
// 	for _, v := range carts {
// 		product := &models.CartProductResponse{
// 			ProductID:   &v.ProductID,
// 			ProductName: &v.ProductName,
// 			UnitPrice:   v.UnitPrice,
// 			Qty:         v.Qty,
// 			Price:       v.Price,
// 		}
// 		response.Products = append(response.Products, product)
// 	}
// 	c.JSON(http.StatusOK, response)
// }

// func (handler *OrderController) FindCartByCustomerPaging(c *gin.Context) {
// 	panic("not implemented") // TODO: Implement
// }

// func (handler *OrderController) UpdateCartQty(c *gin.Context) {
// 	panic("not implemented") // TODO: Implement
// }

// func (handler *OrderController) FindOrderById(c *gin.Context) {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	orders, err := handler.storedb.FindOrderById(c, int16(id))
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, models.NewError(err))
// 		return
// 	}
// 	c.JSON(http.StatusOK, orders)
// }

// func (handler *OrderController) FindAllOrder(c *gin.Context) {
// 	orders, err := handler.storedb.FindAllOrder(c)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, models.NewError(err))
// 		return
// 	}
// 	c.JSON(http.StatusOK, orders)
// }

// func (handler *OrderController) CreateOrder(c *gin.Context) {
// 	var payload models.CreateOrderRequest
// 	if err := c.ShouldBindJSON(&payload); err != nil {
// 		c.JSON(http.StatusUnprocessableEntity, models.NewValidationError(err))
// 		return
// 	}

// 	args := &db.CreateOrderParams{
// 		CustomerID: payload.CustomerID,
// 		EmployeeID: payload.EmployeeID,
// 		ShipVia:    payload.ShipVia,
// 	}

// 	newOrder, err := handler.storedb.CreateOrderTx(c, *args)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, models.NewError(err))
// 		return
// 	}
// 	c.JSON(http.StatusCreated, newOrder)
// }