package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

type commandHelp struct {
	Usage []string
}

func init() {
	newGuildCommand("help", "Gives you the help", help).add()
}

func help(s *discordgo.Session, m *discordgo.MessageCreate) {
	cmds := []string{}

	for key := range activeCommands {
		cmds = append(cmds, key)
	}

	cmdsStrings := strings.Join(cmds, ", ")

	s.ChannelMessageSend(m.ChannelID, cmdsStrings)
}
