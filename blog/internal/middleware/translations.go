package middleware

/*
	国际化

	go-playground/locales:
		多语言包, 是从 CLDR 项目(Unicode 通用语言环境数据存储库)生成的一组多
		语言环境, 主要在 i18n 软件包中使用, 该库是与 universal-translator
		配套使用的;

	go-playground/universal-translator:
		通用翻译器, 是一个使用 CLDR 数据 + 复数规则的 Go 语言 i18n 转换器;

	go-playground/validator/v10/translations: validator 的翻译器


*/

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	validator "github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

func Translations() gin.HandlerFunc {
	return func(c *gin.Context) {
		uni := ut.New(en.New(), zh.New(), zh_Hant_TW.New())
		locale := c.GetHeader("locale") // TODO: header 字段中的地区信息
		trans, _ := uni.GetTranslator(locale)
		v, ok := binding.Validator.Engine().(*validator.Validate)
		if ok {
			// 调用 RegisterDefaultTranslations 注册验证器和对应语言类型的
			// Translator, 实现验证器的多语言支持
			switch locale {
			case "zh":
				_ = zh_translations.RegisterDefaultTranslations(v, trans)
				break
			case "en":
				_ = en_translations.RegisterDefaultTranslations(v, trans)
				break
			default:
				_ = zh_translations.RegisterDefaultTranslations(v, trans)
				break
			}
			c.Set("trans", trans) // pkg/app/form.go, Line36
		}

		c.Next()
	}
}
