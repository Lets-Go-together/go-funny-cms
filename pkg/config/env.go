package config

import (
	"github.com/joho/godotenv"
	"gocms/pkg/logger"
	"os"
)

// 初始化env配置
func InitEnv() {
	err := godotenv.Load()
	logger.PanicError(err, "load env", true)
}

// 获取env
func GetEnv(key string) string {
	return os.Getenv(key)
}
