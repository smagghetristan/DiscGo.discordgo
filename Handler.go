package main

import (
	"os"
	"strings"

	"DiscGo.discordgo/cmd"
	"DiscGo.discordgo/config"
	"github.com/bwmarrin/discordgo"
)

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.

func CommandHandle(session *discordgo.Session, m *discordgo.MessageCreate) {
	user := m.Author
	if BotUsr, err := session.User("@me"); user.ID == BotUsr.ID || user.Bot || err != nil {
		return
	}
	if strings.HasPrefix(m.Content, config.Prefix+"restart") {
		if m.Author.ID == "144472011924570113" {
			os.Exit(0)
		}
	}
	commands.Help(session, m)
	commands.Sound(session, m)
	commands.Administration(session, m)
	commands.Miscellaneous(session, m)
	commands.TicTacToe(session, m)
	commands.HangMan(session, m)
}
func messageCreate(session *discordgo.Session, m *discordgo.MessageCreate) {
	go CommandHandle(session, m)
	return
}
