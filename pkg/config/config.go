package config

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/panjf2000/ants/v2"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"gocms/app/models"
	"gocms/pkg/logger"
	"net/http"
)

// viper 作为全局的 app 容器存在
var (
	Db        *gorm.DB
	Redis     *redis.Client
	Viper     *viper.Viper
	AuthAdmin *models.AuthAdmin
	Pool      *ants.Pool
	Router    *gin.Engine
	Request   *http.Request
)

func init() {
	InitViper()
}

func Initialize() {}

// 初始化全局容器
func InitViper() {
	// step1: 初始化
	Viper = viper.New()

	// step2: 设置文件名称
	Viper.SetConfigName(GetEnvFile())

	// step3: 设置配置文件类型
	Viper.SetConfigType("env")

	// step4: 加载路径
	Viper.AddConfigPath(".")

	// step5: 加载文件配置
	// 参考: https://github.com/spf13/viper#writing-config-files
	err := Viper.ReadInConfig()
	logger.PanicError(err, "加载env配置", true)

	Viper.AutomaticEnv()
	Viper.AllowEmptyEnv(true)
}

// 获取env文件
func GetEnvFile() string {
	return ".env"
}

// 读取env 配置
func Env(envName string, defaultValue ...interface{}) interface{} {
	if len(defaultValue) > 0 {
		return Get(envName, defaultValue)
	}
	return Get(envName)
}

// 获取配置变量
func Get(key string, defaultValue ...interface{}) interface{} {
	if !Viper.IsSet(key) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		} else {
			return nil
		}
	}
	return Viper.Get(key)
}

// 添加配置变量
func Add(key string, value interface{}) {
	Viper.Set(key, value)
}

// ----如下这份代码来自于
// ----https://github.com/summerblue/goblog/blob/master/pkg/config/config.go
// GetString 获取 String 类型的配置信息
func GetString(path string, defaultValue ...interface{}) string {
	return cast.ToString(Get(path, defaultValue...))
}

// GetInt 获取 Int 类型的配置信息
func GetInt(path string, defaultValue ...interface{}) int {
	return cast.ToInt(Get(path, defaultValue...))
}

// GetInt64 获取 Int64 类型的配置信息
func GetInt64(path string, defaultValue ...interface{}) int64 {
	return cast.ToInt64(Get(path, defaultValue...))
}

// GetUint 获取 Uint 类型的配置信息
func GetUint(path string, defaultValue ...interface{}) uint {
	return cast.ToUint(Get(path, defaultValue...))
}

// GetBool 获取 Bool 类型的配置信息
func GetBool(path string, defaultValue ...interface{}) bool {
	return cast.ToBool(Get(path, defaultValue...))
}

// 历史代码
// 初始化env配置
//func InitEnv() {
//	err := godotenv.Load()
//	logger.PanicError(err, "load env", true)
//
//	// 记录一下当前的系统环境变量
//	logger.Info("当前env 环境", os.Getenv("APP_ENV"))
//}
//
//// 获取env
//func GetEnv(key string) string {
//	return os.Getenv(key)
//}
