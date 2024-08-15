package db

import "github.com/jackc/pgx/v5/pgxpool"

type Store interface {
	Querier
}

type StoreManager struct {
	*Queries
	db *pgxpool.Conn
}

func NewStoreManager(dbConn *pgxpool.Conn) Store {
	return &StoreManager{
		Queries: New(dbConn),
		db:      dbConn,
	}
}
