package validate

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

func init() {
	registerTrans()
}

// 初始化翻译器
func getTrans() ut.Translator {
	return registerCustomTrans()
}

// 获取验证器
func getValidate() *validator.Validate {
	return registerCustomValidate()
}

// 初始化注册验证器和翻译器
func registerTrans() {
	trans = getTrans()
	validate = getValidate()
	_ = zh_translations.RegisterDefaultTranslations(validate, trans)
}

// 验证器
// 返回验证器验证结果错误消息 和 bool (是否验证成功)
func BaseValidate(validatModel interface{}) (string, bool) {
	errs := validate.Struct(validatModel)

	if errs != nil {
		errs := errs.(validator.ValidationErrors)
		if len(errs) > 0 {
			err := errs[0]
			return err.Translate(trans), false
		}
	}

	return "", true
}
