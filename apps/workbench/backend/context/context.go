package context

import (
	"github.com/deamgo/workbench/service/developer"
	"github.com/deamgo/workbench/service/mail"
	"github.com/deamgo/workbench/service/workspace"
)

type ApplicationContext struct {
	UserService      developer.UserService
	MailService      mail.MailService
	WorkspaceService workspace.WorkspaceService
}
