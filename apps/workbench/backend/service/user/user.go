package user

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/deamgo/workbench/dao/user"
	"github.com/deamgo/workbench/db"
	"github.com/deamgo/workbench/pkg/logger"
	"github.com/deamgo/workbench/service/mail"
)

type UserService interface {
	UserAdd(ctx context.Context, user *User) (string, error)
	UserGetByEmail(ctx context.Context, u *User) (*User, error)
	UserGetByInvitationCode(ctx context.Context, u *User) (*User, error)
	UserGetByUserName(ctx context.Context, u *User) (*User, error)
	UserDeactivateModifyByEmail(ctx context.Context, u *User) error
}

type UserServiceParams struct {
	Dao user.UserDao
}

type userService struct {
	dao user.UserDao
}

func NewUserService(params UserServiceParams) UserService {
	return &userService{
		dao: params.Dao,
	}
}

func (us userService) UserAdd(ctx context.Context, u *User) (string, error) {
	ud := convertUserDao(u)
	//send email and get  code
	code := mail.SendMail(ctx, "2734170020@qq.com")
	codeHash, err := SaveCode(u.Email, code)
	if err != nil {
		logger.LoggersObj.Error(err)
		return "", err
	}

	fmt.Println(code)
	var uId uint
	uId, err = us.dao.UserAdd(ctx, ud)
	if err != nil {
		logger.LoggersObj.Error(err)

		return "", err
	}
	fmt.Println(uId)
	return codeHash, nil
}

func (us userService) UserDeactivateModifyByEmail(ctx context.Context, u *User) error {
	ud := convertUserDao(u)
	err := us.dao.UserDeactivateModifyByEmail(ctx, ud)
	if err != nil {
		logger.LoggersObj.Error(err)

		return err
	}
	return nil
}

func (us userService) UserGetByUserName(ctx context.Context, u *User) (*User, error) {
	var err error
	ud := convertUserDao(u)
	ud, err = us.dao.UserGetByUserName(ctx, ud)
	if err != nil {
		logger.LoggersObj.Error(err)

		return nil, err
	}
	return convertUser(ud), err
}

func (us userService) UserGetByEmail(ctx context.Context, u *User) (*User, error) {
	var err error
	ud := convertUserDao(u)
	ud, err = us.dao.UserGetByEmail(ctx, ud)
	if err != nil {
		logger.LoggersObj.Error(err)

		return nil, err
	}
	return convertUser(ud), err
}
func (us userService) UserGetByInvitationCode(ctx context.Context, u *User) (*User, error) {
	var err error
	ud := convertUserDao(u)
	ud, err = us.dao.UserGetByInvitationCode(ctx, ud)
	if err != nil {
		logger.LoggersObj.Error(err)

		return nil, err
	}
	return convertUser(ud), err
}

func convertUserDao(u *User) *user.UserDO {
	return &user.UserDO{
		UID:            u.UID,
		Username:       u.Username,
		Email:          u.Email,
		Password:       u.Password,
		InvitationCode: u.InvitationCode,
	}
}

func convertUser(u *user.UserDO) *User {
	return &User{
		UID:            u.UID,
		Username:       u.Username,
		Email:          u.Email,
		Password:       u.Password,
		InvitationCode: u.InvitationCode,
	}
}

func SaveCode(email string, code int) (string, error) {

	codeHash := getEmailHashStr(email)

	err := db.RedisDB.HSet(codeHash, "code", code).Err()
	if err != nil {
		logger.LoggersObj.Error(err)
		return "", err
	}

	err = db.RedisDB.Expire(codeHash, 60*time.Second).Err()
	if err != nil {
		logger.LoggersObj.Error(err)
		return "", err
	}

	return codeHash, nil
}
func getEmailHashStr(email string) string {
	codeHash := md5.New()
	data := []byte(email)
	codeHash.Write(data)
	hashBytes := codeHash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	return hashString
}
