package services

import (
	db "b30northwindapi/db/sqlc"
	"context"
	"log"
)

func (sm *StoreManager) CreateOrderTx(args db.CreateOrderParams) (*db.Order, error) {
	tx, err := sm.dbConn.Begin(context.Background())
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(context.Background())

	qtx := sm.Queries.WithTx(tx)

	//populate cart list
	carts, err := qtx.FindCartByCustomerId(context.Background(), *args.CustomerID)
	if err != nil {
		return nil, err
	}

	log.Println(carts)

	//create order
	newOrder, err := qtx.CreateOrder(context.Background(), args)
	if err != nil {
		return nil, err
	}
	if err = tx.Commit(context.Background()); err != nil {
		return nil, err
	}
	return newOrder, nil

}
