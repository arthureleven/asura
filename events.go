package main

import (
	"asura/handler"
	"asura/services"
	"context"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

func OnReady(s *discordgo.Session, r *discordgo.Ready) {
	log.Printf("[#%v] logged in as %s", s.ShardID, r.User.String())
}

func OnInteractionCreate(s *discordgo.Session, it *discordgo.InteractionCreate) {
	data := it.ApplicationCommandData()

	if command := handler.GetCommand(data.Name); command.Run != nil {
		ctx := context.Background()

		if cooldown, ok := handler.GetCooldown(ctx, it.Member.User.ID, command); ok {
			since := int(time.Since(cooldown).Seconds())
			rem := command.Cooldown - since

			s.InteractionRespond(it.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: services.T("Cooldown", it, rem),
				},
			})
		} else {
			handler.SetCooldown(ctx, it.Member.User.ID, command)

			command.Run(ctx, s, it)
		}
	}
}
