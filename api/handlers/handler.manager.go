package handlers

import db "b30northwindapi/db/sqlc"

type HandlerManager struct {
	*CategoryHandler
}

func NewHandlerManager(store db.Store) *HandlerManager {
	return &HandlerManager{
		CategoryHandler: NewCategoryHandler(store),
	}
}
