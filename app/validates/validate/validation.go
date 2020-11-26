package validate

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"gocms/pkg/logger"
	"regexp"
)

type Validation struct {
	tag           string
	translation   string
	override      bool
	validateFn    validator.Func
	registerFn    validator.RegisterTranslationsFunc
	translationFn validator.TranslationFunc
}

func (that *Validation) register(v *validator.Validate, t ut.Translator) (err error) {
	err = that.registerValidator(v)
	if err == nil {
		err = that.registerTranslation(v, t)
	}
	return
}

func (that *Validation) registerValidator(v *validator.Validate) error {
	return v.RegisterValidation(that.tag, that.validateFn)
}

func (that *Validation) registerTranslation(v *validator.Validate, t ut.Translator) (err error) {

	if that.translationFn != nil && that.registerFn != nil {

		err = v.RegisterTranslation(that.tag, t, that.registerFn, that.translationFn)

	} else if that.translationFn != nil && that.registerFn == nil {

		err = v.RegisterTranslation(that.tag, t, registrationFunc(that.tag, that.translation, that.override), that.translationFn)

	} else if that.translationFn == nil && that.registerFn != nil {

		err = v.RegisterTranslation(that.tag, t, that.registerFn, translateFunc)

	} else {
		err = v.RegisterTranslation(that.tag, t, registrationFunc(that.tag, that.translation, that.override), translateFunc)
	}

	return
}

// 创建正则验证器
func validationOfRegexp(tag string, regex string, translation string) Validation {

	re, err := regexp.Compile(regex)

	if err != nil {
		logger.PanicError(err, "创建正则自定义验证器: "+regex, true)
	}
	// 闭包持有外部变量整个伴随自己的生命周期
	fn := func(fl validator.FieldLevel) bool {
		field := fl.Field().String()
		return re.MatchString(field)
	}
	return Validation{
		tag:         tag,
		translation: translation,
		validateFn:  fn,
	}
}

func translateFunc(ut ut.Translator, fe validator.FieldError) string {
	t, err := ut.T(fe.Tag(), fe.Field())
	if err != nil {
		logger.Error("翻译字段错误", fe.Error())
		return fe.(error).Error()
	}
	return t
}

func registrationFunc(tag string, translation string, override bool) validator.RegisterTranslationsFunc {
	return func(ut ut.Translator) error {
		return ut.Add(tag, translation, override)
	}
}
