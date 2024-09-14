package services

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/bwmarrin/discordgo"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var Bundle *i18n.Bundle

var Languages = map[string]string{
	"pt-BR": "pt",
	"en-US": "en",
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

	for _, language := range Languages {
		basepath := fmt.Sprintf("locales/%s", language)

		if files, err := os.ReadDir(basepath); err == nil {
			for _, file := range files {
				path := fmt.Sprintf("%s/%s", basepath, file.Name())

				Bundle.MustLoadMessageFile(path)
			}
		}
	}
}
