package company

import (
	"context"
	dao "github.com/deamgo/uipaas-home/backend/dao/company"
	"github.com/deamgo/uipaas-home/backend/pkg/log"

	"go.uber.org/zap"
)

type CompanyService interface {
	CompanyGet(ctx context.Context, pageSize int, pageNum int) ([]*Company, int64, error)
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

func convertCompanyList(companyDao []*dao.CompanyDO) []*Company {
	companyList := make([]*Company, 0)
	for _, dao := range companyDao {
		companyList = append(companyList, &Company{
			ID:                     dao.ID,
			CompanyName:            dao.CompanyName,
			CompanySize:            dao.CompanySize,
			Name:                   dao.Name,
			Phone:                  dao.Phone,
			RequirementDescription: dao.RequirementDescription,
			Date:                   dao.Date,
		})
	}
	return companyList

}
