package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/deamgo/workbench/api/developer"
	"github.com/deamgo/workbench/context"
	"github.com/deamgo/workbench/middleware"
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
	// api.Any("/*", middleware.JWTAuthMiddleware())
	{
		api.POST("/signup", developer.SignUp(ctx))
		api.POST("/signup_verify", developer.SignUpVerify(ctx))
		api.POST("/signin", developer.SignIn(ctx))
		api.POST("/forgot_verify", developer.ForgotVerifySend(ctx))
		api.PUT("/reset_password", developer.ResetPassword(ctx))
		//api.PUT("/developer/username/:id")
	}

}
