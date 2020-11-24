package validate

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"gocms/pkg/logger"
)

var (
	validate *validator.Validate
	trans    ut.Translator
)

// 注册自定义验证器
func registerCustomValidate() *validator.Validate {
	validate = validator.New()
	registerPhone()

	return validate
}

// 注册手机号码验证
func registerPhone() {
	_ = validate.RegisterValidation("phone", func(fl validator.FieldLevel) bool {
		return false
	})
}

// 注册翻译器
func registerCustomTrans() ut.Translator {
	zh := zh.New()
	uni := ut.New(zh, zh)
	trans, _ := uni.GetTranslator("zh")

	var err error
	if err = trans.Add("phone", "{0} 必须符合手机号码格式", false); err != nil {
		logger.PanicError(err, "register phone", true)
	}

	return trans
}
