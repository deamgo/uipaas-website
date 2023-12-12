package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"workbench/api/user"
	"workbench/context"
	"workbench/middleware"
)

func NewRouter(ctx context.ApplicationContext) http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.Use(middleware.Cors())
	mountAPIs(e, ctx)
	return e
}
func mountAPIs(e *gin.Engine, ctx context.ApplicationContext) {
	api := e.Group("v1")
	{
		api.POST("/signup", user.SignUp(ctx))
		api.POST("/signup_verify", user.SignUpVerify(ctx))
		api.POST("/signin", user.SignIn(ctx))
	}

}
