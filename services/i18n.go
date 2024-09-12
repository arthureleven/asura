package services

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/bwmarrin/discordgo"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var Bundle *i18n.Bundle

var Idioms = map[string]string{
	"pt-BR": "pt",
	"en-US": "en",
}

var Paths = []string{
	"general",
	"commands",
}

func T(id string, it *discordgo.InteractionCreate, data ...interface{}) string {
	localizer := i18n.NewLocalizer(Bundle, string(it.Interaction.Locale))
	config := &i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: id,
		},
	}

	if len(data) > 0 {
		config.TemplateData = data[0]
	}

	return localizer.MustLocalize(config)
}

func init() {
	Bundle = i18n.NewBundle(language.English)

	Bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	for _, lang := range Idioms {
		for _, file := range Paths {
			Bundle.MustLoadMessageFile(fmt.Sprintf("locales/%s/%s.%s.toml", lang, file, lang))
		}
	}
}
