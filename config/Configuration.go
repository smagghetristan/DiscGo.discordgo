package config

import (
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
)

var Token = os.Getenv("TOKEN")
var Prefix = "g!"
var StartTime = time.Now()

var Menus []Menu
var Categories []Category
var Commands []Command
var HelpEmbeds []*discordgo.MessageEmbed

type Menu struct {
	Name        string
	Description string
	Main        bool
}

type Category struct {
	Name string
	Menu Menu
}

type Command struct {
	Category         Category
	Command          string
	ShortDescription string
	LongDescription  string
	Function         func(*discordgo.Session, *discordgo.MessageCreate)
}
