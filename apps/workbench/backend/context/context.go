package context

import (
	"github.com/deamgo/workbench/service/developer"
	"github.com/deamgo/workbench/service/mail"
)

type ApplicationContext struct {
	UserService developer.UserService
	MailService mail.MailService
}
