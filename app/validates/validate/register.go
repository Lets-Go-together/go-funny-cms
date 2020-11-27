package validate

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

var CustomValidator *customValidator

// 所有验证器
var validations []Validation

type customValidator struct {
	validate *validator.Validate
	trans    ut.Translator
}

func init() {

	validations = append(validations,
		// 国内手机号码
		validationOfRegexp("phone", "^1[0-9]{10}$", "{0} 必须是手机号码"),
		// 常规用户名
		validationOfRegexp("username", "^[a-zA-Z][a-zA-Z0-9_]{4,15}$", "{0} 必须只包含大小写字母, 数字, 下划线, 且长度为 4-15"),
		// 标准域名
		validationOfRegexp("domain", "[a-zA-Z0-9][-a-zA-Z0-9]{0,62}(/.[a-zA-Z0-9][-a-zA-Z0-9]{0,62})+/.?", "{0} 必须是标准域名"),
		// 强密码
		//validationOfRegexp("strong_password", "^(?![0-9]+$)(?![a-zA-Z]+$)[0-9A-Za-z]{8,16}$", "{0} 必须包含写字母和数字, 且长度为 8-16"),
		// 中国邮政编码
		//validationOfRegexp("cn_postal_code", "[1-9]\\d{5}(?!\\d)", "{0} 必须是中国邮政编码"),
		// 中国大陆身份证号
		validationOfRegexp("cn_id_number", "^\\d{15}|\\d{18}$", "{0} 必须是中国身份证号码"),
	)

	CustomValidator, _ = New()
}

func New() (cv *customValidator, err error) {

	v := validator.New()
	local := zh.New()
	uniTrans := ut.New(local, local)
	translator, _ := uniTrans.GetTranslator("zh")

	for i := range validations {
		validation := validations[i]
		err = validation.register(v, translator)
		if err != nil {
			return
		}
	}

	// registerTranslation chinese as default translators for validate.
	err = zhTranslations.RegisterDefaultTranslations(v, translator)

	if err != nil {
		return
	}
	cv = &customValidator{
		validate: v,
		trans:    translator,
	}
	return
}

func (that *customValidator) verify(s interface{}) (bool, string) {
	errs := that.validate.Struct(s)
	if errs != nil {
		errs := errs.(validator.ValidationErrors)
		if len(errs) > 0 {
			err := errs[0]
			return false, err.Translate(that.trans)
		}
	}
	return true, ""
}
