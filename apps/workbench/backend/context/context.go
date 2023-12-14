package context

import (
	"github.com/deamgo/workbench/service/mail"
	"github.com/deamgo/workbench/service/users"
)

type ApplicationContext struct {
	UserService users.UserService
	MailService mail.MailService
}
