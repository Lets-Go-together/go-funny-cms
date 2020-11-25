package example

import (
	"fmt"
	"gocms/app/validates/validate"
	"gocms/pkg/logger"
)

// https://github.com/go-playground/validator
// 验证齐使用参考如下方法
func Validate() {
	type User struct {
		FirstName      string `validate:"required"`
		LastName       string `validate:"required"`
		Age            uint8  `validate:"gte=0,lte=130"`
		Email          string `validate:"required,phone"`
		FavouriteColor string `validate:"iscolor"`
		Phone          string `validate:"phone"`
	}

	user := &User{
		FirstName: "好的",
		LastName:  "222",
		Phone:     "123",
	}

	msg, isSuccess := validate.BaseValidate(user)
	fmt.Println(msg, isSuccess)
	logger.Debug(msg)
}
