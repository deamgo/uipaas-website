package middleware

import (
	auth "github.com/deamgo/workbench/auth/permission"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		auth.AuthenticationPermissions(c)
	}
}
