package context

import "github.com/deamgo/uipaas-home/backend/service/user"

type ApplicationContext struct {
	UserService user.UserService
}
