package context

import (
	"github.com/deamgo/workbench/service/application"
	"github.com/deamgo/workbench/service/devdepot"
	"github.com/deamgo/workbench/service/developer"
	"github.com/deamgo/workbench/service/mail"
	"github.com/deamgo/workbench/service/workspace"
)

type ApplicationContext struct {
	UserService        developer.UserService
	MailService        mail.MailService
	WorkspaceService   workspace.WorkspaceService
	DevDepotService    devdepot.DevDepotService
	ApplicationService application.ApplicationService
}
