package workspace

import (
	"context"
	"database/sql/driver"
	"testing"
	"time"

	dao "github.com/deamgo/workbench/dao/workspace"
	mockTest "github.com/deamgo/workbench/mock"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func setupWorkspaceServiceTest(t *testing.T) (WorkspaceService, sqlmock.Sqlmock) {
	mockDB, mock, err := mockTest.GetNewDbMock()
	assert.NoError(t, err)

	workspaceDao := dao.NewWorkspaceDao(mockDB)
	params := WorkspaceServiceParams{Dao: workspaceDao}
	workspaceService2 := NewWorkspaceService(params)

	return workspaceService2, mock
}

func TestWorkspaceService_WorkspaceGetListById(t *testing.T) {
	tests := []struct {
		name          string
		expectedError error
		developerId   uint64
	}{
		{
			name:          "TestWorkspaceService_WorkspaceGetListById 1",
			expectedError: nil,
			developerId:   1736939743536484352,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			workspaceservice, mock := setupWorkspaceServiceTest(t)
			//select w.* from workspace_developer_relation r left join workspaces w on w.id = r.workspace_id where developer_id = 1736939743536484352;
			rows := sqlmock.NewRows([]string{"id", "name", "label", "description", "logo"}).
				AddRow("d30340", "test021", "", "这是用于测试workspace的测试数据", "/public/Golang.png")

			mock.ExpectQuery("^select w.* from workspace_developer_relation r left join workspaces w on w.id = r.workspace_id where developer_id = (.+) and w.is_deleted = 0;").
				WithArgs(test.developerId).
				WillReturnRows(rows)

			newWorkspace, err := workspaceservice.WorkspaceGetListById(context.Background(), test.developerId)

			assert.NoError(t, err)
			assert.NotNil(t, newWorkspace)
		})
	}
}

func TestWorkspaceService_WorkspaceCreate(t *testing.T) {
	tests := []struct {
		name          string
		expectedError error
		workspace     *Workspace
	}{
		{
			name:          "data 1",
			expectedError: nil,
			workspace: &Workspace{
				Name:        "test1",
				Logo:        "/public/head.jpg",
				Label:       "短描述",
				Description: "这是测试偷偷编写的workspace的长描述",
				CreatedBy:   1,
				UpdateBy:    1,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			workspaceservice, mock := setupWorkspaceServiceTest(t)
			workspace := test.workspace
			//INSERT INTO `workspace` (`id`,`name`,`logo`,`label`,`description`,`created_by`,`created_at`,`updated_by`,`updated_at`,`deleted_by`,`is_deleted`) VALUES ('68c1bf','第三个2据121','/public/Golang.png','工作空间测试数据','这是用于测试workspace的测试数据',0,'2023-12-18 18:19:41.942',0,'2023-12-18 18:19:41.942',0,0)

			mock.ExpectBegin()

			mock.ExpectExec("INSERT INTO `workspaces`").
				WithArgs(hashTop(workspace.Name, 6), workspace.Name, workspace.Logo, workspace.Label, workspace.Description, 1, AnyTime{}, 1, AnyTime{}, 0, 0).
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectExec("INSERT INTO `workspace_developer_relation`").
				WithArgs(hashTop(workspace.Name, 6), 1, 1, 1, AnyTime{}, 1, AnyTime{}, 0, 0).
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()

			newWorkspace, err := workspaceservice.WorkspaceCreate(context.Background(), workspace)

			assert.NoError(t, err)
			assert.NotNil(t, newWorkspace)
		})
	}
}

func TestWorkspaceService_WorkspaceDel(t *testing.T) {
	us, mock := setupWorkspaceServiceTest(t)

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `workspaces`").WithArgs(1, "1").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("UPDATE `workspace_developer_relation`").WithArgs(1, "1").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	u := &Workspace{Id: "1"}

	err := us.WorkspaceDel(context.Background(), u, "1")

	assert.NoError(t, err)
}

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}
