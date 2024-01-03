package main

import (
	"github.com/deamgo/workbench/context"
	application2 "github.com/deamgo/workbench/dao/application"
	devdepot2 "github.com/deamgo/workbench/dao/devdepot"
	user2 "github.com/deamgo/workbench/dao/developer"
	workspace2 "github.com/deamgo/workbench/dao/workspace"
	"github.com/deamgo/workbench/db"
	"github.com/deamgo/workbench/initialize"
	"github.com/deamgo/workbench/pkg/logger"
	routes "github.com/deamgo/workbench/router"
	"github.com/deamgo/workbench/service/application"
	"github.com/deamgo/workbench/service/devdepot"
	"github.com/deamgo/workbench/service/developer"
	"github.com/deamgo/workbench/service/mail"
	"github.com/deamgo/workbench/service/workspace"

	"github.com/gin-gonic/gin"
)

func main() {

	initialize.InitConfig()
	db.InitDB()
	db.InitRedis()
	developerDao := user2.NewADeveloperDao(db.DB)
	workspaceDao := workspace2.NewWorkspaceDao(db.DB)
	devdepotDao := devdepot2.NewDevDepotDao(db.DB)
	applicationDao := application2.NewApplicationDao(db.DB)
	ctx := context.ApplicationContext{
		UserService:        developer.NewDeveloperService(developer.DeveloperServiceParams{Dao: developerDao, MailService: mail.NewMailService()}),
		MailService:        mail.NewMailService(),
		WorkspaceService:   workspace.NewWorkspaceService(workspace.WorkspaceServiceParams{Dao: workspaceDao, DeveloperDao: developerDao}),
		DevDepotService:    devdepot.NewDepotService(devdepot.DevDepotServiceParams{Dao: devdepotDao, DeveloperDao: developerDao, MailService: mail.NewMailService()}),
		ApplicationService: application.NewApplicationService(application.ApplicationServiceParams{Dao: applicationDao}),
	}

	r := gin.Default()

	// routing
	user := routes.NewRouter(ctx)
	r.Any("/*any", gin.WrapH(user))
	err := r.Run(":8989")

	if err != nil {
		logger.Error("Start failed!")
	}
}
