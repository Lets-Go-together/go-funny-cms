package config

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"gocms/pkg/logger"
	"reflect"
	"strconv"
	"time"
)

type DataBase struct {
	*gorm.DB
}

func InitDatabase() {
	Db = MysqlClient()
	Redis = RedisClient()
}

// MysqlClient 初始化 Redis 服务器
func RedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     GetString("REDIS_HOST") + ":" + GetString("REDIS_PORT"),
		Password: GetString("REDIS_PASSWORD"),
		DB:       GetInt("REDIS_DB"),
	})

	_, err := client.Ping().Result()
	logger.PanicError(err, "Redis 连接", true)

	logger.Info("连接Redis 成功", "redis connect")
	return client
}

// MysqlClient 初始化 mysql 服务器
func MysqlClient() *DataBase {
	username := GetString("DB_USERNAME")
	password := GetString("DB_PASSWORD")
	host := GetString("DB_HOST")
	port := GetString("DB_PORT")
	database := GetString("DB_DATABASE")
	charset := GetString("DB_CHARSET")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true", username, password, host, port, database, charset)
	Db, err := gorm.Open("mysql", dsn)
	logger.PanicError(err, "mysql connect err \n dns : "+dsn, true)

	//设置连接池空闲连接数
	Db.DB().SetMaxIdleConns(GetInt("DB_MAX_IDLE_CONNS"))
	//设置打开最大连接数
	Db.DB().SetMaxOpenConns(GetInt("DB_MAX_OPEN_CONNS"))
	//连接可空闲最长时间
	Db.DB().SetConnMaxIdleTime(time.Duration(GetInt64("DB_CONN_MAX_IDLE_TIME")))
	//连接可以重用最长时间
	Db.DB().SetConnMaxLifetime(time.Duration(GetInt64("DB_CONN_MAX_LIFE_TIME")))
	Db.LogMode(true)

	// 全局禁用表名复数
	// Db.SingularTable(true)

	logger.Info("连接Mysql 成功", "mysql connect")

	return &DataBase{
		Db,
	}
}

func (that *DataBase) Where(query interface{}, param ...interface{}) *DataBase {
	return &DataBase{
		that.DB.Where(query, param),
	}
}

func (that *DataBase) WhereCombOr(query interface{}) *DataBase {
	t := reflect.TypeOf(query)
	v := reflect.ValueOf(query)

	flag := false
	db := that.DB

	for i := 0; i < t.NumField(); i++ {
		where := t.Field(i).Name
		param := fieldValueStr(v.Field(i))
		if flag {
			db = db.Or(where+"=?", param)
			continue
		}
		flag = true
		db = db.Where(where+"=?", param)
	}

	return &DataBase{db}
}

func fieldValueStr(value reflect.Value) (res string) {
	kind := value.Kind()
	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		res = strconv.FormatInt(value.Int(), 10)

	case reflect.Float64, reflect.Float32:
		res = strconv.FormatFloat(value.Float(), 'f', 10, 64)

	case reflect.String:
		res = value.String()

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		res = strconv.FormatUint(value.Uint(), 10)

	case reflect.Bool:
		res = strconv.FormatBool(value.Bool())

	default:
		logger.PanicError(errors.New("unsupported type "+kind.String()), "fieldValueStr db.go:40", true)
	}
	return
}
