package devdepot

import (
	"context"

	"gorm.io/gorm"
)

type DevDepotDo interface {
	DevInfoList(ctx context.Context, dwr *DeveloperWorkspaceRelationDO, pageSize, pageNum int) (*[]DevDepotItem, error)
	DevInfoGetByNameLike(ctx context.Context, dwr *DeveloperWorkspaceRelationDO, name string, pageSize, pageNum int) (*[]DevDepotItem, error)
	DevDepotDel(ctx context.Context, dwr *DeveloperWorkspaceRelationDO) error
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
	err := d.db.WithContext(ctx).Table("workspace_developer_relation wdr").
		Joins("JOIN developer d ON d.id = wdr.developer_id").
		Where("wdr.workspace_id = ? and is_deleted=0", dwr.WorkspaceId).
		Select("d.email,d.username, wdr.role").
		Limit(pageSize).Offset(pageSize * (pageNum - 1)).
		Find(&devDepotItems).Error
	return devDepotItems, err
}

func (d devDepotDao) DevInfoGetByNameLike(ctx context.Context, dwr *DeveloperWorkspaceRelationDO, name string, pageSize, pageNum int) (*[]DevDepotItem, error) {
	var devDepotItems *[]DevDepotItem
	err := d.db.Debug().WithContext(ctx).Table("workspace_developer_relation wdr").
		Joins("JOIN developer d ON d.id = wdr.developer_id").
		Where("wdr.workspace_id = ? and d.username like ?  and is_deleted=0", dwr.WorkspaceId, "%"+name+"%").
		Select("d.email,d.username, wdr.role").
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
