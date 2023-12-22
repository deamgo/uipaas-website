package workspace

import (
	"context"

	"gorm.io/gorm"
)

type WorkspaceDao interface {
	WorkspaceCreate(ctx context.Context, workspace *WorkspaceDO) (*WorkspaceDO, error)

	WorkspaceDel(ctx context.Context, workspace *WorkspaceDO) error

	WorkspaceGetListById(ctx context.Context, developerId uint64) ([]*WorkspaceDO, error)

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
	dwrDO := &DeveloperWorkspaceRelationDO{
		WorkspaceId: workspace.Id,
		DeveloperId: workspace.CreatedBy,
		Role:        1,
		CreatedBy:   workspace.CreatedBy,
		UpdatedBy:   workspace.UpdatedBy,
	}

	err := dao.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Model(&WorkspaceDO{}).Create(&workspace).Error; err != nil {
			return err
		}
		if err := tx.Model(&DeveloperWorkspaceRelationDO{}).Create(&dwrDO).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}
	return workspace, err

}

func (dao workspaceDao) WorkspaceGetListById(ctx context.Context, developerId uint64) ([]*WorkspaceDO, error) {
	var WorkspaceDOs []*WorkspaceDO
	err := dao.db.WithContext(ctx).Debug().
		Raw("select w.* from workspace_developer_relation r left join workspaces w on w.id = r.workspace_id where developer_id = ? and w.is_deleted = 0; ", developerId).Scan(&WorkspaceDOs).Error
	if err != nil {
		return nil, err
	}
	return WorkspaceDOs, nil
}

func (dao workspaceDao) WorkspaceDel(ctx context.Context, workspace *WorkspaceDO) error {
	err := dao.db.WithContext(ctx).Model(workspace).UpdateColumn("is_deleted", 1).Error
	return err
}
