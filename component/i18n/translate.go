package i18n

import (
	ko_translations "github.com/kenSevLeb/go-framework/component/i18n/ko"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/ja"
	"github.com/go-playground/locales/ko"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	ja_translations "github.com/go-playground/validator/v10/translations/ja"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	zh_tw_translations "github.com/go-playground/validator/v10/translations/zh_tw"
)

const (
	LANG_ZH    = "zh"         //简体
	LANG_ZH_TW = "zh_hant_tw" //繁体
	LANG_EN    = "en"         //英语
	LANG_JA    = "ja"         //日语
	LANG_KO    = "ko"         //韩语
)

var mappingLanguages = map[string]string{
	"zh-Hant-TW": LANG_ZH_TW,
}

func getMappingLanguage(lang string) string {
	if item, exist := mappingLanguages[lang]; exist {
		return item
	}
	return lang
}

func TranslateValidatorErrors(language string, err error) string {
	for _, err := range err.(validator.ValidationErrors) {
		return err.Translate(getTranslator(language))
	}
	return ""

}

var uni = ut.New(en.New(), zh.New(), zh_Hant_TW.New(), ja.New(), ko.New())

func getTranslator(language string) ut.Translator {
	trans, found := uni.GetTranslator(language)
	if found {
		return trans
	}
	return uni.GetFallback()
}

func RegisterTranslations(validate *validator.Validate) {
	_ = zh_translations.RegisterDefaultTranslations(validate, getTranslator(LANG_ZH))
	_ = en_translations.RegisterDefaultTranslations(validate, getTranslator(LANG_ZH))
	_ = ja_translations.RegisterDefaultTranslations(validate, getTranslator(LANG_ZH))
	_ = zh_tw_translations.RegisterDefaultTranslations(validate, getTranslator(LANG_ZH))
	_ = ko_translations.RegisterDefaultTranslations(validate, getTranslator(LANG_ZH))
}
