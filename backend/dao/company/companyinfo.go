package company

import (
	"context"
	"gorm.io/gorm"
)

type CompanyDao interface {
	CompanyGet(ctx context.Context, pageSize int, pageNum int) (mes []*CompanyDO, total int64, err error)
}

type companyDao struct {
	db *gorm.DB
}

func NewACompanyFormDao(db *gorm.DB) CompanyDao {
	return &companyDao{
		db: db,
	}
}

func (dao *companyDao) CompanyGet(ctx context.Context, pageSize int, pageNum int) (mes []*CompanyDO, total int64, err error) {

	m := dao.db.Model(&CompanyDO{})
	m.Count(&total)
	if err = m.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&mes).Error; err != nil {

		return nil, 0, err
	}

	return mes, total, err
}
