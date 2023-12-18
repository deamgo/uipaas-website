package workspace

import (
	"github.com/DATA-DOG/go-sqlmock"
	dao "github.com/deamgo/workbench/dao/workspace"
	mockTest "github.com/deamgo/workbench/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func setupWorkspaceServiceTest(t *testing.T) (WorkspaceService, sqlmock.Sqlmock) {
	mockDB, mock, err := mockTest.GetNewDbMock()
	assert.NoError(t, err)

	workspaceDao := dao.NewWorkspaceDao(mockDB)
	params := WorkspaceServiceParams{Dao: workspaceDao}
	workspaceService := NewWorkspaceService(params)

	return workspaceService, mock
}

func TestWorkspaceService_WorkspaceCreate(t *testing.T) {
	tests := []struct {
		name          string
		expectedError error
		workspace     *Workspac
	}{
		{
			name:          "page add 1",
			expectedError: nil,
			workspace: &Workspace{
				Name:        "workspace-测试数据",
				Logo:        "/public/head.jpg",
				Lable:       "短描述",
				Description: "这是测试偷偷编写的workspace的长描述",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//workspaceservice, mock := setupWorkspaceServiceTest(t)

		})
	}
}
