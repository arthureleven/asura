package handler

import (
	"context"

	"github.com/bwmarrin/discordgo"
)

type Command struct {
	Name        string
	Description string
	Options     []*discordgo.ApplicationCommandOption
	Run         func(context.Context, *discordgo.Session, *discordgo.InteractionCreate)
	Developer   bool
	Cooldown    int
	Cache       []string
}

var Commands = map[string]Command{}

func RegisterCommand(command Command) {
	Commands[command.Name] = command
}

func GetCommand(name string) Command {
	if command, ok := Commands[name]; ok {
		return command
	}

	return Command{}
}
