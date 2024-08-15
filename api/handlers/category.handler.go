package handlers

import (
	db "b30northwindapi/db/sqlc"
	"b30northwindapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	storedb db.Store
}

func NewCategoryHandler(store db.Store) *CategoryHandler {
	return &CategoryHandler{
		storedb: store,
	}
}

func (handler *CategoryHandler) GetListCategory(c *gin.Context) {
	categories, err := handler.storedb.FindAllCategory(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrDataNotFound)
	}

	c.JSON(http.StatusOK, categories)
}
