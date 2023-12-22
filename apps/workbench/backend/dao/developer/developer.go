// Package developer provides data access objects for the application.
package developer

import (
	"context"
	"github.com/pkg/errors"

	daolayer "github.com/deamgo/workbench/dao"

	"gorm.io/gorm"
)

type DeveloperDao interface {
	DeveloperAdd(ctx context.Context, user *DeveloperDO) error
	DeveloperGetByID(ctx context.Context, user *DeveloperDO) (*DeveloperDO, error)
	DeveloperGetByEmail(ctx context.Context, user *DeveloperDO) (*DeveloperDO, error)
	DeveloperGetByUserName(ctx context.Context, user *DeveloperDO) (*DeveloperDO, error)
	DeveloperStatusModifyByEmail(ctx context.Context, user *DeveloperDO) error
	DeveloperNameModifyByID(ctx context.Context, user *DeveloperDO) error
	DeveloperEmailModifyByEmail(ctx context.Context, oldEmail string, u *DeveloperDO) error
	DeveloperPasswordModifyByEmail(ctx context.Context, user *DeveloperDO) error
	DeveloperGetByEmailAndPwd(ctx context.Context, dlp *DeveloperDO) (dpl *DeveloperDO, err error)
}

type developerDao struct {
	db *gorm.DB
}

func NewADeveloperDao(db *gorm.DB) DeveloperDao {
	return &developerDao{
		db: db,
	}
}

func (u developerDao) DeveloperAdd(ctx context.Context, user *DeveloperDO) error {
	err := u.db.WithContext(ctx).Model(&DeveloperDO{}).Create(&user).Error
	return err
}

// Update deactivate
func (u developerDao) DeveloperStatusModifyByEmail(ctx context.Context, user *DeveloperDO) error {
	email := user.Email
	status := user.Status
	err := u.db.WithContext(ctx).Model(&user).
		Where("email=?", email).UpdateColumn("status", status).Error
	if err != nil {
		return errors.Wrap(daolayer.DBError, err.Error())
	}
	return nil
}

func (u developerDao) DeveloperPasswordModifyByEmail(ctx context.Context, user *DeveloperDO) error {
	email := user.Email
	pwd := user.Password
	err := u.db.WithContext(ctx).Model(&user).
		Where("email=?", email).UpdateColumn("password", pwd).Error
	if err != nil {
		return errors.Wrap(daolayer.DBError, err.Error())
	}
	return nil
}
func (u developerDao) DeveloperNameModifyByID(ctx context.Context, user *DeveloperDO) error {
	//id := user.ID
	uname := user.Username
	err := u.db.WithContext(ctx).Model(&user).
		UpdateColumn("username", uname).Error
	if err != nil {
		return errors.Wrap(daolayer.DBError, err.Error())
	}
	return nil
}

// Search by email
func (u developerDao) DeveloperGetByEmail(ctx context.Context, user *DeveloperDO) (*DeveloperDO, error) {
	email := user.Email
	err := u.db.WithContext(ctx).Model(&user).Where("email=? and status=1", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// DeveloperGetByUserName Search by username
func (u developerDao) DeveloperGetByUserName(ctx context.Context, user *DeveloperDO) (*DeveloperDO, error) {
	uname := user.Username
	err := u.db.WithContext(ctx).Model(&user).Where("username=?", uname).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u developerDao) DeveloperGetByID(ctx context.Context, user *DeveloperDO) (*DeveloperDO, error) {
	err := u.db.WithContext(ctx).Model(&user).Select("username, avatar, email").Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u developerDao) DeveloperGetByEmailAndPwd(ctx context.Context, dlp *DeveloperDO) (dpl *DeveloperDO, err error) {
	err = u.db.WithContext(ctx).Model(&DeveloperDO{}).Where("email=? and password=?", dlp.Email, dlp.Password).First(&dlp).Error
	if err != nil {
		return nil, err
	}
	return dlp, nil
}

func (u developerDao) DeveloperEmailModifyByEmail(ctx context.Context, oldEmail string, user *DeveloperDO) error {
	email := user.Email
	err := u.db.WithContext(ctx).Model(&user).Where("email=?", oldEmail).
		UpdateColumn("email", email).Error
	if err != nil {
		return errors.Wrap(daolayer.DBError, err.Error())
	}
	return nil
}
