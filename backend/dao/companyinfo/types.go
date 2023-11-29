package companyinfo

import (
	"time"
)

type CompanyInfoDO struct {
	ID                     uint      `gorm:"column:id"`
	CompanyName            string    `gorm:"column:company_name"`
	CompanySize            string    `gorm:"column:company_size"`
	Name                   string    `gorm:"column:name"`
	Phone                  string    `gorm:"column:phone"`
	RequirementDescription string    `gorm:"column:requirement_description"`
	Date                   time.Time `gorm:"column:date"`
}

func (CompanyInfoDO) TableName() string {
	return "company_info"
}
