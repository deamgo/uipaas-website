package user

type UserDO struct {
	UserID   string `gorm:"column:userid"`
	UserName string `gorm:"column:username"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
}

func (UserDO) TableName() string {
	return "user"
}
