package configs

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

func InitI18n() *i18n.Bundle {
	// Inisialisasi bundle
	i18nBundle := i18n.NewBundle(language.English)
	i18nBundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)
	i18nBundle.MustLoadMessageFile("file/locales/en.yaml")
	i18nBundle.MustLoadMessageFile("file/locales/id.yaml")
	return i18nBundle
}

func Translate(bundle *i18n.Bundle, lang string, key string, data map[string]any) string {
	localize := i18n.NewLocalizer(bundle, lang)
	return localize.MustLocalize(&i18n.LocalizeConfig{
		MessageID:    key,
		TemplateData: data,
	})
}
