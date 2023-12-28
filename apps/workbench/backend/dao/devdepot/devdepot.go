package devdepot

import (
	"context"

	"gorm.io/gorm"
)

type DevDepotDo interface {
	DevInfoList(ctx context.Context, dwr *DeveloperWorkspaceRelationDO, pageSize, pageNum int) (*[]DevDepotItem, error)
	DevInfoGetByNameLike(ctx context.Context, dwr *DeveloperWorkspaceRelationDO, name string, pageSize, pageNum int) (*[]DevDepotItem, error)
	DevDepotDel(ctx context.Context, dwr *DeveloperWorkspaceRelationDO) error
	DevDepotRoleModify(ctx context.Context, dwr *DeveloperWorkspaceRelationDO) error
	DevDepotAdd(ctx context.Context, dwr *DeveloperWorkspaceRelationDO) error
	GetCreatorByWokSpID(ctx context.Context, workspaceID string) (uint64, error)
	DevDepotGetByEmail(ctx context.Context, dwr *DeveloperWorkspaceRelationDO) (*DevDepotItem, error)
}

type devDepotDao struct {
	db *gorm.DB
}

func NewDevDepotDao(db *gorm.DB) DevDepotDo {
	return &devDepotDao{
		db: db,
	}
}

func (d devDepotDao) DevInfoList(ctx context.Context, dwr *DeveloperWorkspaceRelationDO, pageSize, pageNum int) (*[]DevDepotItem, error) {
	var devDepotItems *[]DevDepotItem
	err := d.db.Debug().WithContext(ctx).Table("workspace_developer_relation wdr").
		Joins("JOIN developer d ON d.id = wdr.developer_id").
		Where("wdr.workspace_id = ? and is_deleted=0", dwr.WorkspaceId).
		Select("wdr.email,d.username, wdr.role ,wdr.status").
		Limit(pageSize).Offset(pageSize * (pageNum - 1)).
		Find(&devDepotItems).Error
	return devDepotItems, err
}

func (d devDepotDao) DevInfoGetByNameLike(ctx context.Context, dwr *DeveloperWorkspaceRelationDO, name string, pageSize, pageNum int) (*[]DevDepotItem, error) {
	var devDepotItems *[]DevDepotItem
	err := d.db.Debug().WithContext(ctx).Table("workspace_developer_relation wdr").
		Joins("JOIN developer d ON d.id = wdr.developer_id").
		Where("wdr.workspace_id = ? and d.username like ?  and is_deleted=0", dwr.WorkspaceId, "%"+name+"%").
		Select("wdr.email,d.username, wdr.role ,wdr.status").
		Limit(pageSize).Offset(pageSize * (pageNum - 1)).
		Find(&devDepotItems).Error
	return devDepotItems, err
}

func (d devDepotDao) DevDepotDel(ctx context.Context, dwr *DeveloperWorkspaceRelationDO) error {
	err := d.db.WithContext(ctx).Model(dwr).
		Where(" developer_id = ?", dwr.DeveloperId).
		Update("is_deleted", 1).Error
	return err
}

func (d devDepotDao) DevDepotRoleModify(ctx context.Context, dwr *DeveloperWorkspaceRelationDO) error {
	err := d.db.WithContext(ctx).Model(dwr).
		Where(" developer_id = ?", dwr.DeveloperId).
		Update("role", dwr.Role).Error
	return err
}

func (d devDepotDao) DevDepotAdd(ctx context.Context, dwr *DeveloperWorkspaceRelationDO) error {
	err := d.db.WithContext(ctx).Model(&DeveloperWorkspaceRelationDO{}).Create(dwr).Error
	return err
}

func (d devDepotDao) GetCreatorByWokSpID(ctx context.Context, workspaceID string) (uint64, error) {
	var wdr DeveloperWorkspaceRelationDO
	err := d.db.Debug().WithContext(ctx).Table("workspace_developer_relation wdr").
		Where("wdr.workspace_id = ? and wdr.role=0 and is_deleted=0", workspaceID).
		Select("wdr.created_by").
		Find(&wdr).Error
	return wdr.CreatedBy, err

}

func (d devDepotDao) DevDepotGetByEmail(ctx context.Context, dwr *DeveloperWorkspaceRelationDO) (*DevDepotItem, error) {
	var devDepotItem *DevDepotItem
	err := d.db.WithContext(ctx).Model(&DeveloperWorkspaceRelationDO{}).Where("email = ?", dwr.Email).Find(&devDepotItem).Error
	return devDepotItem, err
}
