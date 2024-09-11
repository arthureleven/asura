package commands

import (
	"asura/handler"
	"asura/services"

	"github.com/bwmarrin/discordgo"
)

func init() {
	handler.RegisterCommand(handler.Command{
		Name:        "ping",
		Description: "Pong!",
		Run:         run,
		Cooldown:    3,
	})
}

func run(s *discordgo.Session, it *discordgo.InteractionCreate) {
	s.InteractionRespond(it.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: services.T("Ping.pong", it),
		},
	})
}
