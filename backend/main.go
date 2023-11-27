package main

import (
	"github.com/deamgo/uipaas-home/backend/context"
	dao "github.com/deamgo/uipaas-home/backend/dao/user"
	"github.com/deamgo/uipaas-home/backend/db"
	"github.com/deamgo/uipaas-home/backend/pkg/log"
	"github.com/deamgo/uipaas-home/backend/router"
	"github.com/deamgo/uipaas-home/backend/service/user"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	dao := dao.NewAUserDao(db.DB)
	ctx := context.ApplicationContext{UserService: user.NewUserService(
		user.UserServiceParams{Dao: dao},
	)}

	r := gin.Default()
	user := router.NewRouter(ctx)
	r.Any("/*any", gin.WrapH(user))
	err := r.Run(":8080")
	if err != nil {
		log.Fatalw("gin run error",
			zap.Any("error message: ", err.Error()),
		)
	}

}
