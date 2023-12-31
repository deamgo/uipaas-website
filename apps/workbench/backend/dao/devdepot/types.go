package devdepot

import "time"

type DevDepotItem struct {
	WorkspaceId string `json:"workspace_id" gorm:"type:char(6);size 6;primary_key;comment:'workspace ID'"`
	DeveloperId string `json:"developer_id" gorm:"type:varchar;size 20;not null;comment:'Developer ID'"`
	Role        string `json:"role" gorm:"type:varchar(10);not null;comment:'role 0-owner 1-Admin 2-Developer 3-Reviewer'"`
	Status      uint8  `json:"status" gorm:"type:tinyint;not null;comment:'status 0-Pending 1-Accept'"`
	Username    string `json:"username" gorm:"type:varchar(50);not null;comment:'username'"`
	Email       string `json:"email" gorm:"type:varchar(50);not null;comment:'email'"`
}

type DeveloperWorkspaceRelationDO struct {
	WorkspaceId string `gorm:"type:char(6);size 6;primary_key;comment:'workspace ID'"`
	Email       string `gorm:"type:varchar(50);not null;comment:'email'"`
	DeveloperId string `gorm:"type:varchar;size 20;not null;comment:'Developer ID'"`
	Role        uint8  `gorm:"type:tinyint;not null;comment:'role  1-Admin 2-Developer 3-Reviewer'"`
	Status      uint8  `json:"status" gorm:"type:tinyint;not null;comment:'status 0-Pending 1-Accept'"`

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
