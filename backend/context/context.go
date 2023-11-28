package context

import (
	"github.com/deamgo/uipaas-home/backend/service/companyinfo"
	"github.com/deamgo/uipaas-home/backend/service/user"
)

type ApplicationContext struct {
	UserService        user.UserService
	CompanyInfoService companyinfo.CompanyInfoService
}
