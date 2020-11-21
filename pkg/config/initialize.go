package config

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

var (
	Db    *gorm.DB
	Redis *redis.Client
)

func Initialize() {}
