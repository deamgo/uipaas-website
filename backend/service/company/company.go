package company

import (
	"context"
	"github.com/deamgo/uipaas-home/backend/pkg/e"
	"github.com/pkg/errors"
	"regexp"

	dao "github.com/deamgo/uipaas-home/backend/dao/company"
	"github.com/deamgo/uipaas-home/backend/pkg/log"

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
	re := regexp.MustCompile(`^[\p{L}\p{N}]+$`) // (Chinese and English numerals but excluding special characters)
	match1 := re.MatchString(company.CompanySize)
	match2, _ := regexp.MatchString(`^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`, company.BusinessEmail) // Email (email format)
	match3 := re.MatchString(company.CompanyName)
	match4 := re.MatchString(company.Name)
	match5 := re.MatchString(company.RequirementDescription)

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
