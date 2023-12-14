// Package users provides data access objects for the application.
package user

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	daolayer "github.com/deamgo/workbench/dao"
)

type UserDao interface {
	UserAdd(ctx context.Context, user *DeveloperDO) error
	UserGetByEmail(ctx context.Context, user *DeveloperDO) (*DeveloperDO, error)
	UserGetByUserName(ctx context.Context, user *DeveloperDO) (*DeveloperDO, error)
	UserStatusModifyByEmail(ctx context.Context, user *DeveloperDO) error
	UserPasswordModifyByEmail(ctx context.Context, user *DeveloperDO) error
}

type userDao struct {
	db *gorm.DB
}

func NewAUserDao(db *gorm.DB) UserDao {
	return &userDao{
		db: db,
	}
}

func (u userDao) UserAdd(ctx context.Context, user *DeveloperDO) error {
	err := u.db.WithContext(ctx).Model(&DeveloperDO{}).Create(&user).Error
	return err
}

// Update deactivate
func (u userDao) UserStatusModifyByEmail(ctx context.Context, user *DeveloperDO) error {
	email := user.Email
	status := user.Status
	err := u.db.WithContext(ctx).Model(&user).
		Where("email=?", email).UpdateColumn("status", status).Error
	if err != nil {
		return errors.Wrap(daolayer.DBError, err.Error())
	}
	return nil
}

func (u userDao) UserPasswordModifyByEmail(ctx context.Context, user *DeveloperDO) error {
	email := user.Email
	pwd := user.Password
	err := u.db.WithContext(ctx).Model(&user).
		Where("email=?", email).UpdateColumn("password", pwd).Error
	if err != nil {
		return errors.Wrap(daolayer.DBError, err.Error())
	}
	return nil
}

// Search by email
func (u userDao) UserGetByEmail(ctx context.Context, user *DeveloperDO) (*DeveloperDO, error) {
	email := user.Email
	err := u.db.WithContext(ctx).Model(&user).Where("email=? and status=?", email, 1).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Search by username
func (u userDao) UserGetByUserName(ctx context.Context, user *DeveloperDO) (*DeveloperDO, error) {
	uname := user.Username
	err := u.db.WithContext(ctx).Model(&user).Where("username=?", uname).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil

}
