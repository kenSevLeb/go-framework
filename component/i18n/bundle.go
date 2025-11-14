package i18n

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
	"log"
)

// Bundle
type Bundle struct {
	bundle *i18n.Bundle
}

type Config struct {
	Files []string
}

func newI18n(conf *Config) *Bundle {
	instance := new(Bundle)
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)

	for _, path := range conf.Files {
		_, err := bundle.LoadMessageFile(path)
		if err != nil {
			log.Println("LoadMessageFile:", path, err.Error())
		}
	}
	instance.bundle = bundle

	return instance
}

// MustLocalize 本地化
func (component *Bundle) MustLocalize(language string, message string) string {
	translate, err := component.locale(language).Localize(&i18n.LocalizeConfig{MessageID: message})
	if err != nil { // 如果没有配置，则原样返回
		return message
	}
	return translate
}

// locale
func (component *Bundle) locale(language string) *i18n.Localizer {
	locale := i18n.NewLocalizer(component.bundle, language)
	return locale
}
