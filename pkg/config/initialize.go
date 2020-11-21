package config

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"gocms/pkg/database"
)

var (
	Db    *gorm.DB
	Redis *redis.Client
)

func init() {
	InitEnv()
	Db = database.MysqlClient()
	Redis = database.RedisClient()
}

func Initialize() {}
