package services

import (
	db "b30northwindapi/db/sqlc"
	"b30northwindapi/models"
	"context"
	"log"
)

func (sm *StoreManager) CreateOrderTx(ctx context.Context, OrderArgs db.CreateOrderFranchisesParams) (*models.OrderFranchiseDetailsObj, error){
	tx, err := sm.dbConn.Begin(context.Background())
	if err != nil {
		return nil, err
	}

	defer tx.Rollback(context.Background())

	qtx := sm.Queries.WithTx(tx)

	// populate cart list
	carts, err := qtx.FindCartByCustomerId(ctx, *&OrderArgs.OrfiUserID)
	if err != nil {
		return nil ,err
	}

	log.Println(carts)

	// create order
	newOrder, err := qtx.CreateOrderFranchises(ctx, OrderArgs)
	if err != nil {
		return nil, err
	}

	// Create Order Details
	var newOrderDetailsSlice []*db.OrderFranchisesDetail
	var newOrderDetailsResponse = &models.OrderFranchiseDetailsObj{}
	var detailsFranchiseSlice []*models.DetailsFranchise
	sumTotalPrice := 0.0
	
	for _, v := range carts {
		totalPrice := (*v.CartPrice) * float64(*v.CartQty)

		OrdDetailsArgs := &db.CreateOrderFranchisesDetailParams{
			OfdeStartDate: v.CartStartDate,
			OfdeEndDate: v.CartEndDate,
			OfdeQtyUnit: v.CartQty,
			OfdePrice: v.CartPrice,
			OfdeTotalPrice: &totalPrice,
			OfdeOrfiID: &newOrder.OrfiID,
			OfdeFrchID: v.CartFrID,
		}

		newOrderDetails, err := qtx.CreateOrderFranchisesDetail(ctx, *OrdDetailsArgs)
		if err != nil {
			return nil, err
		}

		newOrderDetailsSlice := append(newOrderDetailsSlice, newOrderDetails)
		log.Println(newOrderDetailsSlice)

		franchise, err := qtx.GetFranchisesById(ctx, *v.CartFrID)

		DetailsFranchiseResponse := &models.DetailsFranchise{
			FrchID: int16(*v.CartFrID),
			FrchName: franchise.FrchName,
			FrchDesc: franchise.FrchDesc,
			FrchPriceBuyout: franchise.FrchPriceBuyout,
			FrchPriceYearly: franchise.FrchPriceYearly,
			FrchModified: v.CartModified,
			FrchCateID: franchise.FrchCateID,
		}
		detailsFranchiseSlice = append(detailsFranchiseSlice, DetailsFranchiseResponse)
		

		sumTotalPrice += totalPrice
	}

	tax := float64(*OrderArgs.OrfiTax)


	newOrderDetailsResponse.OrderID = int16(newOrder.OrfiID)
	newOrderDetailsResponse.Franchise = detailsFranchiseSlice
	newOrderDetailsResponse.TotalPrice = sumTotalPrice * tax

	// type UpdateOrderParams struct {
	// 	OrfiID         int32    `json:"orfi_id"`
	// 	OrfiPurchaseNo *string  `json:"orfi_purchase_no"`
	// 	OrfiTax        *float64 `json:"orfi_tax"`
	// 	OrfiSubtotal   *float64 `json:"orfi_subtotal"`
	// 	OrfiPatrxNo    *string  `json:"orfi_patrx_no"`
	// 	OrfiType       *string  `json:"orfi_type"`
	// }

	argsUpdate := &db.UpdateOrderParams{
		OrfiID: newOrder.OrfiID,
		OrfiPurchaseNo: OrderArgs.OrfiPurchaseNo,
		OrfiTax: OrderArgs.OrfiTax,
		OrfiSubtotal: &sumTotalPrice,
		OrfiType: OrderArgs.OrfiType,
	}

	_, err = qtx.UpdateOrder(ctx, *argsUpdate)
	if err != nil {
		return nil, err
	}


	if err = tx.Commit(context.Background()); err != nil {
		return nil, err
	}

	return newOrderDetailsResponse, nil
}

// func (sm *StoreManager) CreateOrderTx(ctx context.Context, args db.CreateOrderParams) (*db.Order, error) {
// 	tx, err := sm.dbConn.Begin(context.Background())
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer tx.Rollback(context.Background())

// 	qtx := sm.Queries.WithTx(tx)

// 	//populate cart list
// 	carts, err := qtx.FindCartByCustomerId(ctx, *args.CustomerID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	log.Println(carts)

// 	//create order
// 	newOrder, err := qtx.CreateOrder(ctx, args)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if err = tx.Commit(context.Background()); err != nil {
// 		return nil, err
// 	}
// 	return newOrder, nil

// }
