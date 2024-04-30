package middlewares

import (
	"fmt"
	"go-jwt/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthenticated",
				"message": err.Error(),
			})
			return
		}

		fmt.Println(verifyToken["id"], "ini id")

		c.Set("userData", verifyToken)
		c.Next()
	}
}
