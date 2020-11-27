package example

import (
	"gocms/app/validates/validate"
	"gocms/pkg/logger"
)

// https://github.com/go-playground/validator
// 验证齐使用参考如下方法
func Validate() {
	type User struct {
		FirstName string `validate:"required"`
		LastName  string `validate:"required"`
		//Age            uint8  `validate:"gte=0,lte=130"`
		//Email          string `validate:"required,phone"`
		//FavouriteColor string `validate:"iscolor"`
		Id    string `validate:"cn_id_number"`
		Phone string `validate:"phone"`
		Gt    int    `validate:"great_then=3"`
	}

	var user = &User{
		FirstName: "好的",
		LastName:  "222",
		Phone:     "13311313311",
		Id:        "5123512512351235123",
		Gt:        3,
	}
	succ, m := validate.Validate(user)
	if succ {
		logger.Debug("验证成功")
	} else {
		logger.Debug(m)
	}
}
