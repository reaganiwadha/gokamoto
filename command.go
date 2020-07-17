package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	activeCommands = make(map[string]command)
)

type command struct {
	Name        string
	Description string
	Exec        func(*discordgo.Session, *discordgo.MessageCreate)
	FireOnGuild bool
	FireOnDM    bool
}

func newCommand(name string, description string, f func(*discordgo.Session, *discordgo.MessageCreate)) command {
	return command{
		Name:        name,
		Description: description,
		Exec:        f,
		FireOnGuild: true,
		FireOnDM:    true,
	}
}

func newGuildCommand(name string, description string, f func(*discordgo.Session, *discordgo.MessageCreate)) command {
	return command{
		Name:        name,
		Description: description,
		Exec:        f,
		FireOnGuild: true,
		FireOnDM:    false,
	}
}

func newDMCommand(name string, description string, f func(*discordgo.Session, *discordgo.MessageCreate)) command {
	return command{
		Name:        name,
		Description: description,
		Exec:        f,
		FireOnGuild: false,
		FireOnDM:    true,
	}
}

func (c command) add() command {
	activeCommands[strings.ToLower(c.Name)] = c
	return c
}
