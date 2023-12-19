package workspace

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"strings"

	dao "github.com/deamgo/workbench/dao/workspace"
	"github.com/deamgo/workbench/pkg/logger"
	"github.com/pkg/errors"
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

// WorkspaceCreate todo the token get developer id set developer_workspace_reaction
func (w workspaceService) WorkspaceCreate(ctx context.Context, workspace *Workspace) (*Workspace, error) {
	var err error
	workspace.Id = hashTop6(workspace.Name)
	workspace.Lable = strings.Split(workspace.Lable, "\n")[0]

	err = equalParameterLen(workspace.Logo, 1, 50)
	if err != nil {
		return nil, err
	}
	err = equalParameterLen(workspace.Name, 1, 20)
	if err != nil {
		return nil, err
	}
	err = equalParameterLen(workspace.Lable, 0, 50)
	if err != nil {
		return nil, err
	}
	err = equalParameterLen(workspace.Description, 0, 1023)
	if err != nil {
		return nil, err
	}
	err = equalParameterIMG(workspace.Logo)
	if err != nil {
		return nil, err
	}

	workspaceDo := convertWorkspaceDao(workspace)

	developIdStr := ctx.Value("username").(string)
	i, err := strconv.ParseInt(developIdStr, 10, 64)
	if err != nil {
		return nil, err
	}
	workspaceDo.CreatedBy = uint64(i)
	workspaceDo.UpdatedBy = uint64(i)

	newWorkspaceDO, err := w.dao.WorkspaceCreate(ctx, workspaceDo)
	if err != nil {
		logger.Error(err)
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

// hashTop6 hash
func hashTop6(str string) string {
	hash := sha256.New()
	hash.Write([]byte(str))
	hashValue := hash.Sum(nil)
	hexString := hex.EncodeToString(hashValue)
	return hexString[:6]

}

// equalParameterLen min<=len(str)<=max
func equalParameterLen(str string, min, max int) error {
	l := len(str)
	if min <= l && l <= max {
		return nil
	}
	return errors.New("parameter exception")
}

// equalParameterIMG equal jpg、jpeg、png file
func equalParameterIMG(str string) error {
	if strings.HasSuffix(str, ".jpg") ||
		strings.HasSuffix(str, ".jpeg") ||
		strings.HasSuffix(str, ".png") {
		return nil
	}
	return errors.New("parameter exception")
}
