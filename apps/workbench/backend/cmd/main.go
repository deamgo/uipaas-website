package main

import (
	"github.com/gin-gonic/gin"
	"workbench/context"
	user2 "workbench/dao/user"
	"workbench/db"
	"workbench/pkg/logger"
	routes "workbench/router"
	"workbench/service/user"
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
