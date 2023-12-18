package main

import (
	"github.com/gin-gonic/gin"

	"github.com/deamgo/workbench/context"
	user2 "github.com/deamgo/workbench/dao/user"
	workspace2 "github.com/deamgo/workbench/dao/workspace"
	"github.com/deamgo/workbench/db"
	"github.com/deamgo/workbench/pkg/logger"
	routes "github.com/deamgo/workbench/router"
	"github.com/deamgo/workbench/service/user"
	"github.com/deamgo/workbench/service/workspace"
)

func main() {

	db.InitRedis()
	dao := user2.NewAUserDao(db.DB)
	workspacedao := workspace2.NewWorkspaceDao(db.DB)
	ctx := context.ApplicationContext{
		UserService:      user.NewUserService(user.UserServiceParams{Dao: dao}),
		WorkspaceService: workspace.NewWorkspaceService(workspace.WorkspaceServiceParams{Dao: workspacedao}),
	}

	r := gin.Default()

	// routing
	user := routes.NewRouter(ctx)
	r.Any("/*any", gin.WrapH(user))
	err := r.Run(":8989")

	if err != nil {
		logger.LoggersObj.Error("Start failed!")
	}
}
