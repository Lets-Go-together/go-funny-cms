package auth

import (
	"fmt"
	"gocms/pkg/config"
	"gocms/pkg/help"
	"time"
)

// 生成并写入密钥
func GerateSign() {
	sign := help.Md5V(time.Now().String() + config.GetString("APP_NAME"))

	fmt.Println("sign："+sign, "\n"+"使用: 生成的密钥写入你的env配置中， key为JWT_SIGN")
}
