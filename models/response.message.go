package models

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type CreateUserReq struct {
	UserName     *string `json:"user_name" binding:"required"`
	UserPassword *string `json:"user_password" binding:"required"`
	UserPhone	 *string `json:"user_phone"`
}

type UserResponse struct {
	UserID       int32               `json:"user_id"`
	UserName     *string             `json:"user_name"`
	UserPassword *string             `json:"user_password"`
	UserPhone    *string             `json:"user_phone"`
	UserToken    *string             `json:"user_token"`
	Roles        []*UserRoleResponse `json:"roles"`
}

type UserRoleResponse struct {
	RoleID   int32   `json:"role_id"`
	RoleName *string `json:"role_name"`
}

type CategoryPostReq struct {
	CategoryName string  `json:"cate_name" binding:"required"`
}

type CategoryUpdateReq struct {
	CategoryName string  `json:"cate_name"`
}

type CreateCartParams struct {
	CustomerID string   `json:"customer_id"`
	ProductID  int32    `json:"product_id"`
	UnitPrice  *float32 `json:"unit_price"`
	Qty        *int32   `json:"qty"`
}

// type CartResponse struct {
// 	CartId      int                    `json:"cart_id"`
// 	CustomerID  string                 `json:"customer_id"`
// 	CompanyName string                 `json:"company_name"`
// 	Products    []*CartProductResponse `json:"products"`
// }

type CartProductResponse struct {
	ProductID   *int16         `json:"product_id"`
	ProductName *string        `json:"product_name"`
	UnitPrice   *float32       `json:"unit_price"`
	Qty         *int32         `json:"qty"`
	Price       pgtype.Numeric `json:"price"`
}

type CreateOrderRequest struct {
	CustomerID     *string     `json:"customer_id" binding:"required"`
	EmployeeID     *int16      `json:"employee_id" binding:"required"`
	OrderDate      pgtype.Date `json:"order_date"`
	RequiredDate   pgtype.Date `json:"required_date"`
	ShippedDate    pgtype.Date `json:"shipped_date"`
	ShipVia        *int16      `json:"ship_via" binding:"required"`
	Freight        *float32    `json:"freight"`
	ShipName       *string     `json:"ship_name"`
	ShipAddress    *string     `json:"ship_address"`
	ShipCity       *string     `json:"ship_city"`
	ShipRegion     *string     `json:"ship_region"`
	ShipPostalCode *string     `json:"ship_postal_code"`
	ShipCountry    *string     `json:"ship_country"`
}

type CreateFranchisesReq struct {
	FrchName        *string     		`json:"frch_name"`
	FrchDesc        *string     		`json:"frch_desc"`
	FrchPriceBuyout *float64    		`json:"frch_price_buyout"`
	FrchPriceYearly *float64   			`json:"frch_price_yearly"`
	FrchModified    pgtype.Timestamp 	`json:"frch_modified"`
	FrchCateID      *int32      		`json:"frch_cate_id"`
}



type CreateOrderFranchisesReq struct {
	OrfiPurchaseNo *string     `json:"orfi_purchase_no"`
	OrfiTax        *float64    `json:"orfi_tax"`
	OrfiSubtotal   *float64    `json:"orfi_subtotal"`
	OrfiPatrxNo    *string     `json:"orfi_patrx_no"`
	OrfiType       *string     `json:"orfi_type"`
	OrfiModified   pgtype.Date `json:"orfi_modified"`
	OrfiUserID     *int32      `json:"orfi_user_id"`
}

type CreateCartReq struct {
	CartUserID     *int32           `json:"cart_user_id"`
	CartFrID       *int32           `json:"cart_fr_id"`
	CartStartDate  pgtype.Timestamp `json:"cart_start_date"`
	CartEndDate    pgtype.Timestamp `json:"cart_end_date"`
	CartQty        *int32           `json:"cart_qty"`
	CartPrice      *float64         `json:"cart_price"`
	CartTotalPrice *float64         `json:"cart_total_price"`
	CartModified   pgtype.Timestamp `json:"cart_modified"`
	CartStatus     *string          `json:"cart_status"`
	CartCartID     *int32           `json:"cart_cart_id"`
}

type CartResponse struct {
	CartID         int32            			`json:"cart_id"`
	CartUserID     *int32           			`json:"cart_user_id"`
	Franchises     []*CartFranchisesResponse	`json:"franchises"`
}

type CartFranchisesResponse struct {
	FrchID 		*int32		`json:"franchise_id"`
	UserID 		*int32		`json:""user_id`
	FrchName 	*string		`json:"franchise_name"`
	QTY 		*int32		`json:"qty"`
	Price 		*float64	`json:price`
}

type OrderFranchiseDetailsObj struct{
	OrderID			int16	`json:"order_id"`
	Franchise		[]*DetailsFranchise
	TotalPrice 		float64 `json:"total_price"`
}

type DetailsFranchise struct{
	FrchID  			int16
	FrchName    		*string
	FrchDesc    		*string           
	FrchPriceBuyout     *float64         
	FrchPriceYearly 	*float64         
	FrchModified     	pgtype.Timestamp         
	FrchCateID     		*int32
}