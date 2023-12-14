package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/deamgo/workbench/api/user"
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
		api.POST("/signup", user.SignUp(ctx))
		api.POST("/signup_verify", user.SignUpVerify(ctx))
		api.POST("/signin", user.SignIn(ctx))
		api.POST("/forgot_verify", user.ForgotVerifySend(ctx))
		api.PATCH("/reset_password", user.ResetPassword(ctx))
	}

}
