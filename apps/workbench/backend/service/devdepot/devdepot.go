package devdepot

import (
	"context"
	"strconv"

	"github.com/deamgo/workbench/dao/devdepot"
	"github.com/deamgo/workbench/pkg/consts"
	"github.com/deamgo/workbench/pkg/logger"
)

type DevDepotService interface {
	DevItemList(ctx context.Context, workspaceID string, pageNum int) (*[]devdepot.DevDepotItem, error)
	DevInfoSearch(ctx context.Context, workspaceID, name string, pageNum int) (*[]devdepot.DevDepotItem, error)
	DevDepotDel(ctx context.Context, developerID string) error
}

type DevDepotServiceParams struct {
	Dao devdepot.DevDepotDo
}

type devDepotService struct {
	dao devdepot.DevDepotDo
}

func NewDepotService(params DevDepotServiceParams) DevDepotService {
	return &devDepotService{
		dao: params.Dao,
	}
}

func (d devDepotService) DevItemList(ctx context.Context, workspaceID string, pageNum int) (*[]devdepot.DevDepotItem, error) {
	devSpaceRel := &devdepot.DeveloperWorkspaceRelationDO{
		WorkspaceId: workspaceID,
	}
	devInfoList, err := d.dao.DevInfoList(ctx, devSpaceRel, consts.PAGESIZE, pageNum)
	convertRoleNum(devInfoList)
	return devInfoList, err
}

func (d devDepotService) DevInfoSearch(ctx context.Context, workspaceID, name string, pageNum int) (*[]devdepot.DevDepotItem, error) {
	devSpaceRel := &devdepot.DeveloperWorkspaceRelationDO{
		WorkspaceId: workspaceID,
	}
	devInfoList, err := d.dao.DevInfoGetByNameLike(ctx, devSpaceRel, name, consts.PAGESIZE, pageNum)
	convertRoleNum(devInfoList)
	return devInfoList, err
}

func (d devDepotService) DevDepotDel(ctx context.Context, developerID string) error {
	id, err := strconv.ParseInt(developerID, 10, 64)
	if err != nil {
		logger.Error(err)
		return err
	}
	dwr := &devdepot.DeveloperWorkspaceRelationDO{
		DeveloperId: uint64(id),
	}
	err = d.dao.DevDepotDel(ctx, dwr)
	return err
}

func convertRoleNum(devInfoList *[]devdepot.DevDepotItem) *[]devdepot.DevDepotItem {
	// Map the Role value to a human-readable string
	roleMapping := map[string]string{
		"0": "owner",
		"1": "Admin",
		"2": "Developer",
		"3": "Reviewer",
	}

	for i := range *devInfoList {
		item := &(*devInfoList)[i]
		if roleName, ok := roleMapping[item.Role]; ok {
			item.Role = roleName
		}
	}
	return devInfoList

}

func convertRole(devInfo *devdepot.DevDepotItem) *devdepot.DevDepotItem {
	// Map the Role value to a human-readable string
	roleMapping := map[string]string{
		"owner":     "0",
		"Admin":     "1",
		"Developer": "2",
		"Reviewer":  "3",
	}

	if roleName, ok := roleMapping[devInfo.Role]; ok {
		devInfo.Role = roleName
	}
	return devInfo

}
