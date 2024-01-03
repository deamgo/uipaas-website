package application

type Application struct {
	ID          string `json:"id" gorm:"type:varchar(25);primary_key;"`
	WorkspaceID string `json:"workspace_id" gorm:"type:char(6);not null;"`
	Name        string `json:"name" gorm:"type:varchar(50);not null;"`
	Description string `json:"description" gorm:"type:varchar(255);"`
	Icon        string `json:"icon" gorm:"type:varchar(255);"`
	Status      uint8  `json:"status" gorm:"type:tinyint;not null;"`

	CreatedBy uint64 `json:"created_by" gorm:"type:bigint;default:0;"`
	DeletedBy uint64 `json:"deleted_by" gorm:"type:bigint;default:0;"`
}
