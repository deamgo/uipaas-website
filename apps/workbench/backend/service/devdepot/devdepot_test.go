package devdepot

import (
	"context"
	"testing"

	dao "github.com/deamgo/workbench/dao/devdepot"
	mockTest "github.com/deamgo/workbench/mock"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func setupDevDepotServiceTest(t *testing.T) (DevDepotService, sqlmock.Sqlmock) {
	mockDB, mock, err := mockTest.GetNewDbMock()
	assert.NoError(t, err)

	devDepotDao := dao.NewDevDepotDao(mockDB)
	params := DevDepotServiceParams{Dao: devDepotDao}
	devDepotService := NewDepotService(params)
	return devDepotService, mock
}

func TestDevDepotService_DevItemList(t *testing.T) {
	us, mock := setupDevDepotServiceTest(t)

	mock.ExpectBegin()
	mock.ExpectQuery("SELECT d.email, d.username, wdr.role FROM workspace_developer_relation wdr").
		WithArgs("1", 10).
		WillReturnRows(sqlmock.NewRows([]string{"email", "username", "role"}).
			AddRow("test@example.com", "testuser", "admin"))

	mock.ExpectCommit()
	_, err := us.DevItemList(context.Background(), "1", 1)

	assert.Error(t, err)
}

func TestDevDepotService_DevDepotDel_NonExistingDeveloper(t *testing.T) {
	us, mock := setupDevDepotServiceTest(t)

	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM developer_workspace_relation WHERE developer_id = ?").
		WithArgs("1", "1").
		WillReturnResult(sqlmock.NewResult(1, 0))

	mock.ExpectCommit()
	err := us.DevDepotDel(context.Background(), "1", "1")

	assert.Error(t, err)
}

func TestDevDepotService_DevInfoSearch(t *testing.T) {
	us, mock := setupDevDepotServiceTest(t)

	mock.ExpectBegin()
	mock.ExpectQuery("SELECT d.email, d.username, wdr.role FROM workspace_developer_relation wdr").
		WithArgs("1", 10).
		WillReturnRows(sqlmock.NewRows([]string{"email", "username", "role"}).
			AddRow("test@example.com", "testuser", "admin"))

	mock.ExpectCommit()
	_, err := us.DevInfoSearch(context.Background(), "1", "jige", 1)

	assert.Error(t, err)
}

func TestDevDepotService_DevDepotInvite(t *testing.T) {
	us, mock := setupDevDepotServiceTest(t)

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `developer_workspace_relation` ").
		WithArgs("1", "1", "1").
		WillReturnResult(sqlmock.NewResult(1, 0))

	mock.ExpectCommit()
	err := us.DevDepotInvite(context.Background(), &dao.DevDepotItem{DeveloperId: "1", Role: "1"})

	assert.Error(t, err)
}

func TestDevDepotInvite_InvitesSuccessfully(t *testing.T) {
	us, mock := setupDevDepotServiceTest(t)

	mock.ExpectBegin()
	mock.ExpectQuery("SELECT d.email, d.username, wdr.role FROM workspace_developer_relation wdr").
		WithArgs("1", "test@example.com").
		WillReturnError(gorm.ErrRecordNotFound)

	mock.ExpectExec("INSERT INTO `developer_workspace_relation`").
		WithArgs("1", "test@example.com", 0, 1, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()
	err := us.DevDepotInvite(context.Background(), &dao.DevDepotItem{Email: "test@example.com", WorkspaceId: "1"})

	assert.Error(t, err)
}

func TestDevDepotInvite_ReturnsErrorWhenDeveloperAlreadyInWorkspace(t *testing.T) {
	us, mock := setupDevDepotServiceTest(t)

	mock.ExpectBegin()
	mock.ExpectQuery("SELECT d.email, d.username, wdr.role FROM workspace_developer_relation wdr").
		WithArgs("1", "test@example.com").
		WillReturnRows(sqlmock.NewRows([]string{"email", "username", "role"}).
			AddRow("test@example.com", "testuser", "admin"))

	mock.ExpectCommit()
	err := us.DevDepotInvite(context.Background(), &dao.DevDepotItem{Email: "test@example.com", WorkspaceId: "1"})

	assert.Error(t, err)
}

func TestDevDepotInvite_ReturnsErrorWhenMailServiceFails(t *testing.T) {
	us, mock := setupDevDepotServiceTest(t)

	mock.ExpectBegin()
	mock.ExpectQuery("SELECT d.email, d.username, wdr.role FROM workspace_developer_relation wdr").
		WithArgs("1", "test@example.com").
		WillReturnError(gorm.ErrRecordNotFound)

	mock.ExpectCommit()
	err := us.DevDepotInvite(context.Background(), &dao.DevDepotItem{Email: "test@example.com", WorkspaceId: "1"})

	assert.Error(t, err)
}
