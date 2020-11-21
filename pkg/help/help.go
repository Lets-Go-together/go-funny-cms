package help

import "os"

// 获取env
func GetEnv(key string) string {
	return os.Getenv(key)
}
