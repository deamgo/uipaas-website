package workspace

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"path"
	"strconv"
	"strings"

	developerDO "github.com/deamgo/workbench/dao/developer"
	dao "github.com/deamgo/workbench/dao/workspace"
	"github.com/deamgo/workbench/pkg/logger"

	"github.com/pkg/errors"
)

type WorkspaceService interface {
	WorkspaceCreate(ctx context.Context, workspace *Workspace) (*Workspace, error)
	WorkspaceDefaultCreate(ctx context.Context, workspace *Workspace) (*Workspace, error)
	WorkspaceDel(ctx context.Context, workspace *Workspace, developerID string) error
	WorkspaceGetListById(ctx context.Context, developerId uint64) ([]*Workspace, error)
	WorkspaceGetFilePath(file *multipart.FileHeader) (string, error)
	WorkspaceNameModify(ctx context.Context, workspace *Workspace) error
}

type UploadFileResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Path string `json:"path"`
}

type WorkspaceServiceParams struct {
	Dao          dao.WorkspaceDao
	DeveloperDao developerDO.DeveloperDao
}

type workspaceService struct {
	dao          dao.WorkspaceDao
	developerDao developerDO.DeveloperDao
}

func NewWorkspaceService(params WorkspaceServiceParams) WorkspaceService {
	return &workspaceService{
		dao:          params.Dao,
		developerDao: params.DeveloperDao,
	}
}

func (w workspaceService) WorkspaceGetFilePath(header *multipart.FileHeader) (string, error) {
	var err error
	var data UploadFileResp

	err = equalParameterIMG(header.Filename)
	if err != nil {
		return "", err
	}
	ext := path.Ext(header.Filename)
	newFileName := hashTop(header.Filename, 10) + ext

	file, err := header.Open()
	if err != nil {
		return "", err
	}

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	fw, err := writer.CreateFormFile("file", newFileName)
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(fw, file)
	if err != nil {
		log.Fatal(err)
	}
	writer.Close()

	resp, err := http.Post("http://121.41.78.218:8700/api/v1/extract-public/upload", writer.FormDataContentType(), body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(bodyText, &data)
	if err != nil {
		log.Fatal(err)
	}
	return data.Path, nil
}

func (w workspaceService) WorkspaceGetListById(ctx context.Context, developerId uint64) ([]*Workspace, error) {
	var err error
	var workspaces []*Workspace

	newWorkspaceDOs, err := w.dao.WorkspaceGetListById(ctx, developerId)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	for _, workspaceDO := range newWorkspaceDOs {
		workspaces = append(workspaces, convertWorkspace(workspaceDO))
	}

	return workspaces, nil
}

// WorkspaceCreate
func (w workspaceService) WorkspaceCreate(ctx context.Context, workspace *Workspace) (*Workspace, error) {
	var err error

	workspace.Id = hashTop(workspace.Name, 6)
	workspace.Label = strings.Split(workspace.Label, "\n")[0]

	err = equalParameterIMG(workspace.Logo)
	if err != nil {
		workspace.Logo = ""
	}
	err = equalParameterLen(workspace.Logo, 0, 50)
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

	workspaceDo := convertWorkspaceDao(workspace)
	developer := &developerDO.DeveloperDO{ID: strconv.FormatUint(workspace.CreatedBy, 10)}
	d, err := w.developerDao.DeveloperGetByID(ctx, developer)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	newWorkspaceDO, err := w.dao.WorkspaceCreate(ctx, workspaceDo, d.Email)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return convertWorkspace(newWorkspaceDO), nil
}

// WorkspaceCreate
func (w workspaceService) WorkspaceDefaultCreate(ctx context.Context, workspace *Workspace) (*Workspace, error) {
	var err error

	workspace.Id = hashTop(workspace.Name, 6)
	workspace.Label = strings.Split(workspace.Label, "\n")[0]

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

	workspaceDo := convertWorkspaceDao(workspace)
	developer := &developerDO.DeveloperDO{ID: strconv.FormatUint(workspace.CreatedBy, 10)}
	d, err := w.developerDao.DeveloperGetByID(ctx, developer)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	newWorkspaceDO, err := w.dao.WorkspaceCreate(ctx, workspaceDo, d.Email)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return convertWorkspace(newWorkspaceDO), nil
}

func (w workspaceService) WorkspaceDel(ctx context.Context, workspace *Workspace, developerID string) error {
	workspaceDo := convertWorkspaceDao(workspace)
	err := w.dao.WorkspaceDel(ctx, workspaceDo, developerID)
	return err
}

func (w workspaceService) WorkspaceNameModify(ctx context.Context, workspace *Workspace) error {
	workspaceDo := convertWorkspaceDao(workspace)
	err := w.dao.WorkspaceNameModify(ctx, workspaceDo)
	return err
}

func convertWorkspaceDao(workspace *Workspace) *dao.WorkspaceDO {
	return &dao.WorkspaceDO{
		Id:   workspace.Id,
		Name: workspace.Name,
		Logo: workspace.Logo,

		Label:       workspace.Label,
		Description: workspace.Description,
		CreatedBy:   workspace.CreatedBy,
		UpdatedBy:   workspace.UpdatedBy,
	}
}

func convertWorkspace(workspaceDao *dao.WorkspaceDO) *Workspace {
	return &Workspace{
		Id:   workspaceDao.Id,
		Name: workspaceDao.Name,

		Label:       workspaceDao.Label,
		Description: workspaceDao.Description,
		Logo:        workspaceDao.Logo,
		CreatedBy:   workspaceDao.CreatedBy,
		UpdatedBy:   workspaceDao.UpdatedBy,
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
