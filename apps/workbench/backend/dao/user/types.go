package user

import "time"

type UserDO struct {
	UID            uint      `gorm:"primaryKey" json:"uid"`
	Username       string    `gorm:"unique" json:"username"`
	Avatar         string    `json:"avatar"`
	Email          string    `gorm:"unique;not null" json:"email"`
	Password       string    `gorm:"not null" json:"-"`
	InvitationCode string    `gorm:"unique;not null" json:"invitation_code"`
	CreateDate     time.Time `gorm:"autoCreateTime" json:"create_date"`
	UpdateDate     time.Time `gorm:"autoUpdateTime" json:"update_date"`
	//账号是否停用
	Deactivate bool   `gorm:"default:false" json:"deactivate"`
	Status     string `json:"status"`
}

// TableName 设置表名，如果没有这个方法，表名默认是结构体的复数形式（UserProfiles）
func (UserDO) TableName() string {
	return "user"
}
