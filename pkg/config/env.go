package config

import (
	"github.com/joho/godotenv"
	"gocms/pkg/logger"
	"os"
)

func init() {
	InitEnv()
}

// 初始化env配置
func InitEnv() {
	err := godotenv.Load()
	logger.PanicError(err, "load env", true)

	// 记录一下当前的系统环境变量
	logger.Info("当前env 环境", os.Getenv("APP_ENV"))
}

// 获取env
func GetEnv(key string) string {
	return os.Getenv(key)
}
