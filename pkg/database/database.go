package database

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/cast"
	"gocms/pkg/config"
	"gocms/pkg/logger"
)

func init() {
	config.Db = MysqlClient()
	config.Redis = RedisClient()
}

func Initialize() {}

// 初始化 Redis 服务器
func RedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     config.GetEnv("REDIS_HOST") + ":" + config.GetEnv("REDIS_PORT"),
		Password: config.GetEnv("REDIS_PASSWORD"),
		DB:       cast.ToInt(config.GetEnv("REDIS_DB")),
	})

	_, err := client.Ping().Result()
	logger.PanicError(err, "Redis 连接", true)

	logger.Info("连接Redis 成功", "redis connect")
	return client
}

// mysql 服务器
func MysqlClient() *gorm.DB {
	username := config.GetEnv("DB_USERNAME")
	password := config.GetEnv("DB_PASSWORD")
	host := config.GetEnv("DB_HOST")
	port := config.GetEnv("DB_PORT")
	database := config.GetEnv("DB_DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", username, password, host, port, database)
	fmt.Println(dsn)
	Db, err := gorm.Open("mysql", dsn)
	logger.PanicError(err, "mysql connect err", true)

	logger.Info("连接Mysql 成功", "mysql connect")
	return Db
}
