package company

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

var CompanyExistError = errors.New("company name is exist")

type CompanyDao interface {
	CompanyGet(ctx context.Context, pageSize int, pageNum int) (mes []*CompanyDO, total int64, err error)
	CompanyAdd(ctx context.Context, company *CompanyDO) error
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

func (dao *companyDao) CompanyAdd(ctx context.Context, company *CompanyDO) error {
	var count int64
	company.Date = time.Now()
	dao.db.Model(&CompanyDO{}).Where("company_name=?", company.CompanyName).Find(&company).Count(&count)
	if count != 0 {
		return CompanyExistError
	}
	if err := dao.db.Create(&company).Error; err != nil {
		return err
	}
	return nil
}
