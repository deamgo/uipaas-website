package router

import (
	"github.com/deamgo/workbench/api/workspace"
	"github.com/gin-gonic/gin"
	"net/http"

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
	{
		api.POST("/signup", user.SignUp(ctx))
		api.POST("/signup_verify", user.SignUpVerify(ctx))
		api.POST("/signin", user.SignIn(ctx))
		api.POST("/workspace/create", workspace.WorkspaceCreate(ctx))
	}

}
