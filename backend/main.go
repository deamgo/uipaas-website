package main

import (
	"github.com/deamgo/uipaas-home/backend/context"
	formdao "github.com/deamgo/uipaas-home/backend/dao/companyinfo"
	dao "github.com/deamgo/uipaas-home/backend/dao/user"
	"github.com/deamgo/uipaas-home/backend/db"
	"github.com/deamgo/uipaas-home/backend/pkg/log"
	"github.com/deamgo/uipaas-home/backend/router"
	"github.com/deamgo/uipaas-home/backend/service/companyinfo"
	"github.com/deamgo/uipaas-home/backend/service/user"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	dao := dao.NewAUserDao(db.DB)
	formdao := formdao.NewAUseFormDao(db.DB)
	ctx := context.ApplicationContext{UserService: user.NewUserService(
		user.UserServiceParams{Dao: dao},
	),
		CompanyInfoService: companyinfo.NewCompanyInfoService(companyinfo.CompanyInfoServiceParams{Dao: formdao})}

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
