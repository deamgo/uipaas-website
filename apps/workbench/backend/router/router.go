package router

import (
	"net/http"

	"github.com/deamgo/workbench/api/account"
	"github.com/deamgo/workbench/api/developer"
	"github.com/deamgo/workbench/api/workspace"
	"github.com/deamgo/workbench/context"
	"github.com/deamgo/workbench/middleware"

	"github.com/gin-gonic/gin"
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
    api.POST("/signup", account.SignUp(ctx))
		api.POST("/signup_verify", account.SignUpVerify(ctx))
		api.POST("/signin", account.SignIn(ctx))
		api.GET("/logout", middleware.JWTAuthMiddleware(), account.Logout())
		api.POST("/forgot_verify", account.SendForgotVerify(ctx))
		api.PUT("/reset_password", account.ResetPassword(ctx))
		// Get developer
		api.GET("/developer/:id", middleware.JWTAuthMiddleware(), developer.DeveloperGetByID(ctx))
		// Modify the username
		api.PUT("/developer/username/:id", middleware.JWTAuthMiddleware(), developer.DeveloperNameModify(ctx))
		// Modify the email
		api.POST("/developer/email/firststep", middleware.JWTAuthMiddleware(), developer.VerifyEmailAndPwd(ctx))
		api.POST("/developer/email/secondstep", middleware.JWTAuthMiddleware(), developer.SendModifyEmailVerify(ctx))
		api.PUT("/developer/email/thirdstep", middleware.JWTAuthMiddleware(), developer.VerifyEmailVerificationCode(ctx))
		// Modify the password
		api.POST("/developer/password/firststep", middleware.JWTAuthMiddleware(), developer.SendModifyPwdVerify(ctx))
		api.POST("/developer/password/secondstep", middleware.JWTAuthMiddleware(), developer.VerifyPwdVerificationCode(ctx))
	}
	workspaceApi := api.Group("/workspace")
	{
		workspaceApi.POST("/create", middleware.JWTAuthMiddleware(), workspace.WorkspaceCreate(ctx))
		workspaceApi.GET("/list", middleware.JWTAuthMiddleware(), workspace.WorkspaceGetListById(ctx))
		workspaceApi.POST("/logo", middleware.JWTAuthMiddleware(), workspace.WorkspaceGetLogoPath(ctx))
	}

}
