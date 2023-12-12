package context

import "workbench/service/user"

type ApplicationContext struct {
	UserService user.UserService
}
