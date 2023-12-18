package db

import (
	"context"
	"fmt"
	"github.com/deamgo/workbench/initialize"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

//var dbConfigPath *string

func InitDB() *gorm.DB {

	config := initialize.GetConfig()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.DBName,
		config.Database.Charset)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	DB = db

	return db
}

func NewDBClient(ctx context.Context) *gorm.DB {

	// Get the database connection.
	db := DB
	// Return a new database client with the context.
	return db.WithContext(ctx)
}
