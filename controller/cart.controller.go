package controller

import (
	db "b30northwindapi/db/sqlc"
	"b30northwindapi/models"
	"b30northwindapi/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type CartController struct {
	serviceManager *services.ServiceManager
}

// constructor
func NewCartController(servicesManager services.ServiceManager) *CartController {
	return &CartController{
		serviceManager: &servicesManager,
	}
}

type CreateCartParams struct {
	CustomerID string   `json:"customer_id"`
	ProductID  int32    `json:"product_id"`
	UnitPrice  *float32 `json:"unit_price"`
	Qty        *int32   `json:"qty"`
}

type CartResponse struct {
	CartId      int                    `json:"cart_id"`
	CustomerID  string                 `json:"customer_id"`
	CompanyName string                 `json:"company_name"`
	Products    []*CartProductResponse `json:"products"`
}

type CartProductResponse struct {
	ProductID   *int16         `json:"product_id"`
	ProductName *string        `json:"product_name"`
	UnitPrice   *float32       `json:"unit_price"`
	Qty         *int32         `json:"qty"`
	Price       pgtype.Numeric `json:"price"`
}

func (handler *CartController) AddToCart(c *gin.Context) {
	var payload CreateCartParams
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.NewValidationError(err))
		return
	}

	args := &db.FindCartByCustomerAndProductParams{
		CustomerID: payload.CustomerID,
		ProductID:  payload.ProductID,
	}

	product, _ := handler.serviceManager.CartService.FindCartByCustomerAndProduct(c, *args)

	var response = &CartResponse{}
	var cart = &db.Cart{}
	var err error

	if product == nil || product.CartID == 0 {
		argsAddCart := &db.CreateCartParams{
			CustomerID: payload.CustomerID,
			ProductID:  payload.ProductID,
			UnitPrice:  payload.UnitPrice,
			Qty:        payload.Qty,
		}
		// create new cart & return
		cart, err = handler.serviceManager.CartService.CreateCart(c, *argsAddCart)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.NewError(err))
			return
		}

	} else {
		argsUpdateCart := &db.UpdateCartQtyParams{
			CartID: product.CartID,
			Qty:    payload.Qty,
		}
		// update cart quantity
		cart, err = handler.serviceManager.CartService.UpdateCartQty(c, *argsUpdateCart)

		if err != nil {
			c.JSON(http.StatusInternalServerError, models.NewError(err))
			return
		}
	}

	//fetch all list product in carts
	carts, err := handler.serviceManager.CartService.FindCartByCustomerId(c, cart.CustomerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrDataNotFound)
		return
	}

	response.CartId = int(carts[0].CartID)
	response.CustomerID = carts[0].CustomerID
	response.CompanyName = carts[0].CompanyName

	//fill carts data to dto response
	for _, v := range carts {
		product := &CartProductResponse{
			ProductID:   &v.ProductID,
			ProductName: &v.ProductName,
			UnitPrice:   v.UnitPrice,
			Qty:         v.Qty,
			Price:       v.Price,
		}
		response.Products = append(response.Products, product)
	}
	c.JSON(http.StatusCreated, response)
}
