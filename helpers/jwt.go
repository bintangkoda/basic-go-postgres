package helpers

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var secretKey = "rahasia"

func GenerateToken(id uint, email string) string {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
	}

	fmt.Println(claims, "claims")

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := parseToken.SignedString([]byte(secretKey))
	if err != nil {
		fmt.Println(err, "error")
	}

	return signedToken
}

func VerifyToken(c *gin.Context) (jwt.MapClaims, error) {
	errResponse := errors.New("sign in to procced")
	headerToken := c.Request.Header.Get("Authorization")

	bearer := strings.HasPrefix(headerToken, "Bearer")
	if !bearer {
		return nil, errResponse
	}

	stringToken := strings.Split(headerToken, " ")[1]
	fmt.Println(stringToken, "ini token")

	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}
		return []byte(secretKey), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errResponse
	}

	fmt.Println(token.Claims.(jwt.MapClaims)["id"], "ini id")

	return token.Claims.(jwt.MapClaims), nil
	// claims, ok := token.Claims.(jwt.MapClaims)
	// if !ok && !token.Valid {
	// 	return nil, errResponse
	// }

	// return claims, nil
}
