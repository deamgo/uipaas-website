package company

import (
	"time"
)

type CompanyDO struct {
	ID                     uint      `gorm:"column:id"`
	CompanyName            string    `gorm:"column:company_name"`
	CompanySize            string    `gorm:"column:company_size"`
	Name                   string    `gorm:"column:name"`
	BusinessEmail          string    `gorm:"column:business_email"`
	RequirementDescription string    `gorm:"column:requirement_description"`
	Date                   time.Time `gorm:"column:date"`
}

func (CompanyDO) TableName() string {

	return "company"

}
