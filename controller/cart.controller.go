package controller

import (
	db "b30northwindapi/db/sqlc"
	"b30northwindapi/models"
	"b30northwindapi/services"
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type CartController struct {
	storedb services.Store
}

func NewCartController(store services.Store) *CartController{
	return &CartController{
		storedb: store,
	}
}

func (cart *CartController) FindCartByCustomerId (c *gin.Context){
	id, err:= strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}

	id32 := int32(id)
	id32ptr := &id32
	carts, err := cart.storedb.FindCartByCustomerId(c, id32ptr)

	c.JSON(http.StatusOK, carts)
}

func (cart *CartController) CreateCart (c *gin.Context){
	var payload *models.CreateCartReq
	
	err := c.BindJSON(&payload) 
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.NewError(err))
		return
	}

	// Set timestamps to now
	timestamps := time.Now()
	ts := pgtype.Timestamp{
		Time:   timestamps,
		Valid:  true, // Indicates that the timestamp is valid (not NULL)
	}

	args := &db.CreateCartParams{
		CartUserID: payload.CartUserID,
		CartFrID: payload.CartFrID,
		CartStartDate: payload.CartStartDate,
		CartEndDate: payload.CartEndDate,
		CartQty: payload.CartQty,
		CartPrice: payload.CartPrice,
		CartTotalPrice: payload.CartTotalPrice,
		CartModified: ts,
		CartStatus: payload.CartStatus,
		CartCartID: payload.CartCartID,
	}

	carts, err := cart.storedb.CreateCart(c, *args)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError(err))
	}

	c.JSON(http.StatusCreated, carts)
}

func (cart *CartController) DeleteCart (c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))

	_,err := cart.storedb.FindCartsbyId(c, int32(id))
	if err != nil {
		if err == sql.ErrNoRows{
			c.JSON(http.StatusNotFound, models.ErrDataNotFound)
			return
		}
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}

	err = cart.storedb.DeleteCart(c, int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"status" : "success", "message" : "cart has been deleted"})
}


func (cartC *CartController) AddToCart(c *gin.Context){


	var payload models.CreateCartReq
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.NewValidationError(err))
		return
	}

	args := &db.FindCartByCustomerAndFranchisesParams{
		CartUserID: payload.CartUserID,
		CartFrID  : payload.CartFrID,
	}

	franchise,_ := cartC.storedb.FindCartByCustomerAndFranchises(c, *args)

	var response = &models.CartResponse{}
	var cart = &db.Cart{}
	var err error

	if franchise == nil || franchise.CartID == 0 {
		// Set timestamps to now
		timestamps := time.Now()
		ts := pgtype.Timestamp{
		Time:   timestamps,
		Valid:  true, // Indicates that the timestamp is valid (not NULL)
		}

		var floatCartPrice *float64
    	var intQty *int32
	
		// Assign the pointers
		floatCartPrice = payload.CartPrice
		intQty = payload.CartQty

		floatValCartPrice := *floatCartPrice
        intValQty := *intQty

        // Convert the int32 value to float64
        intValAsFloat := float64(intValQty)

        // Multiply the values
        result := floatValCartPrice * intValAsFloat

		argsAddCart := &db.CreateCartParams{
			CartUserID     : payload.CartUserID,
			CartFrID       : payload.CartFrID,
			CartStartDate  : payload.CartStartDate,
			CartEndDate    : payload.CartEndDate,
			CartQty        : payload.CartQty,
			CartPrice      : payload.CartPrice,
			CartTotalPrice : &result, //total price * qty
			CartModified   : ts, //time now
			CartStatus     : payload.CartStatus,
			CartCartID     : payload.CartCartID,
		}

		cart, err = cartC.storedb.CreateCart(c, *argsAddCart)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.NewError(err))
			return
		}
	} else {
		argsUpdateCart := &db.UpdateCartQTYParams{
			CartID: 	franchise.CartID,
			CartQty: 	payload.CartQty,
		}

		cart, err = cartC.storedb.UpdateCartQTY(c, *argsUpdateCart)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.NewError(err))
			return
		}
	}

	carts, err := cartC.storedb.FindCartByCustomerId(c, cart.CartUserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrDataNotFound)
		return
	}

	response.CartID = int32(carts[0].CartID)
	response.CartUserID = carts[0].CartUserID

	for _, v := range carts {
		franchise := &models.CartFranchisesResponse{
			FrchID: v.CartFrID,
			UserID: v.CartUserID,
			FrchName: v.FrchName,
			QTY: v.CartQty,
			Price: v.CartPrice,
		}
		response.Franchises = append(response.Franchises, franchise)
	}
	c.JSON(http.StatusCreated, response)
}