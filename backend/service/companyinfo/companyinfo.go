package companyinfo

import (
	"context"
	dao "github.com/deamgo/uipaas-home/backend/dao/companyinfo"
	"github.com/deamgo/uipaas-home/backend/pkg/log"

	"go.uber.org/zap"
)

type CompanyInfoService interface {
	CompanyInfoGet(ctx context.Context, pageSize int, pageNum int) ([]*CompanyInfo, int64, error)
}

type CompanyInfoServiceParams struct {
	Dao dao.CompanyInfoDao
}

type companyInfoService struct {
	dao dao.CompanyInfoDao
}

func NewCompanyInfoService(params CompanyInfoServiceParams) CompanyInfoService {
	return &companyInfoService{
		dao: params.Dao,
	}
}

func (u companyInfoService) CompanyInfoGet(ctx context.Context, pageSize int, pageNum int) ([]*CompanyInfo, int64, error) {
	list, total, err := u.dao.CompanyInfoGet(ctx, pageSize, pageNum)
	if err != nil {
		log.Errorw("get companyinfolist failed",
			zap.Error(err),
			zap.Any("companyInfoList", list),
		)
		return nil, 0, err
	}

	return convertCompanyInfoList(list), total, err
}

func convertCompanyInfoList(companyInfoDao []*dao.CompanyInfoDO) []*CompanyInfo {
	CompanyInfoList := make([]*CompanyInfo, 0)
	for _, dao := range companyInfoDao {
		CompanyInfoList = append(CompanyInfoList, &CompanyInfo{
			ID:                     dao.ID,
			CompanyName:            dao.CompanyName,
			CompanySize:            dao.CompanySize,
			Name:                   dao.Name,
			Phone:                  dao.Phone,
			RequirementDescription: dao.RequirementDescription,
			Date:                   dao.Date,
		})
	}
	return CompanyInfoList

}
