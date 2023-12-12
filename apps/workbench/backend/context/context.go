package context

import "github.com/deamgo/workbench/service/user"

type ApplicationContext struct {
	UserService user.UserService
}
