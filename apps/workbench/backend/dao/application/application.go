package application

import (
	"context"

	"gorm.io/gorm"
)

type ApplicationDo interface {
	ApplicationAdd(ctx context.Context, app *ApplicationDO) error
	ApplicationGet(ctx context.Context, app *ApplicationDO) ([]*ApplicationDO, error)
	ApplicationGetByName(ctx context.Context, app *ApplicationDO) ([]*ApplicationDO, error)
	ApplicationGetByWorkspaceIdAndName(ctx context.Context, app *ApplicationDO) (*ApplicationDO, error)
}

type applicationDao struct {
	db *gorm.DB
}

func NewApplicationDao(db *gorm.DB) ApplicationDo {
	return &applicationDao{
		db: db,
	}
}

func (d applicationDao) ApplicationAdd(ctx context.Context, app *ApplicationDO) error {
	err := d.db.WithContext(ctx).Create(&app).Error
	return err
}

func (d applicationDao) ApplicationGet(ctx context.Context, app *ApplicationDO) ([]*ApplicationDO, error) {
	var appArr []*ApplicationDO
	err := d.db.WithContext(ctx).Where("workspace_id = ? and is_deleted=0", app.WorkspaceID).Find(&appArr).Error
	return appArr, err
}

func (d applicationDao) ApplicationGetByName(ctx context.Context, app *ApplicationDO) ([]*ApplicationDO, error) {
	var appArr []*ApplicationDO

	err := d.db.WithContext(ctx).Where("workspace_id = ? and name like ? and is_deleted=0", app.WorkspaceID, "%"+app.Name+"%").
		Find(&appArr).Error
	return appArr, err
}

func (d applicationDao) ApplicationGetByWorkspaceIdAndName(ctx context.Context, app *ApplicationDO) (*ApplicationDO, error) {
	err := d.db.WithContext(ctx).Where("workspace_id = ? and name = ?  and is_deleted=0", app.WorkspaceID, app.Name).Find(&app).Error
	return app, err
}
