package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/deamgo/workbench/auth/jwt"
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
		// Subsequent processing functions can be handled by c. Get("username") to get the requested users information
		c.Next()
	}

}
