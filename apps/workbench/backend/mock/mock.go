package mock

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/DATA-DOG/go-sqlmock"
)

func GetNewDbMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, mock, err
	}

	gormDB, err := gorm.Open(mysql.Dialector{Config: &mysql.Config{DriverName: "mysql", Conn: db, SkipInitializeWithVersion: true}}, &gorm.Config{})

	if err != nil {
		return gormDB, mock, err
	}

	return gormDB, mock, err
}
