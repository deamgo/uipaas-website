package main

import (
	"github.com/gin-gonic/gin"

	"github.com/deamgo/workbench/context"
	user2 "github.com/deamgo/workbench/dao/user"
	"github.com/deamgo/workbench/db"
	"github.com/deamgo/workbench/pkg/logger"
	routes "github.com/deamgo/workbench/router"
	"github.com/deamgo/workbench/service/user"
)

func main() {

	db.InitRedis()
	dao := user2.NewAUserDao(db.DB)
	ctx := context.ApplicationContext{
		UserService: user.NewUserService(user.UserServiceParams{Dao: dao}),
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
