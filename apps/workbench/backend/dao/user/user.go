// Package user provides data access objects for the application.
package user

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	daolayer "github.com/deamgo/workbench/dao"
	"github.com/deamgo/workbench/db"
)

var UserNotExistError = errors.New("user not exist")
var UserLoginError = errors.New("incorrect username or password")

type UserDao interface {
	UserAdd(ctx context.Context, user *UserDO) (uint, error)
	UserGetByEmail(ctx context.Context, user *UserDO) (*UserDO, error)
	UserGetByInvitationCode(ctx context.Context, user *UserDO) (*UserDO, error)
	UserGetByUserName(ctx context.Context, user *UserDO) (*UserDO, error)
	UserDeactivateModifyByEmail(ctx context.Context, user *UserDO) error
}

type userDao struct {
	db *gorm.DB
}

func NewAUserDao(db *gorm.DB) UserDao {
	return &userDao{
		db: db,
	}
}

func (u userDao) UserAdd(ctx context.Context, user *UserDO) (uint, error) {
	err := db.DB.WithContext(ctx).Model(&UserDO{}).Create(&user).Error
	uId := user.UID
	return uId, err
}

// Update deactivate
func (u userDao) UserDeactivateModifyByEmail(ctx context.Context, user *UserDO) error {
	email := user.Email
	deactivate := user.Deactivate
	err := db.DB.WithContext(ctx).Model(&user).
		Where("email=?", email).UpdateColumn("deactivate", deactivate).Error
	if err != nil {
		return errors.Wrap(daolayer.DBError, err.Error())
	}
	return nil
}

// Search by email
func (u userDao) UserGetByEmail(ctx context.Context, user *UserDO) (*UserDO, error) {
	email := user.Email
	err := db.DB.WithContext(ctx).Model(&user).Where("email=?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, UserNotExistError
		}
		return nil, errors.Wrap(daolayer.DBError, err.Error())
	}
	return user, nil
}

// Search by username
func (u userDao) UserGetByUserName(ctx context.Context, user *UserDO) (*UserDO, error) {
	uname := user.Username
	err := db.DB.WithContext(ctx).Model(&user).Where("username=?", uname).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, UserNotExistError
		}
		return nil, errors.Wrap(daolayer.DBError, err.Error())
	}
	return user, nil

}

// Search by invitation code
func (u userDao) UserGetByInvitationCode(ctx context.Context, user *UserDO) (*UserDO, error) {
	invitationCode := user.InvitationCode
	err := db.DB.WithContext(ctx).
		Model(&user).
		Where("invitation_code=?", invitationCode).
		First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, UserNotExistError
		}
		return nil, errors.Wrap(daolayer.DBError, err.Error())
	}

	return user, nil
}
