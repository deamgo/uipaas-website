package user

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
	dao "workbench/dao/user"
	mockTest "workbench/mock"
)

func setupUserServiceTest(t *testing.T) (UserService, sqlmock.Sqlmock) {
	mockDB, mock, err := mockTest.GetNewDbMock()
	assert.NoError(t, err)

	userDao := dao.NewAUserDao(mockDB)
	params := UserServiceParams{Dao: userDao}
	userservice := NewUserService(params)

	return userservice, mock
}
func TestUserService_UserAdd(t *testing.T) {

	tests := []struct {
		name          string
		user          *User
		expectedError error
	}{
		{
			name: "page add1",
			user: &User{
				Username:       "glancake",
				Email:          "dsfd242@qq.com",
				Password:       "2j532lj5k32",
				InvitationCode: "biECiZ",
			},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service, mock := setupUserServiceTest(t)

			info := tt.user
			if tt.expectedError == nil {
				mock.ExpectQuery("SELECT").WithArgs().WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))

				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta("INSERT INTO")).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
				_, err := service.UserAdd(context.Background(), info)
				assert.Equal(t, tt.expectedError, err)
			}
		})
	}

}
