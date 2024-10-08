package commands

import (
	"asura/handler"
	"asura/services"
	"context"

	"github.com/bwmarrin/discordgo"
)

func init() {
	handler.RegisterCommand(handler.Command{
		Name:        "ping",
		Description: "Pong!",
		Run:         runPing,
		Cooldown:    3,
	})
}

func runPing(_ context.Context, s *discordgo.Session, it *discordgo.InteractionCreate) {
	s.InteractionRespond(it.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: services.T("Ping", it),
		},
	})
}
