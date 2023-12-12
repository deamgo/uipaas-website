package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"workbench/auth/jwt"
)

// 基于JWT认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 2003,
				"msg":  "请求头中的auth为空",
			})
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)

		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"code": 2004,
				"msg":  "请求头中的auth格式错误",
			})
			//阻止调用后续的函数
			c.Abort()
			return
		}
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 2005,
				"msg":  "无效的token",
			})
			c.Abort()
			return
		}
		//将当前请求的username信息保存到请求的上下文c上
		c.Set("username", mc.Username)
		//后续的处理函数可以通过c.Get("username")来获取请求的用户信息
		c.Next()
	}

}
