package middleware

import (
	"strings"
	"time"

	"github.com/deamgo/uipaas-home/backend/pkg/e"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"net/http"
)

const TokenExpireDuration = time.Hour * 24

var CustomSecret = []byte("uipaas-home")

type CustomClaims struct {
	jwt.RegisteredClaims
}

func GenToken(username string) (string, error) {
	claims := CustomClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
			Issuer:    "uipaas-home",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(CustomSecret)
}

func ParseToken(tokenString string) (*CustomClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return CustomSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": e.ErrorEmptyToken,
				"msg":  e.EmptyToken,
			})
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"code": e.ErrorTokenFormat,
				"msg":  e.TokenFormatError,
			})
			c.Abort()
			return
		}

		_, err := ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": e.ErrorInvalidToken,
				"msg":  e.TokenInvalid,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
