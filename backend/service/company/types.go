package company

import (
	"time"
)

type Company struct {
	ID                     uint      `json:"id"`
	CompanyName            string    `json:"companyname"`
	CompanySize            string    `json:"companysize"`
	Name                   string    `json:"name"`
	Phone                  string    `json:"phone"`
	RequirementDescription string    `json:"requirementdescription"`
	Date                   time.Time `json:"date"`
}