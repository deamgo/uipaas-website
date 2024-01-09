package router

import (
	"net/http"

	"github.com/deamgo/workbench/api/account"
	"github.com/deamgo/workbench/api/application"
	"github.com/deamgo/workbench/api/devdepot"
	"github.com/deamgo/workbench/api/developer"
	"github.com/deamgo/workbench/api/workspace"
	"github.com/deamgo/workbench/api/ws"
	"github.com/deamgo/workbench/context"
	"github.com/deamgo/workbench/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter(ctx context.ApplicationContext) http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.Use(middleware.Cors())
	e.GET("/v1/ws", ws.WSHandler())
	mountAPIs(e, ctx)
	return e
}
func mountAPIs(e *gin.Engine, ctx context.ApplicationContext) {
	api := e.Group("v1")
	{
		api.POST("/signup", account.SignUp(ctx))
		api.POST("/signup_verify", account.SignUpVerify(ctx))
		api.POST("/signin", account.SignIn(ctx))
		api.POST("/forgot_verify", account.SendForgotVerify(ctx))
		api.PUT("/reset_password", account.ResetPassword(ctx))
		api.GET("/logout", middleware.JWTAuthMiddleware(), account.Logout())
	}
	developerAPI := e.Group("v1", middleware.JWTAuthMiddleware())
	{

		// Get developer
		developerAPI.GET("/developer", developer.DeveloperGetByID(ctx))
		// Modify the username
		developerAPI.PUT("/developer/username/:id", developer.DeveloperNameModify(ctx))
		// Modify the email
		developerAPI.POST("/developer/email/firststep", developer.VerifyEmailAndPwd(ctx))
		developerAPI.POST("/developer/email/secondstep", developer.SendModifyEmailVerify(ctx))
		developerAPI.PUT("/developer/email/thirdstep", developer.VerifyEmailVerificationCode(ctx))
		// Modify the password
		developerAPI.POST("/developer/password/firststep", developer.SendModifyPwdVerify(ctx))
		developerAPI.POST("/developer/password/secondstep", developer.VerifyPwdVerificationCode(ctx))
		developerAPI.PUT("/developer/password/thirdstep", developer.ModifyPwd(ctx))
	}
	workspaceApi := api.Group("/workspace", middleware.JWTAuthMiddleware())
	{
		workspaceApi.DELETE("/:workspace_id/settings", middleware.AuthMiddleware(), workspace.WorkspaceDel(ctx))
		workspaceApi.PUT("/:workspace_id/settings", middleware.AuthMiddleware(), workspace.WorkspaceNameModify(ctx))
		workspaceApi.POST("/create", workspace.WorkspaceCreate(ctx))
		workspaceApi.GET("/list", workspace.WorkspaceGetListById(ctx))
		workspaceApi.POST("/logo", workspace.WorkspaceGetLogoPath(ctx))
		workspaceApi.GET("/:workspace_id/developer", middleware.AuthMiddleware(), devdepot.DevdepotList(ctx))
		workspaceApi.GET("/:workspace_id/developer/search", middleware.AuthMiddleware(), devdepot.DevdepotSearch(ctx))
		workspaceApi.DELETE("/:workspace_id/developer", middleware.AuthMiddleware(), devdepot.DevdepotDel(ctx))
		workspaceApi.PUT("/:workspace_id/developer", middleware.AuthMiddleware(), devdepot.DevdepotRoleModify(ctx))
		workspaceApi.POST("/:workspace_id/developer/invite", middleware.AuthMiddleware(), devdepot.DevDepotInvite(ctx))

	}
	applicationAPI := e.Group("v1", middleware.JWTAuthMiddleware(), middleware.AuthMiddleware())
	{
		applicationAPI.POST("/workspace/:workspace_id/application", application.ApplicationCreate(ctx))
		applicationAPI.GET("/workspace/:workspace_id/application", application.ApplicationList(ctx))
		applicationAPI.GET("/workspace/:workspace_id/application/search", application.ApplicationSearch(ctx))
	}

}
