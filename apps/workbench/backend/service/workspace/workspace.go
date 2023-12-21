package workspace

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strings"

	dao "github.com/deamgo/workbench/dao/workspace"
	"github.com/deamgo/workbench/pkg/logger"
	"github.com/pkg/errors"
)

type WorkspaceService interface {
	WorkspaceCreate(ctx context.Context, workspace *Workspace) (*Workspace, error)
	WorkspaceGetListById(ctx context.Context, developerId uint64) ([]*Workspace, error)
	WorkspaceGetFilePath(file *os.File) (string, error)
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
	workspace.Id = hashTop(workspace.Name, 6)
	workspace.Label = strings.Split(workspace.Label, "\n")[0]

	err = equalParameterLen(workspace.Logo, 1, 50)
	if err != nil {
		return nil, err
	}
	err = equalParameterLen(workspace.Name, 1, 20)
	if err != nil {
		return nil, err
	}
	err = equalParameterLen(workspace.Label, 0, 50)
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
		Label:       workspace.Label,
		Description: workspace.Description,
		CreatedBy:   workspace.CreatedBy,
		UpdatedBy:   workspace.UpdateBy,
	}
}

func convertWorkspace(workspaceDao *dao.WorkspaceDO) *Workspace {
	return &Workspace{
		Id:          workspaceDao.Id,
		Name:        workspaceDao.Name,
		Label:       workspaceDao.Label,
		Description: workspaceDao.Description,
		Logo:        workspaceDao.Logo,
		CreatedBy:   workspaceDao.CreatedBy,
		UpdateBy:    workspaceDao.UpdatedBy,
	}
}

func hashTop(str string, length int) string {
	hash := sha256.New()
	hash.Write([]byte(str))
	hashValue := hash.Sum(nil)
	hexString := hex.EncodeToString(hashValue)
	return hexString[:length]
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
