package developer

import (
	"context"
	"crypto/sha1"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"time"

	"github.com/deamgo/workbench/dao/developer"
	"github.com/deamgo/workbench/db"
	"github.com/deamgo/workbench/pkg/logger"
	"github.com/deamgo/workbench/service/mail"
)

type UserService interface {
	DeveloperAdd(ctx context.Context, user *Developer) (string, error)
	DeveloperGetByID(ctx context.Context, id string) (*Developer, error)
	DeveloperGetByEmail(ctx context.Context, u *Developer) (*developer.DeveloperDO, error)
	ForgotVerifySend(ctx context.Context, u *Developer) (string, error)
	DeveloperGetByUserName(ctx context.Context, u *Developer) (*Developer, error)
	DeveloperNameModifyByID(ctx context.Context, u *Developer) error
	DeveloperStatusModifyByEmail(ctx context.Context, u *Developer) error
	DeveloperPasswordModifyByEmail(ctx context.Context, u *Developer) error
}

type DeveloperServiceParams struct {
	Dao         developer.DeveloperDao
	MailService mail.MailService
}

type developerService struct {
	dao  developer.DeveloperDao
	mail mail.MailService
}

func NewDeveloperService(params DeveloperServiceParams) UserService {
	return &developerService{
		dao:  params.Dao,
		mail: params.MailService,
	}
}

func (us developerService) DeveloperAdd(ctx context.Context, u *Developer) (string, error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		logger.Error(err)
	}
	generateID := node.Generate().String()

	ud := convertDeveloperDO(u)
	//send email and get  code
	code := us.mail.SendMail(ctx, u.Email)
	codeHash, err := SaveCode(u.Email, code)
	if err != nil {
		logger.Error(err)
		return "", err
	}
	fmt.Println(code)
	ud.ID = generateID
	err = us.dao.DeveloperAdd(ctx, ud)
	if err != nil {
		logger.Error(err)

		return "", err
	}
	return codeHash, nil
}

func (us developerService) DeveloperStatusModifyByEmail(ctx context.Context, u *Developer) error {
	ud := convertDeveloperDO(u)
	ud.Status = 1
	err := us.dao.DeveloperStatusModifyByEmail(ctx, ud)
	if err != nil {
		logger.Error(err)

		return err
	}
	return nil
}
func (us developerService) DeveloperNameModifyByID(ctx context.Context, u *Developer) error {
	ud := convertDeveloperDO(u)
	err := us.dao.DeveloperNameModifyByID(ctx, ud)
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (us developerService) ForgotVerifySend(ctx context.Context, u *Developer) (string, error) {
	code := mail.MailService.SendMail(mail.NewMailService(), ctx, u.Email)
	codeHash, err := SaveCode(u.Email, code)
	if err != nil {
		logger.Error(err)
		return "", err
	}
	return codeHash, nil
}

func (us developerService) DeveloperPasswordModifyByEmail(ctx context.Context, u *Developer) error {
	ud := convertDeveloperDO(u)
	err := us.dao.DeveloperPasswordModifyByEmail(ctx, ud)
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (us developerService) DeveloperGetByUserName(ctx context.Context, u *Developer) (*Developer, error) {
	var err error
	ud := convertDeveloperDO(u)
	ud, err = us.dao.DeveloperGetByUserName(ctx, ud)
	if err != nil {
		logger.Error(err)

		return nil, err
	}
	return convertDeveloper(ud), err
}
func (us developerService) DeveloperGetByID(ctx context.Context, id string) (*Developer, error) {
	var err error
	u := &Developer{
		ID: id,
	}
	ud := convertDeveloperDO(u)
	ud, err = us.dao.DeveloperGetByUserName(ctx, ud)
	if err != nil {
		logger.Error(err)

		return nil, err
	}
	return convertDeveloper(ud), err
}

func (us developerService) DeveloperGetByEmail(ctx context.Context, u *Developer) (*developer.DeveloperDO, error) {
	var err error
	ud := convertDeveloperDO(u)
	ud, err = us.dao.DeveloperGetByEmail(ctx, ud)
	if err != nil {
		logger.Error(err)

		return nil, err
	}
	return ud, err
}
func convertDeveloperDO(u *Developer) *developer.DeveloperDO {
	return &developer.DeveloperDO{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
		Status:   u.Status,
	}
}

func convertDeveloper(u *developer.DeveloperDO) *Developer {
	return &Developer{
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
