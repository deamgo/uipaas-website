package application

import (
	"context"
	"errors"

	"github.com/deamgo/workbench/dao/application"
	"github.com/deamgo/workbench/pkg/logger"

	"github.com/bwmarrin/snowflake"
	"gorm.io/gorm"
)

type ApplicationService interface {
	ApplicationAdd(ctx context.Context, app *Application) error
	ApplicationDuplicate(ctx context.Context, app *Application) error
	ApplicationList(ctx context.Context, app *Application) ([]*Application, error)
	ApplicationSearchByName(ctx context.Context, app *Application) ([]*Application, error)
}

type ApplicationServiceParams struct {
	Dao application.ApplicationDo
}

type applicationService struct {
	dao application.ApplicationDo
}

func NewApplicationService(params ApplicationServiceParams) ApplicationService {
	return &applicationService{
		dao: params.Dao,
	}
}

func (a applicationService) ApplicationAdd(ctx context.Context, app *Application) error {
	findApp, err := a.dao.ApplicationGetByWorkspaceIdAndName(ctx, convertApplicationDo(app))
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}
	if findApp.ID != "" {
		return errors.New("application name duplication")
	}
	node, err := snowflake.NewNode(2)
	if err != nil {
		logger.Error(err)
		return err
	}
	generateID := node.Generate().String()
	application := &application.ApplicationDO{
		ID:          generateID,
		Name:        app.Name,
		WorkspaceID: app.WorkspaceID,
		CreatedBy:   app.CreatedBy,
		UpdatedBy:   app.CreatedBy,
		Description: "a application",
		Status:      0,
	}
	err = a.dao.ApplicationAdd(ctx, application)
	return err
}

func (a applicationService) ApplicationList(ctx context.Context, app *Application) ([]*Application, error) {
	appDo := convertApplicationDo(app)
	appDoArr, err := a.dao.ApplicationGet(ctx, appDo)
	var appArr []*Application
	for _, v := range appDoArr {
		appArr = append(appArr, convertApplication(v))
	}
	return appArr, err
}

func (a applicationService) ApplicationSearchByName(ctx context.Context, app *Application) ([]*Application, error) {
	appDo := convertApplicationDo(app)
	appDoArr, err := a.dao.ApplicationGetByName(ctx, appDo)
	var appArr []*Application
	for _, v := range appDoArr {
		appArr = append(appArr, convertApplication(v))
	}
	return appArr, err
}

func (a applicationService) ApplicationDuplicate(ctx context.Context, app *Application) error {
	findApp, err := a.dao.ApplicationGetByWorkspaceIdAndName(ctx, convertApplicationDo(app))
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}
	if findApp.ID != "" {
		return errors.New("application name duplication")
	}
	node, err := snowflake.NewNode(2)
	if err != nil {
		logger.Error(err)
	}
	generateID := node.Generate().String()
	application := &application.ApplicationDO{
		ID:          generateID,
		Name:        app.Name,
		WorkspaceID: app.WorkspaceID,
		CreatedBy:   app.CreatedBy,
		UpdatedBy:   app.CreatedBy,
		Description: app.Description,
		Status:      0,
	}
	err = a.dao.ApplicationAdd(ctx, application)
	return err
}

func convertApplication(a *application.ApplicationDO) *Application {
	return &Application{
		ID:          a.ID,
		Name:        a.Name,
		WorkspaceID: a.WorkspaceID,
		Description: a.Description,
		Status:      a.Status,
		Icon:        a.Icon,
		CreatedBy:   a.CreatedBy,
		DeletedBy:   a.DeletedBy,
	}
}

func convertApplicationDo(a *Application) *application.ApplicationDO {
	return &application.ApplicationDO{
		ID:          a.ID,
		Name:        a.Name,
		WorkspaceID: a.WorkspaceID,
		Description: a.Description,
		Status:      a.Status,
		Icon:        a.Icon,
		CreatedBy:   a.CreatedBy,
		DeletedBy:   a.DeletedBy,
	}
}
