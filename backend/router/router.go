package router

import (
	"github.com/deamgo/uipaas-home/backend/middleware"
	"net/http"

	"github.com/deamgo/uipaas-home/backend/api/company"
	"github.com/deamgo/uipaas-home/backend/api/user"
	"github.com/deamgo/uipaas-home/backend/context"

	"github.com/gin-gonic/gin"
)

func NewRouter(ctx context.ApplicationContext) http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	mountAPIs(e, ctx)

	return e
}
func mountAPIs(e *gin.Engine, ctx context.ApplicationContext) {
	api := e.Group("v1")
	{
		api.GET("/user/:id", user.UserGet(ctx))
		api.POST("/login", user.UserLogin(ctx))
		api.GET("/company", middleware.JWTAuthMiddleware(), company.CompanyGet(ctx))
		api.POST("/company", middleware.JWTAuthMiddleware(), company.CompanyAdd(ctx))

	}

}
