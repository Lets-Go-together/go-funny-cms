package auth

import (
	"fmt"
	"gocms/app/models/admin"
	"gocms/pkg/config"
	"gocms/pkg/help"
	"gocms/pkg/logger"
	"golang.org/x/crypto/bcrypt"
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
		Password: string(password),
	}
	fmt.Println(adminModel)
	config.Db.FirstOrCreate(&adminModel, admin.Admin{
		Account: account,
	})

	fmt.Printf("account: %s \npassword: %s \n", account, p)
}

// 创建密码
func CreatePassword(password string) string {
	newPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	logger.PanicError(err, "创建密码", true)

	return string(newPassword)
}

// 密码验证
func ValidatePassword(oldPassword string, newPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(oldPassword), []byte(newPassword))
	if err != nil {
		return true
	}
	return false
}
