package db

import (
	"flag"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB = InitDB()

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	Charset  string `yaml:"charset"`
}

type Config struct {
	Database DBConfig `yaml:"database"`
}

var dbConfigPath = flag.String("dbConfig", "", "db config path")

func InitDB() *gorm.DB {
	path, err := os.Getwd()
	if err != nil {
		log.Fatalf("Cannot open config file: %v", err)
	}
	var (
		configFile []byte
	)
	if *dbConfigPath == "" {
		configFile, err = os.ReadFile(path + "/db/config.yaml")
	} else {
		configFile, err = os.ReadFile(*dbConfigPath)
	}

	if err != nil {
		log.Fatalf("Cannot open config file: %v", err)
	}

	var config Config
	if err = yaml.Unmarshal(configFile, &config); err != nil {
		log.Fatalf("Parsing config file: %v", err)
	}

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

	return db
}
