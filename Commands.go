package main

import "DiscGo.discordgo/config"
import "DiscGo.discordgo/cmd"

func AllCommands() {
	mainMenu := config.Menu{
		Name:        "Help Menu",
		Description: "The prefix of the bot is currently : **g!**",
		Main:        true,
	}

	gameMenu := config.Menu{
		Name:        "Games Menu",
		Description: "The prefix of the bot is currently : **g!**",
		Main:        true,
	}

	administrationCategory := config.Category{
		Name: "Administration :",
		Menu: mainMenu,
	}

	musicCategory := config.Category{
		Name: "Music :",
		Menu: mainMenu,
	}

	miscellaneousCategory := config.Category{
		Name: "Miscellaneous :",
		Menu: mainMenu,
	}

	randomCategory := config.Category{
		Name: "Others :",
		Menu: mainMenu,
	}

	supportCategory := config.Category{
		Name: "Support :",
		Menu: mainMenu,
	}

	tCategory := config.Category{
		Name: "Tic Tac Toe :",
		Menu: gameMenu,
	}

	hCategory := config.Category{
		Name: "Hang Man :",
		Menu: gameMenu,
	}

	kick := config.Command{
		Category:         administrationCategory,
		Command:          config.Prefix + "kick",
		ShortDescription: "**kick** : Kicks a user, you have to mention him.",
		LongDescription:  "This command will let you kick someone, if you do so, you will need to mention the user and only the user. Kicking doesn't mean banning, this means that the user will be able to recconect to the server with a correct invitation link.",
		Function:         commands.Kick,
	}

	ban := config.Command{
		Category:         administrationCategory,
		Command:          config.Prefix + "ban",
		ShortDescription: "**ban** : Bans a user, you have to mention him.",
		LongDescription:  "This command will let you ban someone, if you do so, you will need to mention the user and only the user. Banning means that the user won't be able to recconect to the server with a correct invitation link.",
		Function:         commands.Ban,
	}

	clear := config.Command{
		Category:         administrationCategory,
		Command:          config.Prefix + "clear",
		ShortDescription: "**clear** *x* : The bot will clear x message from a channel. (Maximum : 100)",
		LongDescription:  "This command is useful when you want to clear a channel that has been spammed. You must specify a number after the command because it won't work otherwise.",
		Function:         commands.Clear,
	}

	ping := config.Command{
		Category:         administrationCategory,
		Command:          config.Prefix + "ping",
		ShortDescription: "**ping** : The bot will reply by saying \"Pong\".",
		LongDescription:  "This command is useful when you want to know if you are having some kind of high latency, though it's not really useful since you can see your ping.",
		Function:         commands.Ping,
	}

	help := config.Command{
		Category:         administrationCategory,
		Command:          config.Prefix + "help",
		ShortDescription: "**help** : The bot will send you this menu.",
		LongDescription:  "With this command, the bot will reply to you sending the whole help menu, allowing you to find other commands.",
		Function:         commands.Help,
	}

	play := config.Command{
		Category:         musicCategory,
		Command:          config.Prefix + "play",
		ShortDescription: "**play** *youtube-url* : Will play the song from the youtube link.",
		LongDescription:  "To use the youtube player function you will need to specify a youtube link. For example : `" + config.Prefix + "play https://www.youtube.com/watch?v=v2AC41dglnM`",
		Function:         commands.Play,
	}

	pause := config.Command{
		Category:         musicCategory,
		Command:          config.Prefix + "pause",
		ShortDescription: "**pause** : Will pause the current song.",
		LongDescription:  "This will pause the current song.",
		Function:         commands.Pause,
	}

	resume := config.Command{
		Category:         musicCategory,
		Command:          config.Prefix + "resume",
		ShortDescription: "**resume** : Will resume the current song.",
		LongDescription:  "This will resume the song that was played and paused.",
		Function:         commands.Resume,
	}

	skip := config.Command{
		Category:         musicCategory,
		Command:          config.Prefix + "skip",
		ShortDescription: "**skip** : Will skip the current song.",
		LongDescription:  "Will skip to the next song in the queue",
		Function:         commands.Skip,
	}

	stop := config.Command{
		Category:         musicCategory,
		Command:          config.Prefix + "stop",
		ShortDescription: "**stop** : Will stop playing.",
		LongDescription:  "This command will empty the queue and stop playing the current song.",
		Function:         commands.Stop,
	}

	queue := config.Command{
		Category:         musicCategory,
		Command:          config.Prefix + "queue",
		ShortDescription: "**queue** : Will show you the entire queue",
		LongDescription:  "It will show you the queue and the duration of each song.",
		Function:         commands.QueueCMD,
	}

	say := config.Command{
		Category:         miscellaneousCategory,
		Command:          config.Prefix + "say",
		ShortDescription: "**say** : The bot will repeat what you said after this command.",
		LongDescription:  "It will repeat everything.",
		Function:         commands.Say,
	}

	success := config.Command{
		Category:         miscellaneousCategory,
		Command:          config.Prefix + "success",
		ShortDescription: "**success** title;text: The bot will create a minecraft-like success.",
		LongDescription:  "To use this command, an example is what would help you to understand. For example : `" + config.Prefix + "success You've made it!;First success made.`",
		Function:         commands.Success,
	}

	meme := config.Command{
		Category:         miscellaneousCategory,
		Command:          config.Prefix + "meme",
		ShortDescription: "**meme** : The bot will send a random meme.",
		LongDescription:  "It will.",
		Function:         commands.Meme,
	}

	games := config.Command{
		Category:         miscellaneousCategory,
		Command:          config.Prefix + "games",
		ShortDescription: "**games** : The bot will send you the game help page.",
		LongDescription:  "Try it, it will.",
		Function:         commands.GamesHelp,
	}

	info := config.Command{
		Category:         randomCategory,
		Command:          config.Prefix + "info",
		ShortDescription: "**info** : The bot will send you all of its statistics.",
		LongDescription:  "Try it, it will.",
		Function:         commands.Info,
	}

	supportMessage := config.Command{
		Category:         supportCategory,
		ShortDescription: "The bot is currently being updated to Golang, to follow the development, you can join this server : https://discord.gg/Q8NbBeQ",
	}

	tstart := config.Command{
		Category:         tCategory,
		Command:          config.Prefix + "t start",
		ShortDescription: "**t start** : Will make you start a game.",
		LongDescription:  "Try it, it will.",
		Function:         commands.TCommandStart,
	}

	tp2 := config.Command{
		Category:         tCategory,
		Command:          config.Prefix + "t p2",
		ShortDescription: "**t p2** : Will make you join a game.",
		LongDescription:  "Try it, it will.",
		Function:         commands.TCommandP2,
	}

	t := config.Command{
		Category:         tCategory,
		Command:          config.Prefix + "t",
		ShortDescription: "**t** *X* : Will make you play a move in a game.",
		LongDescription:  "Try it, it will.",
		Function:         commands.TCommandGame,
	}

	hplay := config.Command{
		Category:         hCategory,
		Command:          config.Prefix + "h play",
		ShortDescription: "**h play** : Will start a game and select a random word.",
		LongDescription:  "Try it, it will.",
		Function:         commands.HMPlay,
	}

	h := config.Command{
		Category:         hCategory,
		Command:          config.Prefix + "h",
		ShortDescription: "**h** *letter* : Useful to try and guess.",
		LongDescription:  "Try it, it will.",
		Function:         commands.HM,
	}

	config.AddCommand(kick)
	config.AddCommand(ban)
	config.AddCommand(ping)
	config.AddCommand(clear)
	config.AddCommand(help)

	config.AddCommand(play)
	config.AddCommand(pause)
	config.AddCommand(resume)
	config.AddCommand(skip)
	config.AddCommand(stop)
	config.AddCommand(queue)

	config.AddCommand(say)
	config.AddCommand(success)
	config.AddCommand(meme)
	config.AddCommand(games)

	config.AddCommand(info)

	config.AddCommand(supportMessage)

	config.AddCommand(tstart)
	config.AddCommand(tp2)
	config.AddCommand(t)

	config.AddCommand(hplay)
	config.AddCommand(h)

	config.CreateEmbeds()
}
