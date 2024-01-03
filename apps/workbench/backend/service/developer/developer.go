package developer

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/deamgo/workbench/dao/developer"
	"github.com/deamgo/workbench/db"
	"github.com/deamgo/workbench/pkg/logger"
	"github.com/deamgo/workbench/service/mail"

	"github.com/bwmarrin/snowflake"
	"gorm.io/gorm"
)

// UserService DeveloperService is an interface that defines the methods that our service should implement.
type UserService interface {
	// DeveloperAdd adds a new developer to the system.
	DeveloperAdd(ctx context.Context, user *Developer) (string, error)
	// DeveloperGetByID retrieves a developer by their ID.
	DeveloperGetByID(ctx context.Context, id string) (*Developer, error)
	// DeveloperGetByEmail retrieves a developer by their email.
	DeveloperGetByEmail(ctx context.Context, u *Developer) (*developer.DeveloperDO, error)
	DeveloperGetByEmailAndStatus(ctx context.Context, u *Developer) (*developer.DeveloperDO, error)
	// DeveloperGetByEmailAndPwd retrieves a developer by their email and password.
	DeveloperGetByEmailAndPwd(ctx context.Context, u *Developer) (*developer.DeveloperDO, error)
	// ForgotVerifySend sends a verification code to a developer who forgot their password.
	ForgotVerifySend(ctx context.Context, u *Developer) (string, error)
	// DeveloperGetByUserName retrieves a developer by their username.
	DeveloperGetByUserName(ctx context.Context, u *Developer) (*Developer, error)
	// DeveloperNameModifyByID modifies a developer's name by their ID.
	DeveloperNameModifyByID(ctx context.Context, u *Developer) error
	// SendModifyEmailVerify sends a verification code to a developer who wants to modify their email.
	SendModifyEmailVerify(ctx context.Context, u *Developer) (string, error)
	// DeveloperEmailModifyByEmail modifies a developer's email by their old email.
	DeveloperEmailModifyByEmail(ctx context.Context, oldEmail string, u *Developer) error
	// DeveloperStatusModifyByEmail modifies a developer's status by their email.
	DeveloperStatusModifyByEmail(ctx context.Context, u *Developer) error
	// DeveloperPasswordModifyByEmail modifies a developer's password by their email.
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
	code := us.mail.SendVerificationCodeMail(ctx, u.Email)
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
	code := mail.MailService.SendVerificationCodeMail(mail.NewMailService(), ctx, u.Email)
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

	ud, err = us.dao.DeveloperGetByID(ctx, ud)
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
func (us developerService) DeveloperGetByEmailAndStatus(ctx context.Context, u *Developer) (*developer.DeveloperDO, error) {
	var err error
	ud := convertDeveloperDO(u)
	ud, err = us.dao.DeveloperGetByEmailAndStatus(ctx, ud)
	if err != nil {
		logger.Error(err)

		return nil, err
	}
	return ud, err
}
func (us developerService) DeveloperGetByEmailAndPwd(ctx context.Context, u *Developer) (*developer.DeveloperDO, error) {
	dlp := convertDeveloperDO(u)
	dlp, err := us.dao.DeveloperGetByEmailAndPwd(ctx, dlp)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return dlp, nil
}
func (us developerService) DeveloperEmailModifyByEmail(ctx context.Context, oldEmail string, u *Developer) error {
	dlp := convertDeveloperDO(u)
	err := us.dao.DeveloperEmailModifyByEmail(ctx, oldEmail, dlp)
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (us developerService) SendModifyEmailVerify(ctx context.Context, u *Developer) (string, error) {
	dlp, err := us.dao.DeveloperGetByEmail(ctx, convertDeveloperDO(u))
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Error(err)
			return "", err
		}
	}
	if dlp != nil {
		return "", errors.New("email already exists")
	}
	code := mail.MailService.SendVerificationCodeMail(mail.NewMailService(), ctx, u.Email)
	var codeHash string
	codeHash, err = SaveCode(u.Email, code)
	if err != nil {
		logger.Error(err)
		return "", err
	}

	return codeHash, nil
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
