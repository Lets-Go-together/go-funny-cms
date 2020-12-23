package database

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gocms/pkg/config"
	"gocms/pkg/logger"
	"time"
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
	config.Dsn = dsn
	Db, err := gorm.Open("mysql", dsn)
	logger.PanicError(err, "mysql connect err \n dns : "+dsn, true)

	//设置连接池空闲连接数
	Db.DB().SetMaxIdleConns(config.GetInt("DB_MAX_IDLE_CONNS"))
	//设置打开最大连接数
	Db.DB().SetMaxOpenConns(config.GetInt("DB_MAX_OPEN_CONNS"))
	//连接可空闲最长时间
	Db.DB().SetConnMaxIdleTime(time.Duration(config.GetInt64("DB_CONN_MAX_IDLE_TIME")))
	//连接可以重用最长时间
	Db.DB().SetConnMaxLifetime(time.Duration(config.GetInt64("DB_CONN_MAX_LIFE_TIME")))
	Db.LogMode(true)

	// 全局禁用表名复数
	// Db.SingularTable(true)

	logger.Info("连接Mysql 成功", "mysql connect")

	return Db
}
