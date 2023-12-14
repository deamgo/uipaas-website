package users

import (
	"context"
	"crypto/sha1"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"time"

	"github.com/deamgo/workbench/dao/user"
	"github.com/deamgo/workbench/db"
	"github.com/deamgo/workbench/pkg/logger"
	"github.com/deamgo/workbench/service/mail"
)

type UserService interface {
	UserAdd(ctx context.Context, user *User) (string, error)
	UserGetByEmail(ctx context.Context, u *User) (*user.DeveloperDO, error)
	ForgotVerifySend(ctx context.Context, u *User) (string, error)
	UserGetByUserName(ctx context.Context, u *User) (*User, error)
	UserStatusModifyByEmail(ctx context.Context, u *User) error
	UserPasswordModifyByEmail(ctx context.Context, u *User) error
}

type UserServiceParams struct {
	Dao         user.UserDao
	MailService mail.MailService
}

type userService struct {
	dao  user.UserDao
	mail mail.MailService
}

func NewUserService(params UserServiceParams) UserService {
	return &userService{
		dao:  params.Dao,
		mail: params.MailService,
	}
}

func (us userService) UserAdd(ctx context.Context, u *User) (string, error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		logger.Error(err)
	}
	generateID := node.Generate().String()

	ud := convertUserDao(u)
	//send email and get  code
	code := us.mail.SendMail(ctx, u.Email)
	codeHash, err := SaveCode(u.Email, code)
	if err != nil {
		logger.Error(err)
		return "", err
	}
	fmt.Println(code)
	ud.ID = generateID
	err = us.dao.UserAdd(ctx, ud)
	if err != nil {
		logger.Error(err)

		return "", err
	}
	return codeHash, nil
}

func (us userService) UserStatusModifyByEmail(ctx context.Context, u *User) error {
	ud := convertUserDao(u)
	ud.Status = 1
	err := us.dao.UserStatusModifyByEmail(ctx, ud)
	if err != nil {
		logger.Error(err)

		return err
	}
	return nil
}

func (us userService) ForgotVerifySend(ctx context.Context, u *User) (string, error) {
	code := mail.MailService.SendMail(mail.NewMailService(), ctx, "2734170020@qq.com")
	codeHash, err := SaveCode(u.Email, code)
	if err != nil {
		logger.Error(err)
		return "", err
	}
	return codeHash, nil
}

func (us userService) UserPasswordModifyByEmail(ctx context.Context, u *User) error {
	ud := convertUserDao(u)
	err := us.dao.UserPasswordModifyByEmail(ctx, ud)
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (us userService) UserGetByUserName(ctx context.Context, u *User) (*User, error) {
	var err error
	ud := convertUserDao(u)
	ud, err = us.dao.UserGetByUserName(ctx, ud)
	if err != nil {
		logger.Error(err)

		return nil, err
	}
	return convertUser(ud), err
}

func (us userService) UserGetByEmail(ctx context.Context, u *User) (*user.DeveloperDO, error) {
	var err error
	ud := convertUserDao(u)
	ud, err = us.dao.UserGetByEmail(ctx, ud)
	if err != nil {
		logger.Error(err)

		return nil, err
	}
	return ud, err
}
func convertUserDao(u *User) *user.DeveloperDO {
	return &user.DeveloperDO{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
		Status:   u.Status,
	}
}

func convertUser(u *user.DeveloperDO) *User {
	return &User{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
		Status:   u.Status,
	}
}

func SaveCode(email string, code int) (string, error) {

	codeHash := GetEmailHashStr(email)

	err := db.RedisDB.HSet(codeHash, "code", code).Err()
	if err != nil {
		logger.Error(err)
		return "", err
	}
	// set A Five-Minute Expiration Time
	err = db.RedisDB.Expire(codeHash, 5*60*time.Second).Err()
	if err != nil {
		logger.Error(err)
		return "", err
	}

	return codeHash, nil
}
func GetEmailHashStr(email string) string {
	hash := sha1.Sum([]byte(email))
	hashString := fmt.Sprintf("%x", hash)
	return hashString
}
