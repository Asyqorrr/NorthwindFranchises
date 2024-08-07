package middleware

import (
	"b30northwindapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(jwt *models.JWTHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := jwt.TokenValid(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}
