package database

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gocms/pkg/config"
	"gocms/pkg/logger"
)

func Initialize() {
	config.Db = MysqlClient()
	config.Redis = RedisClient()
}

// 初始化 Redis 服务器
func RedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     config.GetString("REDIS_HOST") + ":" + config.GetString("REDIS_PORT"),
		Password: config.GetString("REDIS_PASSWORD"),
		DB:       config.GetInt("REDIS_DB"),
	})

	_, err := client.Ping().Result()
	logger.PanicError(err, "Redis 连接", true)

	logger.Info("连接Redis 成功", "redis connect")
	return client
}

// mysql 服务器
func MysqlClient() *gorm.DB {
	username := config.GetString("DB_USERNAME")
	password := config.GetString("DB_PASSWORD")
	host := config.GetString("DB_HOST")
	port := config.GetString("DB_PORT")
	database := config.GetString("DB_DATABASE")
	charset := config.GetString("DB_CHARSET")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true", username, password, host, port, database, charset)
	Db, err := gorm.Open("mysql", dsn)
	logger.PanicError(err, "mysql connect err \n dns : "+dsn, true)

	logger.Info("连接Mysql 成功", "mysql connect")
	return Db
}
