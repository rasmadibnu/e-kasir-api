package security

import (
	"net/http"

	"kasir-cepat-api/helper"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := TokenValid(c)
		if err != nil {
			resp := helper.ErrorJSON(c, "Unauthorized", http.StatusUnauthorized, nil)

			c.JSON(http.StatusUnauthorized, resp)
			c.Abort()
			return
		}
		c.Next()
	}
}
