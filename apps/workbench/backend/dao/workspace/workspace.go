package workspace

import (
	"context"

	"gorm.io/gorm"
)

type WorkspaceDao interface {
	WorkspaceCreate(ctx context.Context, workspace *WorkspaceDO) (*WorkspaceDO, error)
}

type workspaceDao struct {
	db *gorm.DB
}

func NewWorkspaceDao(db *gorm.DB) WorkspaceDao {
	return &workspaceDao{
		db: db,
	}
}

func (dao workspaceDao) WorkspaceCreate(ctx context.Context, workspace *WorkspaceDO) (*WorkspaceDO, error) {
	err := dao.db.WithContext(ctx).Model(&WorkspaceDO{}).Create(&workspace).Error
	return workspace, err
}
