package commands

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"DiscGo.discordgo/config"
	"github.com/bwmarrin/discordgo"
)

type MemeStruct struct {
	PostLink  string `json:"postLink"`
	Subreddit string `json:"subreddit"`
	Title     string `json:"title"`
	Url       string `json:"url"`
}

func Say(s *discordgo.Session, m *discordgo.MessageCreate) {
	message := strings.Replace(m.Content, config.Prefix+"say ", "", 1)
	_, err := s.ChannelMessageSend(m.ChannelID, message)
	if err != nil {
		return
	}
	err = s.ChannelMessageDelete(m.ChannelID, m.ID)
	if err != nil {
		return
	}
}

func Success(s *discordgo.Session, m *discordgo.MessageCreate) {
	message := strings.Replace(m.Content, config.Prefix+"success ", "", 1)
	title := strings.Split(message, ";")[0]
	text := strings.Split(message, ";")[1]
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1).Intn(39)
	url := "https://www.minecraftskinstealer.com/achievement/a.php?i=" + strconv.Itoa(r1) + "&h=" + url.QueryEscape(title) + "&t=" + url.QueryEscape(text)
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	AvatarURL := m.Author.AvatarURL("512")
	Author := &discordgo.MessageEmbedAuthor{
		Name:    m.Author.Username,
		IconURL: AvatarURL,
	}
	Image := &discordgo.File{
		Reader: resp.Body,
		Name:   "achievement.png",
	}
	Params := &discordgo.MessageSend{
		Files: []*discordgo.File{Image},
		Embed: &discordgo.MessageEmbed{
			Author: Author,
			Image: &discordgo.MessageEmbedImage{
				URL: "attachment://achievement.png",
			},
			Color: 0xFFDD00,
		},
	}

	_, err = s.ChannelMessageSendComplex(m.ChannelID, Params)
	if err != nil {
		return
	}
	err = s.ChannelMessageDelete(m.ChannelID, m.ID)
	if err != nil {
		return
	}
}

func Meme(s *discordgo.Session, m *discordgo.MessageCreate) {
	response, err := http.Get("https://meme-api.herokuapp.com/gimme")
	if err != nil {
		//
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		JSON := MemeStruct{}
		err := json.Unmarshal(data, &JSON)
		if err != nil {
			//
		}
		AvatarURL := m.Author.AvatarURL("512")
		Author := &discordgo.MessageEmbedAuthor{
			Name:    m.Author.Username,
			IconURL: AvatarURL,
		}
		Image := &discordgo.MessageEmbedImage{
			URL: JSON.Url,
		}
		embed := &discordgo.MessageEmbed{
			Author: Author,
			Image:  Image,
			Color:  0xFFDD00,
		}
		_, err = s.ChannelMessageSendEmbed(m.ChannelID, embed)
		if err != nil {
			return
		}
		err = s.ChannelMessageDelete(m.ChannelID, m.ID)
		if err != nil {
			return
		}
	}
}
