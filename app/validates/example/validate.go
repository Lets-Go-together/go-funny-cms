package example

import (
	"fmt"
	validate "gocms/app/validates/validate"
)

// https://github.com/go-playground/validator
// 验证齐使用参考如下方法
func Validate() {
	type User struct {
		FirstName      string `validate:"required"`
		LastName       string `validate:"required"`
		Age            uint8  `validate:"gte=0,lte=130"`
		Email          string `validate:"required,email"`
		FavouriteColor string `validate:"iscolor"`
	}

	user := &User{
		FirstName:      "好的",
		LastName:       "222",
		Age:            135,
		Email:          "1563",
		FavouriteColor: "",
	}

	msg, isSuccess := validate.Validate(user)
	fmt.Println(msg, isSuccess)
}
