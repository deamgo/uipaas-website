package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/deamgo/workbench/auth/jwt"
	"github.com/deamgo/workbench/db"

	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware based On JWT Certified Middleware
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 2003,
				"msg":  "the Auth In The RequestHeader Is Empty",
			})
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)

		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"code": 2004,
				"msg":  "the Auth Format In The RequestHeader Is Incorrect",
			})
			// prevent Subsequent Functions From Being Called
			c.Abort()
			return
		}
		// Check if the token is on the blacklist

		if jwt.TokenBlacklist[authHeader] {
			c.JSON(http.StatusOK, gin.H{
				"code": 2004,
				"msg":  "The login is invalid, please log in again",
			})
			c.Abort()
			return
		}
		isExpireToken, err := jwt.IsExpireToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 2005,
				"msg":  err,
			})
			c.Abort()
			return
		}
		if isExpireToken {
			id, _ := jwt.ExtractIDFromToken(authHeader)
			newToken, err := jwt.GenToken(id)
			db.RedisDB.Set(id, newToken, time.Hour*2)
			if err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
			}
			c.JSON(http.StatusOK, gin.H{
				"code": 2006,
				"msg":  "Token Expired",
				"data": struct {
					Token string `json:"token"`
				}{newToken},
			})
			c.Abort()
			return
		}

		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 2005,
				"msg":  "invalid Token",
			})
			c.Abort()
			return
		}
		// Save the username information of the current request to the context C of the request
		c.Set("username", mc.ID)
		// Subsequent processing functions can be handled by c. Get("username") to get the requested developer information
		c.Next()
	}

}
