package developer

import (
	"context"
	"testing"

	dao "github.com/deamgo/workbench/dao/developer"
	mockTest "github.com/deamgo/workbench/mock"
	"github.com/deamgo/workbench/service/mail"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func setupDeveloperServiceTest(t *testing.T) (UserService, sqlmock.Sqlmock) {
	mockDB, mock, err := mockTest.GetNewDbMock()
	assert.NoError(t, err)

	userDao := dao.NewADeveloperDao(mockDB)
	params := DeveloperServiceParams{Dao: userDao, MailService: mail.NewMailService()}
	userservice := NewDeveloperService(params)

	return userservice, mock
}

func TestDeveloperService_DeveloperStatusModifyByEmail(t *testing.T) {

	us, mock := setupDeveloperServiceTest(t)

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `developer`").WithArgs(1, "test@example.com").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	u := &Developer{Email: "test@example.com", Status: 1}

	err := us.DeveloperStatusModifyByEmail(context.Background(), u)

	assert.NoError(t, err)
}

func TestDeveloperService_DeveloperGetByEmail(t *testing.T) {

	us, mock := setupDeveloperServiceTest(t)

	rows := sqlmock.NewRows([]string{"id", "email"}).
		AddRow("1", "test@example.com")
	mock.ExpectQuery("^SELECT (.+) FROM `developer`").WithArgs("test@example.com").WillReturnRows(rows)

	email := "test@example.com"
	user := &Developer{Email: email}
	nodes, err := us.DeveloperGetByEmail(context.Background(), user)

	assert.NoError(t, err)
	assert.NotNil(t, nodes)
}

func TestDeveloperService_DeveloperGetByUserNamet(t *testing.T) {
	us, mock := setupDeveloperServiceTest(t)

	rows := sqlmock.NewRows([]string{"id", "username"}).
		AddRow("1", "glancake")
	mock.ExpectQuery("^SELECT (.+) FROM `developer`").WithArgs("glancake").WillReturnRows(rows)

	username := "glancake"
	user := &Developer{Username: username}
	nodes, err := us.DeveloperGetByUserName(context.Background(), user)

	assert.NoError(t, err)
	assert.NotNil(t, nodes)
}

func TestDeveloperService_DeveloperPasswordModifyByEmail(t *testing.T) {

	us, mock := setupDeveloperServiceTest(t)

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `developer`").WithArgs("789798322", "test@example.com").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	u := &Developer{Email: "test@example.com", Password: "789798322"}

	err := us.DeveloperPasswordModifyByEmail(context.Background(), u)

	assert.NoError(t, err)
}

func TestDeveloperService_DeveloperGetByUserName(t *testing.T) {

	us, mock := setupDeveloperServiceTest(t)

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `developer`").WithArgs("crfillee", "1735589460680118272").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	u := &Developer{ID: "1735589460680118272", Username: "crfillee"}

	err := us.DeveloperNameModifyByID(context.Background(), u)

	assert.NoError(t, err)
}

func TestDeveloperService_DeveloperGetByID(t *testing.T) {

	us, mock := setupDeveloperServiceTest(t)

	rows := sqlmock.NewRows([]string{"id", "email", "username", "password", "status"}).
		AddRow("1", "test@example.com", "testuser", "password", 1)
	mock.ExpectQuery("^SELECT (.+) FROM `developer`").WithArgs("1").WillReturnRows(rows)

	id := "1"
	user, err := us.DeveloperGetByID(context.Background(), id)

	assert.NoError(t, err)
	assert.NotNil(t, user)

}

func TestDeveloperService_DeveloperEmailModifyByEmail(t *testing.T) {
	us, mock := setupDeveloperServiceTest(t)

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `developer`").WithArgs("newtest@example.com", "test@example.com").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	u := &Developer{Email: "newtest@example.com"}

	err := us.DeveloperEmailModifyByEmail(context.Background(), "test@example.com", u)

	assert.NoError(t, err)
}

func TestDeveloperService_DeveloperEmailModifyByEmail_EmailNotFound(t *testing.T) {
	us, mock := setupDeveloperServiceTest(t)

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `developer`").WithArgs("newtest@example.com", "test@example.com").
		WillReturnError(gorm.ErrRecordNotFound)
	mock.ExpectRollback()

	u := &Developer{Email: "newtest@example.com"}

	err := us.DeveloperEmailModifyByEmail(context.Background(), "test@example.com", u)

	assert.Error(t, err)
}

func TestDeveloperService_DeveloperGetByID_NotFound(t *testing.T) {
	us, mock := setupDeveloperServiceTest(t)

	mock.ExpectQuery("^SELECT (.+) FROM `developer`").WithArgs("1").
		WillReturnError(gorm.ErrRecordNotFound)

	id := "1"
	_, err := us.DeveloperGetByID(context.Background(), id)

	assert.Error(t, err)
}
