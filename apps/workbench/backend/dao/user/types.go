package user

import "time"

type UserDO struct {
	UID            uint      `json:"uid" gorm:"primaryKey" `
	Username       string    `json:"username" gorm:"unique" `
	Avatar         string    `json:"avatar"`
	Email          string    `json:"email" gorm:"unique;not null" `
	Password       string    `json:"password" gorm:"not null"`
	InvitationCode string    `json:"invitation_code" gorm:"unique;not null" `
	CreateDate     time.Time `json:"create_date" gorm:"autoCreateTime" `
	UpdateDate     time.Time ` json:"update_date" gorm:"autoUpdateTime" `
	//whether The Account Is Deactivated
	Deactivate bool   `json:"deactivate" gorm:"default:false" `
	Status     string `json:"status"`
}

// TableName
func (UserDO) TableName() string {
	return "user"
}
