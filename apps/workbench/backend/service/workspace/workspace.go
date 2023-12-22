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

type UploadFileResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Path string `json:"path"`
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

func (w workspaceService) WorkspaceGetFilePath(file *os.File) (string, error) {
	var err error
	var data UploadFileResp
	//err = equalParameterIMG(file.Name())
	if err != nil {
		return "", err
	}

	ext := path.Ext(file.Name())

	form := new(bytes.Buffer)
	writer := multipart.NewWriter(form)

	newFileName := hashTop(file.Name(), 10) + ext
	fw, err := writer.CreateFormFile("file", newFileName)
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(fw, file)
	if err != nil {
		log.Fatal(err)
	}
	writer.Close()

	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://121.41.78.218:8700/api/v1/extract-public/upload", form)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "Apifox/1.0.0 (https://apifox.com)")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Host", "121.41.78.218:8700")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", writer.FormDataContentType())
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
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
	err = equalParameterIMG(workspace.Logo)
	if err != nil {
		workspace.Logo = ""
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
		strings.HasSuffix(str, ".png") ||
		len(str) == 0 {
		return nil
	}
	return errors.New("parameter exception")
}
