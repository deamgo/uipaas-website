package users

import (
	"context"
	"github.com/deamgo/workbench/service/mail"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	dao "github.com/deamgo/workbench/dao/user"
	mockTest "github.com/deamgo/workbench/mock"
)

func setupUserServiceTest(t *testing.T) (UserService, sqlmock.Sqlmock) {
	mockDB, mock, err := mockTest.GetNewDbMock()
	assert.NoError(t, err)

	userDao := dao.NewAUserDao(mockDB)
	params := UserServiceParams{Dao: userDao, MailService: mail.NewMailService()}
	userservice := NewUserService(params)

	return userservice, mock
}

func TestUserService_UserStatusModifyByEmail(t *testing.T) {

	us, mock := setupUserServiceTest(t)

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `developer`").WithArgs(1, "test@example.com").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	u := &User{Email: "test@example.com", Status: 1}

	err := us.UserStatusModifyByEmail(context.Background(), u)

	assert.NoError(t, err)
}

func TestUserService_UserGetByEmail(t *testing.T) {

	us, mock := setupUserServiceTest(t)

	rows := sqlmock.NewRows([]string{"id", "email"}).
		AddRow("1", "test@example.com")
	mock.ExpectQuery("^SELECT (.+) FROM `developer`").WithArgs("test@example.com").WillReturnRows(rows)

	email := "test@example.com"
	user := &User{Email: email}
	nodes, err := us.UserGetByEmail(context.Background(), user)

	assert.NoError(t, err)
	assert.NotNil(t, nodes)
}

func TestUserService_UserGetByUserName(t *testing.T) {
	us, mock := setupUserServiceTest(t)

	rows := sqlmock.NewRows([]string{"id", "username"}).
		AddRow("1", "glancake")
	mock.ExpectQuery("^SELECT (.+) FROM `developer`").WithArgs("glancake").WillReturnRows(rows)

	username := "glancake"
	user := &User{Username: username}
	nodes, err := us.UserGetByUserName(context.Background(), user)

	assert.NoError(t, err)
	assert.NotNil(t, nodes)
}

func TestUserService_UserPasswordModifyByEmail(t *testing.T) {

	us, mock := setupUserServiceTest(t)

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `developer`").WithArgs("789798322", "test@example.com").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	u := &User{Email: "test@example.com", Password: "789798322"}

	err := us.UserPasswordModifyByEmail(context.Background(), u)

	assert.NoError(t, err)
}

func TestUserService_UserAdd(t *testing.T) {

}
