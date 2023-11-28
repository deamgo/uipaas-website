package user

import (
	"context"
	"github.com/pkg/errors"

	daolayer "github.com/deamgo/uipass-waitlist-page/backend/dao"

	"gorm.io/gorm"
)

var UserNotExistError = errors.New("user not exist")
var UserLoginError = errors.New("incorrect username or password")

type UserDao interface {
	UserGet(ctx context.Context, user *UserDO) (*UserDO, error)
	UserLogin(ctx context.Context, user *UserDO) error
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
	id := user.UserID
	if err := dao.db.WithContext(ctx).Model(&user).Where("id = ?", id).First(&user).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, UserNotExistError
		}
		return nil, errors.Wrap(daolayer.DBError, err.Error())
	}

	return user, nil
}

func (dao *userDao) UserLogin(ctx context.Context, user *UserDO) error {

	if err := dao.db.Where("username = ? AND password = ?", user.UserName, user.Password).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return UserLoginError
		}
		return errors.Wrap(daolayer.DBError, err.Error())

	}
	return nil
}
