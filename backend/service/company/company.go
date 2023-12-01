package company

import (
	"context"
	"errors"
	dao "github.com/deamgo/uipaas-home/backend/dao/company"
	"github.com/deamgo/uipaas-home/backend/pkg/e"
	"github.com/deamgo/uipaas-home/backend/pkg/log"
	"regexp"

	"go.uber.org/zap"
)

type CompanyService interface {
	CompanyGet(ctx context.Context, pageSize int, pageNum int) ([]*Company, int64, error)
	CompanyAdd(ctx context.Context, company *Company) error
}

type CompanyServiceParams struct {
	Dao dao.CompanyDao
}

type companyService struct {
	dao dao.CompanyDao
}

func NewcompanyService(params CompanyServiceParams) CompanyService {
	return &companyService{
		dao: params.Dao,
	}
}

func (u companyService) CompanyGet(ctx context.Context, pageSize int, pageNum int) ([]*Company, int64, error) {
	list, total, err := u.dao.CompanyGet(ctx, pageSize, pageNum)
	if err != nil {
		log.Errorw("get companylist failed",
			zap.Error(err),
			zap.Any("companyList", list),
		)
		return nil, 0, err
	}

	return convertCompanyList(list), total, err
}

func (u companyService) CompanyAdd(ctx context.Context, company *Company) error {
	match1, _ := regexp.MatchString(`^\d{1,10}$`, company.CompanySize)                                      // Company size (number: 1-10 digits)
	match2, _ := regexp.MatchString(`^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`, company.BusinessEmail) // Email (email format)

	re := regexp.MustCompile(`^[\p{L}\p{N}]+$`) //Company Name (Chinese and English numerals but excluding special characters)
	match3 := re.MatchString(company.CompanyName)

	re1 := regexp.MustCompile(`^[\p{L}\p{N}]+$`) //Name (Chinese and English numerals but excluding special characters)
	match4 := re1.MatchString(company.Name)

	re2 := regexp.MustCompile(`^[\p{L}\p{N}]+$`) //RequirementDescription (Chinese and English numerals but excluding special characters)
	match5 := re2.MatchString(company.RequirementDescription)

	if !match1 || !match2 || !match3 || !match4 || !match5 {
		return errors.New(e.AddFormatError)
	}
	info := convertCompanyDao(company)
	err := u.dao.CompanyAdd(ctx, info)
	if err != nil {
		log.Errorw("add companyinfo failed",
			zap.Error(err),
			zap.Any("companyInfo", company),
		)
		return err
	}
	return nil
}

func convertCompanyDao(info *Company) *dao.CompanyDO {
	return &dao.CompanyDO{
		ID:                     info.ID,
		CompanyName:            info.CompanyName,
		CompanySize:            info.CompanySize,
		Name:                   info.Name,
		BusinessEmail:          info.BusinessEmail,
		RequirementDescription: info.RequirementDescription,
		Date:                   info.Date,
	}
}

func convertCompanyList(companyDao []*dao.CompanyDO) []*Company {
	companyList := make([]*Company, 0)
	for _, dao := range companyDao {
		companyList = append(companyList, &Company{
			ID:                     dao.ID,
			CompanyName:            dao.CompanyName,
			CompanySize:            dao.CompanySize,
			Name:                   dao.Name,
			BusinessEmail:          dao.BusinessEmail,
			RequirementDescription: dao.RequirementDescription,
			Date:                   dao.Date,
		})
	}
	return companyList

}
