package main

import (
	"asura/handler"
	"context"
	"log"

	"github.com/bwmarrin/discordgo"
)

func OnReady(s *discordgo.Session, r *discordgo.Ready) {
	log.Printf("[#%v] logged in as %s", s.ShardID, r.User.String())
}

func OnInteractionCreate(s *discordgo.Session, it *discordgo.InteractionCreate) {
	data := it.ApplicationCommandData()

	if command := handler.GetCommand(data.Name); command.Run != nil {
		ctx := context.Background()

		command.Run(ctx, s, it)
	}
}
