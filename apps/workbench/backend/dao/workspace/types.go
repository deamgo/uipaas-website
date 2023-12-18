package workspace

import (
	"time"
)

type WorkspaceDO struct {
	Id          string `gorm:"type:char(6);size 6;primary_key;comment:'工作空间ID'"`
	Name        string `gorm:"type:char(20);size 20;not null;comment:'工作空间名称'"`
	Logo        string `gorm:"type:varchar(255);size 255;not null;comment:'图标地址'"`
	Lable       string `gorm:"type:varchar(255);size 255;not null;default:'';comment:'短描述'"`
	Description string `gorm:"type:varchar(1023);size 1023;not null;default:'';comment:'长描述'"`

	CreatedBy uint64    `gorm:"type:bigint;default:0;not null;comment:'创建人';"`
	CreatedAt time.Time `gorm:"comment:'创建时间';"`
	UpdatedBy uint64    `gorm:"type:bigint;default:0;not null;comment:'最后一次更新的开发者';"`
	UpdatedAt time.Time `gorm:"comment:'更新时间';"`
	DeletedBy uint64    `gorm:"type:bigint;default:0;comment:'删除人';"`
	DeletedAt time.Time `gorm:"default:null';comment:'删除时间';"`
	IsDeleted uint8     `gorm:"type:tinyint;default:0;not null;index;comment:'逻辑删除 0-未删除 1-已删除'"`
}

func (WorkspaceDO) TableName() string {
	return "workspace"
}
