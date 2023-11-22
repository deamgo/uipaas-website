package context

import "github.com/deamgo/uipass-waitlist-page/backend/service/user"

type ApplicationContext struct {
	UserService user.UserService
}
