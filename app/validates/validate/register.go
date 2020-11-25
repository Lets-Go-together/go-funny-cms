package validate

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"gocms/pkg/logger"
	"regexp"
)

var (
	validate     *validator.Validate
	trans        ut.Translator
	validationFn = map[string]validator.Func{}
)

func init() {
	// 国内手机号码
	validationFn["phone"] = validationFnOfRegexp("^1[0-9]{10}$")
	// 常规用户名
	validationFn["username"] = validationFnOfRegexp("^[A-Za-z_0-9]{4,15}$")

	// TODO 2020-11-25 23:04:12 补全以下自定义验证
	//validationFn["password"] = validationFnOfRegexp()
	//validationFn["tel"] = validationFnOfRegexp()
	//validationFn["domain"] = validationFnOfRegexp()
	//validationFn["cn_id_number"] = validationFnOfRegexp()
	//validationFn["cn_postal_code"] = validationFnOfRegexp()
}

// 注册自定义验证器
func registerCustomValidate() *validator.Validate {
	validate = validator.New()
	for tag := range validationFn {
		err := validate.RegisterValidation(tag, validationFn[tag])
		if err == nil {
			logger.PanicError(err, "注册自定义验证器", true)
		}
	}
	return validate
}

// 返回匹配正则的验证器
func validationFnOfRegexp(s string) validator.Func {
	re, err := regexp.Compile(s)
	if err != nil {
		logger.PanicError(err, "创建正则自定义验证器: "+s, true)
	}
	return func(fl validator.FieldLevel) bool {
		field := fl.Field().String()
		return re.MatchString(field)
	}
}

// 注册翻译器
func registerCustomTrans() ut.Translator {
	zhTrans := zh.New()
	enTrans := en.New()
	uniTrans := ut.New(zhTrans, zhTrans, enTrans)

	trans, _ := uniTrans.GetTranslator("zh")
	var err error
	if err = trans.Add("phone", "{0} 必须符合手机号码格式", false); err != nil {
		logger.PanicError(err, "register phone", true)
	}

	return trans
}
