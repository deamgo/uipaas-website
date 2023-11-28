package companyinfo

import (
	"context"
	"gorm.io/gorm"
)

type CompanyInfoDao interface {
	CompanyInfoGet(ctx context.Context, pageSize int, pageNum int) (mes []*CompanyInfoDO, total int64, err error)
}

type companyInfoDao struct {
	db *gorm.DB
}

func NewAUseFormDao(db *gorm.DB) CompanyInfoDao {
	return &companyInfoDao{
		db: db,
	}
}

func (dao *companyInfoDao) CompanyInfoGet(ctx context.Context, pageSize int, pageNum int) (mes []*CompanyInfoDO, total int64, err error) {

	m := dao.db.Model(&CompanyInfoDO{})
	m.Count(&total)
	if err = m.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&mes).Error; err != nil {

		return nil, 0, err
	}

	return mes, total, err
}
