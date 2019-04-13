package commands

import (
	"runtime"
	"strconv"
	"strings"
	"time"

	"DiscGo.discordgo/config"
	"github.com/bwmarrin/discordgo"
)

func Help(session *discordgo.Session, m *discordgo.MessageCreate) {
	Prefix := config.Prefix
	StartTime := config.StartTime

	if strings.HasPrefix(m.Content, Prefix+"help") {
		//ADMIN FIELD
		AdminField := &discordgo.MessageEmbedField{
			Name: "Admin commands:",
			Value: `**kick** : Kicks a user, you have to mention him.
		        **ban** : Bans a user, you have to mention him.
		        **ping** : Will make the bot reply "Pong !"
		        **clear** : Will make the bot clear a channel's message (clear **8**)`,
		}
		//FUN FIELD
		FunField := &discordgo.MessageEmbedField{
			Name: "Fun commands:",
			Value: `**say**: The bot will repeat what you said after the command.
						**success** title;text: The bot will create a minecraft-like success.
						**meme**: The bot will send a random meme.
						**games**: Will send you the game help page.`,
		}
		MusicField := &discordgo.MessageEmbedField{
			Name: "Music commands:",
			Value: `**play** *youtube-url* : Will play the song from the youtube link.
		        **pause** : Will pause the current song.
		        **resume** : Will resume the current song.
		        **stop** : Will stop the current song.
		        **queue** : The bot will send you the queue.`,
		}
		//RANDOM FIELD
		RandomField := &discordgo.MessageEmbedField{
			Name:  "Random Commands :",
			Value: `**info** : Will return the bots info.`,
		}
		//UPDATE FIELD
		AdField := &discordgo.MessageEmbedField{
			Name:  "Support :",
			Value: `The bot is currently being updated to Golang, to follow the development, you can join this server : https://discord.gg/Q8NbBeQ`,
		}

		AllFields := []*discordgo.MessageEmbedField{AdminField, MusicField, FunField, RandomField, AdField}

		embed := &discordgo.MessageEmbed{
			Title:       "Help Menu",
			Description: "The prefix of the bot is currently : **g!**",
			Fields:      AllFields,
			Color:       0xFFDD00,
		}

		session.ChannelMessageSendEmbed(m.ChannelID, embed)
	}

	if strings.HasPrefix(m.Content, Prefix+"games") {
		//TICTACTOE FIELD
		TicTacToeField := &discordgo.MessageEmbedField{
			Name: "TicTacToe :",
			Value: `**t start** : Will make you start a game.
						**t p2** : Will make you join a game.
						**t** *number*: Will make you play the number in a game.`,
		}
		//HANGMAN FIELD
		HangManField := &discordgo.MessageEmbedField{
			Name: "HangMan :",
			Value: `**h play** : Will start a game and select a random word.
						**h** *letter* : Useful to try and guess.`,
		}
		AllFields := []*discordgo.MessageEmbedField{TicTacToeField, HangManField}

		embed := &discordgo.MessageEmbed{
			Title:       "Game Menu",
			Description: "The prefix of the bot is currently : **g!**",
			Fields:      AllFields,
			Color:       0xFFDD00,
		}

		session.ChannelMessageSendEmbed(m.ChannelID, embed)
	}

	if strings.HasPrefix(m.Content, Prefix+"info") {
		users := 0
		for _, guild := range session.State.Ready.Guilds {
			users += len(guild.Members)
		}
		ServerAmount := len(session.State.Guilds)
		Uptime := time.Since(StartTime)

		embed := &discordgo.MessageEmbed{
			Title: "Bot Statistics :",
			Description: `**Servers** : ` + strconv.Itoa(ServerAmount) + `
				**Users** : ` + strconv.Itoa(users) + `
				**Tasks** : ` + strconv.Itoa(runtime.NumGoroutine()) + `
				**Uptime** : ` + strconv.Itoa(int(Uptime.Hours())) + `:` + strconv.Itoa(int(Uptime.Minutes())%60) + `:` + strconv.Itoa(int(Uptime.Seconds())%60),
			Color: 0xFFDD00,
		}
		session.ChannelMessageSendEmbed(m.ChannelID, embed)
	}
	return
}
