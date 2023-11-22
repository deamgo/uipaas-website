package db

import (
	"github.com/deamgo/uipass-waitlist-page/backend/pkg/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB = InitDB()

func InitDB() *gorm.DB {
	dsn := "root:@tcp(localhost:3306)/uipass-waitlist-page?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	return db
}
