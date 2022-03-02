package i18n

import (
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var (
	bundle    *i18n.Bundle
	localizer *i18n.Localizer
)

func init() {
	bundle = i18n.NewBundle(language.Chinese)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustLoadMessageFile("/home/i18n/zh-CN.toml")
	localizer = i18n.NewLocalizer(bundle, "zh-CN")
}

// Load loads i18n files
func Load(path string) {
	bundle.MustLoadMessageFile(path)
}

// Tr translates a message
func Tr(i18nKey string, template map[string]interface{}) string {
	if len(template) == 0 {
		return localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: i18nKey})
	}
	return localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: i18nKey, TemplateData: template})
}
