package developer

import "time"

type DeveloperDO struct {
	ID       string    `json:"id" gorm:"primaryKey" `
	Username string    `json:"username" gorm:"unique" `
	Avatar   string    `json:"avatar"`
	Email    string    `json:"email" gorm:"unique;not null" `
	Password string    `json:"password" gorm:"not null"`
	CreateAt time.Time `json:"create_at" gorm:"autoCreateTime" `
	UpdateAt time.Time ` json:"update_at" gorm:"autoUpdateTime" `
	Status   int       `json:"status"`
}

// TableName
func (DeveloperDO) TableName() string {
	return "developer"
}
