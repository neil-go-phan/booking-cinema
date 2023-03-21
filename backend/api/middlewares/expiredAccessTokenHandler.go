package middlewares

import (
	"booking-cinema-backend/helper"
	"errors"

	"github.com/gin-gonic/gin"
)


func ExpiredAccessTokenHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("x-refresh-token")
		if tokenString == "" {
			c.Error(errors.New(helper.ERROR_VALIDATE_TOKEN_FAIL.ErrorName))
			c.Abort()
			return
		}

		claims, err := validateToken(tokenString)
		if err != nil {
			c.Error(errors.New(helper.ERROR_VALIDATE_TOKEN_FAIL.ErrorName))
			c.Abort()
			return
		}
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)
		c.Next()
	}
}