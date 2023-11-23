package user

import (
	"context"

	dao "github.com/deamgo/uipass-waitlist-page/backend/dao/user"
)

type UserService interface {
	UserGet(ctx context.Context, user *User) (*User, error)
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

func convertUserDao(user *User) *dao.UserDO {
	return &dao.UserDO{}
}

func convertUser(userDao *dao.UserDO) *User {
	return &User{}
}
