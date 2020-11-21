package config

import (
	"github.com/joho/godotenv"
	"os"
)

// 初始化env配置
func InitEnv() {
	err := godotenv.Load()
	CheckError(err, "load env")
}

// 获取env
func GetEnv(key string) string {
	return os.Getenv(key)
}
