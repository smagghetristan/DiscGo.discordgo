package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"DiscGo.discordgo/config"
	"github.com/bwmarrin/discordgo"
	"github.com/fatih/color"
)

func main() {
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)
	dg.AddHandler(botReady)
	AllCommands()
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

func botReady(session *discordgo.Session, evt *discordgo.Ready) {
	color.Green("Bot is now running.  Press CTRL-C to exit.")
	session.UpdateStatus(0, "g!help for help")
}
