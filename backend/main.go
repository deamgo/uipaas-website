package main

import (
	"github.com/deamgo/uipass-waitlist-page/backend/context"
	dao "github.com/deamgo/uipass-waitlist-page/backend/dao/user"
	"github.com/deamgo/uipass-waitlist-page/backend/db"
	"github.com/deamgo/uipass-waitlist-page/backend/router"
	"github.com/deamgo/uipass-waitlist-page/backend/service/user"
	"github.com/gin-gonic/gin"
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
		return
	}

}
