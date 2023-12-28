package devdepot

import (
	"context"
	"errors"
	"strconv"

	"github.com/deamgo/workbench/dao/devdepot"
	developerDO "github.com/deamgo/workbench/dao/developer"
	"github.com/deamgo/workbench/pkg/consts"
	"github.com/deamgo/workbench/service/mail"

	"gorm.io/gorm"
)

type DevDepotService interface {
	DevItemList(ctx context.Context, workspaceID string, pageNum int) (*[]devdepot.DevDepotItem, error)
	DevInfoSearch(ctx context.Context, workspaceID, name string, pageNum int) (*[]devdepot.DevDepotItem, error)
	DevDepotDel(ctx context.Context, workspaceID, developerID string) error
	DevDepotRoleModify(ctx context.Context, devdepot *devdepot.DevDepotItem) error
	DevDepotInvite(ctx context.Context, item *devdepot.DevDepotItem) error
}

type DevDepotServiceParams struct {
	Dao          devdepot.DevDepotDo
	DeveloperDao developerDO.DeveloperDao
	MailService  mail.MailService
}

type devDepotService struct {
	dao          devdepot.DevDepotDo
	developerDao developerDO.DeveloperDao
	mailService  mail.MailService
}

func NewDepotService(params DevDepotServiceParams) DevDepotService {
	return &devDepotService{
		dao:          params.Dao,
		developerDao: params.DeveloperDao,
		mailService:  params.MailService,
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

func (d devDepotService) DevDepotDel(ctx context.Context, workspaceID, developerID string) error {
	dwr := &devdepot.DeveloperWorkspaceRelationDO{
		DeveloperId: developerID,
		WorkspaceId: workspaceID,
	}
	err := d.dao.DevDepotDel(ctx, dwr)
	return err
}

func (d devDepotService) DevDepotRoleModify(ctx context.Context, devdepot *devdepot.DevDepotItem) error {
	dwr, err := convertDevDepot(devdepot)
	if err != nil {
		return err
	}
	err = d.dao.DevDepotRoleModify(ctx, dwr)
	return err
}

func (d devDepotService) DevDepotInvite(ctx context.Context, item *devdepot.DevDepotItem) error {
	// Determine if you are already in the workspace
	_, err := d.dao.DevDepotGetByEmail(ctx, &devdepot.DeveloperWorkspaceRelationDO{Email: item.Email, WorkspaceId: item.WorkspaceId})
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("the developer is already in the workspace")
	}
	// Sending mailbox
	err = d.mailService.SendWorkspaceInviteMail(ctx, item.Email, item.WorkspaceId)
	if err != nil {
		return err
	}
	developer, err := d.developerDao.DeveloperGetByEmail(ctx, &developerDO.DeveloperDO{Email: item.Email})
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}
	item.Status = 0

	if developer != nil {
		item.DeveloperId = developer.ID
	}
	dwr, err := convertDevDepot(item)
	if err != nil {
		return err
	}
	creator, err := d.getCreator(ctx, item.WorkspaceId)
	if err != nil {
		return err
	}
	dwr.CreatedBy = creator
	dwr.UpdatedBy = creator
	err = d.dao.DevDepotAdd(ctx, dwr)
	return err
}

func (d devDepotService) getCreator(ctx context.Context, workspaceID string) (uint64, error) {
	return d.dao.GetCreatorByWokSpID(ctx, workspaceID)
}

// -----
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

func convertDevDepot(devInfo *devdepot.DevDepotItem) (*devdepot.DeveloperWorkspaceRelationDO, error) {
	devInfo = convertRole(devInfo)
	role, err := strconv.Atoi(devInfo.Role)
	if err != nil {
		return nil, err
	}
	return &devdepot.DeveloperWorkspaceRelationDO{
		WorkspaceId: devInfo.WorkspaceId,
		DeveloperId: devInfo.DeveloperId,
		Email:       devInfo.Email,
		Role:        uint8(role),
		Status:      devInfo.Status,
	}, nil
}
