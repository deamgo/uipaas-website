package company

import (
	"context"
	"regexp"
	"testing"
	"time"

	dao "github.com/deamgo/uipaas-home/backend/dao/company"
	mock_test "github.com/deamgo/uipaas-home/backend/mock"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func setupCompanyServiceTest(t *testing.T) (CompanyService, sqlmock.Sqlmock) {
	mockDB, mock, err := mock_test.GetNewDbMock()
	assert.NoError(t, err)

	listDao := dao.NewACompanyFormDao(mockDB)
	params := CompanyServiceParams{Dao: listDao}
	service := NewcompanyService(params)

	return service, mock
}

func TestCompanyService_companyGet(t *testing.T) {
	tests := []struct {
		name          string
		pageNum       int
		pageSize      int
		expectedTotal int64
		expectedList  []*Company
		expectedError error
	}{
		{
			name:          "page list1",
			pageNum:       6,
			pageSize:      1,
			expectedTotal: 6,
			expectedList: []*Company{
				{ID: 6,
					CompanyName:            "公司名字1",
					CompanySize:            "12",
					Name:                   "张三",
					BusinessEmail:          "1231341413@qq.com",
					RequirementDescription: "描述1",
				},
			},
			expectedError: nil,
		},
		{
			name:          "page list2",
			pageNum:       0,
			pageSize:      -1,
			expectedTotal: 0,
			expectedList:  nil,
			expectedError: errors.New("record not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service, mock := setupCompanyServiceTest(t)

			rows := mock.NewRows([]string{"id", "company_name", "company_size", "name", "business_email", "requirement_description"}).
				AddRow("6", "公司名字1", "12", "张三", "1231341413@qq.com", "描述1")

			if tt.expectedError == nil {
				mock.ExpectQuery("SELECT").WithArgs().WillReturnRows(sqlmock.NewRows([]string{"total"}).AddRow(6))
				mock.ExpectQuery("(?i)SELECT\\s+\\*\\s+FROM\\s+`company`\\s+LIMIT\\s+1\\s+OFFSET\\s+5").
					WillReturnRows(rows)
				var pageNum = tt.pageNum
				var pageSize = tt.pageSize

				list, total, err := service.CompanyGet(context.Background(), pageSize, pageNum)
				assert.NoError(t, err)
				for i, index := range list {
					assert.Equal(t, tt.expectedList[i], index)
				}
				assert.Equal(t, tt.expectedTotal, total)
			} else {

				mock.ExpectQuery("SELECT").WithArgs().WillReturnRows(sqlmock.NewRows([]string{"total"}).AddRow(0))

				mock.ExpectQuery("(?i)SELECT\\s+\\*\\s+FROM\\s+`company`\\s+OFFSET\\s+1").
					WillReturnError(gorm.ErrRecordNotFound)

				var pageNum = tt.pageNum
				var pageSize = tt.pageSize
				list, total, err := service.CompanyGet(context.Background(), pageSize, pageNum)
				assert.EqualError(t, err, tt.expectedError.Error())
				assert.Equal(t, tt.expectedList, list)
				assert.Equal(t, tt.expectedTotal, total)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}

}

func TestCompanyInfoService_CompanyInfoAdd(t *testing.T) {

	tests := []struct {
		name          string
		companyInfo   *Company
		expectedError error
	}{
		{
			name: "page add1",
			companyInfo: &Company{
				ID:                     0,
				CompanyName:            "这是公司名字",
				CompanySize:            "12",
				Name:                   "王五",
				BusinessEmail:          "1231341413@qq.com",
				RequirementDescription: "描述1",
				Date:                   time.Now(),
			},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service, mock := setupCompanyServiceTest(t)

			info := tt.companyInfo
			if tt.expectedError == nil {
				mock.ExpectQuery("SELECT").WithArgs().WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))

				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta("INSERT INTO")).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
				err := service.CompanyAdd(context.Background(), info)
				assert.Equal(t, tt.expectedError, err)
			}
		})
	}
}
