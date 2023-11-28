package router

import (
	"github.com/deamgo/uipaas-home/backend/api/companyinfo"
	"github.com/deamgo/uipaas-home/backend/api/user"
	"github.com/deamgo/uipaas-home/backend/context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(ctx context.ApplicationContext) http.Handler {
	e := gin.New()
	// add middleware
	e.Use(gin.Recovery())
	mountAPIs(e, ctx)

	return e
}
func mountAPIs(e *gin.Engine, ctx context.ApplicationContext) {
	api := e.Group("api")
	{
		api.GET("/user/:id", user.UserGet(ctx))
		api.POST("/login", user.UserLogin(ctx))
		api.GET("/companyinfo", companyinfo.CompanyInfoGet(ctx))
	}

}
