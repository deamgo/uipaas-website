package company

import (
	"time"
)

type Company struct {
	ID                     uint      `json:"id" `
	CompanyName            string    `json:"companyname" validate:"min=1,max=50"`
	CompanySize            string    `json:"companysize" validate:"min=1,max=20"`
	Name                   string    `json:"name" validate:"min=1,max=50"`
	BusinessEmail          string    `json:"businessemail" validate:"min=1,max=100"`
	RequirementDescription string    `json:"requirementdescription" validate:"min=1,max=200"`
	Date                   time.Time `json:"date"`
}
