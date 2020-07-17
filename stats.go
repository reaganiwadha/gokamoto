package main

import (
	"os"
	"runtime"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/struCoder/pidusage"
)

func init() {
	newGuildCommand("stats", "Gives you a statistic", stats).add()
}

func stats(s *discordgo.Session, m *discordgo.MessageCreate) {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	sysInfo, _ := pidusage.GetStat(os.Getpid())

	memUsage := strconv.Itoa(int(sysInfo.Memory/1024/1024)) + "mb"
	pid := strconv.Itoa(os.Getpid())

	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name:    "Sakamoto",
			IconURL: "",
		},
		Color: 0x9292e0,
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "mem",
				Value:  memUsage,
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "pid",
				Value:  pid,
				Inline: true,
			},
		},
	}

	s.ChannelMessageSendEmbed(m.ChannelID, embed)
}
