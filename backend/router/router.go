package router

import (
	"net/http"

	"github.com/deamgo/uipass-waitlist-page/backend/api/user"
	"github.com/deamgo/uipass-waitlist-page/backend/context"

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
		api.Handle("GET", "/user", user.UserGet(ctx))
	}

}
