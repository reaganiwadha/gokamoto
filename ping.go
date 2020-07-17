package main

import (
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

func init() {
	newGuildCommand("ping", "Pings the bot", ping).add()
}

func ping(s *discordgo.Session, m *discordgo.MessageCreate) {
	now := time.Now()

	sentMessage, err := s.ChannelMessageSend(m.ChannelID, "🏓 Pinging")

	if err != nil {
		log.Error(err)
	}

	newNow := strconv.FormatInt(time.Since(now).Milliseconds(), 10)
	s.ChannelMessageEdit(m.ChannelID, sentMessage.ID, "Pong! 🏓 "+newNow+"ms")
}
