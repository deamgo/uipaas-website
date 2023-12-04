package user

import (
	"context"
	"reflect"
	"testing"

	dao "github.com/deamgo/uipaas-home/backend/dao/user"
	mock_test "github.com/deamgo/uipaas-home/backend/mock"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
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

func TestUserService_UserLogin(t *testing.T) {
	tests := []struct {
		name           string
		UserName       string
		Password       string
		expectedError  error
		expectedErrMsg string
		expectedToken  string
	}{
		{
			name:          "User login success",
			UserName:      "zhangsan",
			Password:      "1234",
			expectedError: nil,
			expectedToken: "",
		},
		{
			name:           "User login failed",
			UserName:       "lisi",
			Password:       "1234",
			expectedError:  dao.UserLoginError,
			expectedErrMsg: "incorrect username or password",
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userservice, mock := setupUserServiceTest(t)

			rows := mock.NewRows([]string{"userid", "username", "password"}).
				AddRow("1", "zhangsan", "1234")

			if tt.expectedError == nil {
				mock.ExpectQuery("(?i)SELECT\\s+\\*\\s+FROM\\s+`user`\\s+WHERE\\s+username\\s+=\\s+\\?\\s+AND\\s+password\\s+=\\s+\\?").
					WithArgs(tt.UserName, tt.Password).
					WillReturnRows(rows)
				user := User{UserName: tt.UserName,
					Password: tt.Password}
				token, err := userservice.UserLogin(context.Background(), &user)
				assert.NoError(t, err)
				assert.Equal(t, reflect.TypeOf(token).Kind(), reflect.TypeOf(tt.expectedToken).Kind())
				assert.Equal(t, tt.expectedError, err)
			} else {
				mock.ExpectQuery("(?i)SELECT\\s+\\*\\s+FROM\\s+`user`\\s+WHERE\\s+username\\s+=\\s+\\?\\s+AND\\s+password\\s+=\\s+\\?").
					WithArgs(tt.UserName, tt.Password).
					WillReturnError(gorm.ErrRecordNotFound)
				user := User{UserName: tt.UserName,
					Password: tt.Password}
				token, err := userservice.UserLogin(context.Background(), &user)
				assert.EqualError(t, err, tt.expectedError.Error())
				assert.Equal(t, reflect.TypeOf(token).Kind(), reflect.TypeOf(tt.expectedToken).Kind())
				assert.Equal(t, tt.expectedError, err)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
