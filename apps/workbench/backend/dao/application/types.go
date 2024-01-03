package application

import "time"

type ApplicationDO struct {
	ID          string `gorm:"type:varchar(25);primary_key;"`
	WorkspaceID string `gorm:"type:char(6);not null;"`
	Name        string `gorm:"type:varchar(50);not null;"`
	Description string `gorm:"type:varchar(255);"`
	Icon        string `gorm:"type:varchar(255);"`
	Status      uint8  `gorm:"type:tinyint;not null;"`

	CreatedBy uint64    `gorm:"type:bigint;default:0;not null;comment:'creator';"`
	CreatedAt time.Time `gorm:"comment:'creation time';"`
	UpdatedBy uint64    `gorm:"type:bigint;default:0;not null;comment:'The last person to update the data';"`
	UpdatedAt time.Time `gorm:"comment:'update time';"`
	DeletedBy uint64    `gorm:"type:bigint;default:0;comment:'Deleting people';"`
	DeletedAt time.Time `gorm:"default:null';comment:'Delete time';"`
	IsDeleted uint8     `gorm:"type:tinyint;default:0;not null;index;comment:'Logical deletion 0-not deleted 1-deleted'"`
}

func (ApplicationDO) TableName() string {
	return "application"
}
