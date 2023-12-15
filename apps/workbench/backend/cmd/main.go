package main

import (
	"github.com/deamgo/workbench/initialize"
	"github.com/deamgo/workbench/service/mail"
	"github.com/gin-gonic/gin"

	"github.com/deamgo/workbench/context"
	user2 "github.com/deamgo/workbench/dao/developer"
	"github.com/deamgo/workbench/db"
	"github.com/deamgo/workbench/pkg/logger"
	routes "github.com/deamgo/workbench/router"
	"github.com/deamgo/workbench/service/developer"
)

func main() {

	initialize.InitConfig()
	db.InitDB()
	db.InitRedis()
	dao := user2.NewADeveloperDao(db.DB)
	ctx := context.ApplicationContext{
		UserService: developer.NewDeveloperService(developer.DeveloperServiceParams{Dao: dao, MailService: mail.NewMailService()}),
		MailService: mail.NewMailService(),
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
