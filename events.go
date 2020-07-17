package main

import (
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func readyEvent(s *discordgo.Session, m *discordgo.Ready) {
	log.Info("Ready! " + "Loaded " + strconv.Itoa(len(activeCommands)) + " commands")
}

func messageCreateEvent(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	if strings.HasPrefix(m.Content, prefix) {
		cmd := strings.Trim(m.Content, prefix)
		if pick, ok := activeCommands[cmd]; ok {
			if (m.GuildID != "" && pick.FireOnGuild) || (m.GuildID == "" && pick.FireOnDM) {
				start := time.Now()
				pick.Exec(s, m)
				elapsed := time.Since(start)
				log.Info("Executed command \"", cmd, "\". Took me ", elapsed.Milliseconds(), "ms to finish")
			}
		}
	}
}

func guildJoinEvent(s *discordgo.Session, m *discordgo.GuildCreate) {
	log.Info("Bot joined a guild! " + m.Name)
}
