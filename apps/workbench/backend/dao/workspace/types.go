package workspace

import (
	"time"
)

type WorkspaceDO struct {
	Id          string `gorm:"type:char(6);size 6;primary_key;comment:'workspace ID'"`
	Name        string `gorm:"type:char(20);size 20;not null;comment:'workspace Name'"`
	Logo        string `gorm:"type:varchar(255);size 255;not null;comment:'logo'"`
	Label       string `gorm:"type:varchar(255);size 255;not null;default:'';comment:'label'"`
	Description string `gorm:"type:varchar(1023);size 1023;not null;default:'';comment:'description'"`

	CreatedBy uint64    `gorm:"type:bigint;default:0;not null;comment:'creator';"`
	CreatedAt time.Time `gorm:"comment:'creation time';"`
	UpdatedBy uint64    `gorm:"type:bigint;default:0;not null;comment:'The last person to update the data';"`
	UpdatedAt time.Time `gorm:"comment:'update time';"`
	DeletedBy uint64    `gorm:"type:bigint;default:0;comment:'Deleting people';"`
	DeletedAt time.Time `gorm:"default:null';comment:'Delete time';"`
	IsDeleted uint8     `gorm:"type:tinyint;default:0;not null;index;comment:'Logical deletion 0-not deleted 1-deleted'"`
}

func (WorkspaceDO) TableName() string {
	return "workspaces"
}

type DeveloperWorkspaceRelationDO struct {
	WorkspaceId string `gorm:"type:char(6);size 6;primary_key;comment:'workspace ID'"`
	DeveloperId uint64 `gorm:"type:bigint;size 20;not null;comment:'Developer ID'"`
	Role        uint8  `gorm:"type:tinyint;not null;comment:'role  1-Admin 2-Developer 3-Reviewer'"`

	CreatedBy uint64    `gorm:"type:bigint;default:0;not null;comment:'creator';"`
	CreatedAt time.Time `gorm:"comment:'creation time';"`
	UpdatedBy uint64    `gorm:"type:bigint;default:0;not null;comment:'The last person to update the data';"`
	UpdatedAt time.Time `gorm:"comment:'update time';"`
	DeletedBy uint64    `gorm:"type:bigint;default:0;comment:'Deleting people';"`
	DeletedAt time.Time `gorm:"default:null';comment:'Delete time';"`
	IsDeleted uint8     `gorm:"type:tinyint;default:0;not null;index;comment:'Logical deletion 0-not deleted 1-deleted'"`
}

func (DeveloperWorkspaceRelationDO) TableName() string {
	return "workspace_developer_relation"
}
