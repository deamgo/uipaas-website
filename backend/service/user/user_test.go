package user

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/gorm"
	"testing"

	dao "github.com/deamgo/uipass-waitlist-page/backend/dao/user"
	mock_test "github.com/deamgo/uipass-waitlist-page/backend/mock"
	"github.com/stretchr/testify/assert"
)

func setupUserServiceTest(t *testing.T) (UserService, sqlmock.Sqlmock) {
	mockDB, mock, err := mock_test.GetNewDbMock()
	assert.NoError(t, err)

	userDao := dao.NewAUserDao(mockDB)
	params := UserServiceParams{Dao: userDao}
	userservice := NewUserService(params)

	return userservice, mock
}

func TestUserService_UserGet(t *testing.T) {
	tests := []struct {
		name           string
		userID         string
		expectedUser   *User
		expectedError  error
		expectedErrMsg string
	}{
		{
			name:   "User exists",
			userID: "1",
			expectedUser: &User{
				UserID:   "1",
				UserName: "tomoki",
				Email:    "example@gmail.com",
				Password: "123",
			},
			expectedError: nil,
		},
		{
			name:           "User does not exist",
			userID:         "2",
			expectedUser:   nil,
			expectedError:  dao.UserNotExistError,
			expectedErrMsg: "User does not exist",
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userservice, mock := setupUserServiceTest(t)

			rows := mock.NewRows([]string{"userid", "username", "password", "email"}).
				AddRow("1", "tomoki", "123", "example@gmail.com")

			if tt.expectedError == nil {
				mock.ExpectQuery("(?i)SELECT\\s+\\*\\s+FROM\\s+`user`\\s+WHERE\\s+id\\s+=\\s+\\?\\s+ORDER\\s+BY\\s+`user`.`userid`\\s+LIMIT\\s+1").
					WithArgs(tt.userID).
					WillReturnRows(rows)
				user := User{UserID: tt.userID}
				userGet, err := userservice.UserGet(context.Background(), &user)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedUser, userGet)
			} else {
				mock.ExpectQuery("(?i)SELECT\\s+\\*\\s+FROM\\s+`user`\\s+WHERE\\s+id\\s+=\\s+\\?\\s+ORDER\\s+BY\\s+`user`.`userid`\\s+LIMIT\\s+1").
					WithArgs(tt.userID).
					WillReturnError(gorm.ErrRecordNotFound)
				user := User{UserID: tt.userID}
				userGet, err := userservice.UserGet(context.Background(), &user)
				assert.EqualError(t, err, tt.expectedError.Error())
				assert.Equal(t, tt.expectedUser, userGet)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
