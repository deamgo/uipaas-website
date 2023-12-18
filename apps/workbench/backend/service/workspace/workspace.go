package workspace

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	dao "github.com/deamgo/workbench/dao/workspace"
	"github.com/deamgo/workbench/pkg/logger"
)

type WorkspaceService interface {
	WorkspaceCreate(ctx context.Context, workspace *Workspace) (*Workspace, error)
}

type WorkspaceServiceParams struct {
	Dao dao.WorkspaceDao
}

type workspaceService struct {
	dao dao.WorkspaceDao
}

func NewWorkspaceService(params WorkspaceServiceParams) WorkspaceService {
	return &workspaceService{
		dao: params.Dao,
	}
}

func (w workspaceService) WorkspaceCreate(ctx context.Context, workspace *Workspace) (*Workspace, error) {

	workspace.Id = hashTop6(workspace.Name)
	workspaceDo := convertWorkspaceDao(workspace)

	newWorkspaceDO, err := w.dao.WorkspaceCreate(ctx, workspaceDo)
	if err != nil {
		logger.LoggersObj.Error(err)
		return nil, err
	}

	return convertWorkspace(newWorkspaceDO), nil
}

func convertWorkspaceDao(workspace *Workspace) *dao.WorkspaceDO {
	return &dao.WorkspaceDO{
		Id:          workspace.Id,
		Name:        workspace.Name,
		Logo:        workspace.Logo,
		Lable:       workspace.Lable,
		Description: workspace.Description,
	}
}

func convertWorkspace(workspaceDao *dao.WorkspaceDO) *Workspace {
	return &Workspace{
		Id:          workspaceDao.Id,
		Name:        workspaceDao.Name,
		Lable:       workspaceDao.Lable,
		Description: workspaceDao.Description,
		Logo:        workspaceDao.Logo,
	}
}

func hashTop6(str string) string {
	hash := sha256.New()
	hash.Write([]byte(str))
	hashValue := hash.Sum(nil)
	hexString := hex.EncodeToString(hashValue)
	return hexString[:6]
}
