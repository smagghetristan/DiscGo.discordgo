package commands

import (
	"strconv"
	"strings"

	"DiscGo.discordgo/config"
	"github.com/bwmarrin/discordgo"
	"github.com/fatih/color"
)

type Page struct {
	Message string
}

func Play(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Find the guild for that channel.
	g, err := s.Guild(m.GuildID)
	if err != nil {
		// Could not find guild.
		return
	}
	// Look for the message sender in that guild's current voice states.
	for _, vs := range g.VoiceStates {
		if vs.UserID == m.Author.ID {
			url := strings.Replace(m.Content, config.Prefix+"play", "", 1)
			PlayYoutubeLink(s, g.ID, m.ChannelID, vs.ChannelID, url)
			if err != nil {
				color.Red("Error playing sound:", err)
			}
			return
		}
	}
}

func Pause(s *discordgo.Session, m *discordgo.MessageCreate) {
	PausePlaying(m.GuildID, m.ChannelID, s)
}

func Resume(s *discordgo.Session, m *discordgo.MessageCreate) {
	ResumePlaying(m.GuildID, m.ChannelID, s)
}

func Skip(s *discordgo.Session, m *discordgo.MessageCreate) {
	SkipPlaying(m.GuildID, m.ChannelID, s)
}

func Stop(s *discordgo.Session, m *discordgo.MessageCreate) {
	StopPlaying(m.GuildID, m.ChannelID, s)
}

func QueueCMD(s *discordgo.Session, m *discordgo.MessageCreate) {
	i := 0
	exist := false
	for i = 0; i < len(AllQueues); i++ {
		if AllQueues[i].GuildID == m.GuildID {
			exist = true
			break
		}
	}
	if exist {
		page, err := strconv.Atoi(strings.Replace(m.Content, config.Prefix+"queue ", "", 1))
		var pages []Page
		if err != nil {
			page = 1
		}
		if page < 1 {
			page = 1
		}
		CurrentPage := 0
		Cf := Page{
			Message: "",
		}
		pages = append(pages, Cf)
		for k := 0; k < len(AllQueues[i].Queue); k++ {
			if (k-1)%10 == 0 && (k-1) != 0 {
				CurrentPage++
				Cf := Page{
					Message: "",
				}
				pages = append(pages, Cf)
			}
			if k != 0 {
				pages[CurrentPage].Message += strconv.Itoa(k) + ". [" + AllQueues[i].Queue[k].Name + "](" + AllQueues[i].Queue[k].URL + `) | ` + AllQueues[i].Queue[k].Duration + `
				`
			}
		}
		CurrentMusic := &discordgo.MessageEmbedField{
			Name:  "__Now Playing:__",
			Value: "[" + AllQueues[i].Queue[0].Name + "](" + AllQueues[i].Queue[0].URL + `) | ` + AllQueues[i].Queue[0].Duration + ``,
		}
		//RANDOM FIELD
		Next := &discordgo.MessageEmbedField{
			Name:  "__Next:__",
			Value: pages[page-1].Message,
		}
		AllFields := []*discordgo.MessageEmbedField{CurrentMusic, Next}
		AvatarURL := m.Author.AvatarURL("512")
		embed := &discordgo.MessageEmbed{
			Title:  "Queue :",
			Fields: AllFields,
			Footer: &discordgo.MessageEmbedFooter{
				IconURL: AvatarURL,
				Text:    "Page : " + strconv.Itoa(page) + "/" + strconv.Itoa(len(pages)),
			},
			Color: 0xFFDD00,
		}
		_, err = s.ChannelMessageSendEmbed(m.ChannelID, embed)
		if err != nil {
			return
		}
	} else {
		s.ChannelMessageSend(m.ChannelID, ":no_entry: There's nothing in the queue.")
	}
}
