package user

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

var UserNotExistError = errors.New("user not exist")

type UserDao interface {
	UserGet(ctx context.Context, user *UserDO) (*UserDO, error)
}

type userDao struct {
	db *gorm.DB
}

func NewAUserDao(db *gorm.DB) UserDao {
	return &userDao{
		db: db,
	}
}

func (dao *userDao) UserGet(ctx context.Context, user *UserDO) (*UserDO, error) {
	return user, nil
}
