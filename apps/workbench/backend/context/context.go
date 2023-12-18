package context

import (
	"github.com/deamgo/workbench/service/user"
	"github.com/deamgo/workbench/service/workspace"
)

type ApplicationContext struct {
	UserService      user.UserService
	WorkspaceService workspace.WorkspaceService
}
