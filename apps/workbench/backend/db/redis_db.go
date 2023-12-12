package db

import (
	"fmt"
	"github.com/go-redis/redis"
	"gopkg.in/yaml.v3"
	"log"
)

var RedisDB *redis.Client

type RedisConfig struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

func InitRedis() {
	var (
		dbconfigFile []byte
	)

	var rc RedisConfig
	if err := yaml.Unmarshal(dbconfigFile, &rc); err != nil {
		log.Fatalf("Parsing config file: %v", err)
	}

	RedisDB = redis.NewClient(&redis.Options{
		Addr:     rc.Addr,
		Password: rc.Password,
		DB:       rc.DB,
	})
	pong, err := RedisDB.Ping().Result()
	if err != nil {
		fmt.Println(pong, err)
	}
	fmt.Println("redis 连接成功")

}
