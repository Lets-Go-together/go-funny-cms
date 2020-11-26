package auth

import (
	"fmt"
	"gocms/app/models/admin"
	"gocms/pkg/config"
	"gocms/pkg/help"
	"gocms/pkg/logger"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

// 生成并写入密钥
func GerateSign() {
	sign := help.Md5V(time.Now().String() + config.GetString("APP_NAME"))

	fmt.Println("sign："+sign, "\n"+"使用: 生成的密钥写入你的env配置中， key为JWT_SIGN")
}

// 创建admin 用户
func GerateAdminUser(account string) {
	if len(account) == 0 {
		fmt.Println("请输入账号")
		return
	}

	p := "12345678"
	password := CreatePassword(p)
	adminModel := admin.Admin{
		Account:  account,
		Password: password,
	}
	fmt.Println(adminModel)
	config.Db.FirstOrCreate(&adminModel, admin.Admin{
		Account: account,
	})

	fmt.Printf("account: %s \npassword: %s \n", account, p)

	fmt.Println(ValidatePassword(password, p))
}

// 创建密码
func CreatePassword(password string) string {
	newPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	logger.PanicError(err, "创建密码", true)

	return string(newPassword)
}

// 密码验证
func ValidatePassword(hashPassword string, newPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(newPassword))
	logger.PanicError(err, "密码验证", false)
	logger.Info("password validate", fmt.Sprintf("hashPassword %s - newPassword %s", hashPassword, newPassword))
	if err != nil {
		return false
	}
	return true
}

// 验证token
func ValidateToken(token string) string {
	if len(token) < 0 {
		return ""
	}

	t := strings.Split(token, "Bearer ")
	if len(t) > 0 {
		return t[1]
	}
	return ""
}
