package main

import (
	"asura/handler"
	"log"

	"github.com/bwmarrin/discordgo"
)

func OnReady(s *discordgo.Session, r *discordgo.Ready) {
	log.Printf("[#%v] logged in as %s", s.ShardID, r.User.String())
}

func OnInteractionCreate(s *discordgo.Session, it *discordgo.InteractionCreate) {
	data := it.ApplicationCommandData()
	command := handler.GetCommand(data.Name)

	if command.Run != nil {
		command.Run(s, it)
	}
}
