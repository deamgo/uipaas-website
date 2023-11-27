package user

import (
	"context"
	dao "github.com/deamgo/uipass-waitlist-page/backend/dao/user"
	"github.com/deamgo/uipass-waitlist-page/backend/pkg/log"
	"go.uber.org/zap"
)

type UserService interface {
	UserGet(ctx context.Context, user *User) (*User, error)
	UserLogin(ctx context.Context, user *User) error
}

type UserServiceParams struct {
	Dao dao.UserDao
}

type userService struct {
	dao dao.UserDao
}

func NewUserService(params UserServiceParams) UserService {
	return &userService{
		dao: params.Dao,
	}
}

func (u userService) UserGet(ctx context.Context, user *User) (*User, error) {
	userdao := convertUserDao(user)
	userDO, err := u.dao.UserGet(ctx, userdao)
	if err != nil {
		return nil, err
	}

	return convertUser(userDO), nil
}

func (u userService) UserLogin(ctx context.Context, user *User) error {

	userdao := convertUserDao(user)

	err := u.dao.UserLogin(ctx, userdao)
	if err != nil {
		log.Errorw("user login failed",
			zap.Error(err),
			zap.Any("userlogin", user),
		)
		return err
	}

	if err != nil {
		return err
	}
	return nil
}

func convertUserDao(user *User) *dao.UserDO {
	return &dao.UserDO{
		UserID:   user.UserID,
		UserName: user.UserName,
		Email:    user.Email,
		Password: user.Password,
	}
}

func convertUser(userDao *dao.UserDO) *User {
	return &User{
		UserID:   userDao.UserID,
		UserName: userDao.UserName,
		Email:    userDao.Email,
		Password: userDao.Password,
	}
}
