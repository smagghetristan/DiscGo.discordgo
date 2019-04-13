package commands

import (
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/smagghetristan/dca"

	"github.com/rylio/ytdl"
)

type Player struct {
	id     string
	stream *dca.StreamingSession
}

type Queue struct {
	URL      string
	Name     string
	Duration string
}

type Queues struct {
	GuildID string
	Queue   []Queue
}

type Parameters struct {
	InVoice bool
	vc      *discordgo.VoiceConnection
}

var AllQueues []Queues
var players []Player

func RemoveFromArray(s []Player, GuildID string) []Player {
	i := 0
	exist := false
	for i = 0; i < len(players); i++ {
		if players[i].id == GuildID {
			exist = true
			break
		}
	}
	if exist {
		s[i] = s[len(s)-1]
		return s[:len(s)-1]
	} else {
		return s
	}
}

func RemoveFromQueue(s []Queues, GuildID string, URL string) []Queue {
	i := 0
	k := 0
	exist := false
	for i = 0; i < len(s); i++ {
		if s[i].GuildID == GuildID {
			for k = 0; k < len(s[i].Queue); k++ {
				if s[i].Queue[k].URL == URL {
					exist = true
					break
				}
			}
			break
		}
	}
	if exist {
		s[i].Queue[k] = s[i].Queue[len(s[i].Queue)-1]
		return s[i].Queue[:len(s[i].Queue)-1]
	} else {
		return s[i].Queue
	}
}

func RemoveQueue(s []Queues, GuildID string) []Queues {
	i := 0
	exist := false
	for i = 0; i < len(s); i++ {
		if s[i].GuildID == GuildID {
			exist = true
			break
		}
	}
	if exist {
		s[i] = s[len(s)-1]
		return s[:len(s)-1]
	} else {
		return s
	}
}

func Playing(s *discordgo.Session, GuildID string, MessageChannelID string, ChannelID string, videoID string, InVoice Parameters) {
	options := dca.StdEncodeOptions
	options.RawOutput = true
	options.Bitrate = 96
	options.BufferedFrames = 5000
	options.Application = "lowdelay"
	options.PacketLoss = 1

	videoInfo, err := ytdl.GetVideoInfo(`https://www.youtube.com/watch?v=` + videoID)
	if err != nil {
		// Handle the error
	}

	s.ChannelMessageSend(MessageChannelID, ":arrow_down: Downloading : "+videoInfo.Title)

	format := videoInfo.Formats.Extremes(ytdl.FormatAudioBitrateKey, true)[0]
	downloadURL, err := videoInfo.GetDownloadURL(format)
	if err != nil {
		// Handle the error
	}

	resp, err := http.Get(downloadURL.String())
	if err != nil {
		return
	}
	encodingSession, err := dca.EncodeMem(resp.Body, options)
	if err != nil {
		// Handle the error
	}
	defer encodingSession.Cleanup()

	s.ChannelMessageSend(MessageChannelID, ":musical_note: Started playing : "+videoInfo.Title)
	vc := &discordgo.VoiceConnection{}
	if InVoice.InVoice {
		vc = InVoice.vc
	} else {
		vc, err = s.ChannelVoiceJoin(GuildID, ChannelID, false, false)
		if err != nil {
			return
		}
	}
	_ = vc.Speaking(true)
	done := make(chan error)
	stream := dca.NewStream(encodingSession, vc, done)
	stream.SetPaused(false)
	CurrentPlayer := Player{
		id:     GuildID,
		stream: stream,
	}
	players = append(players, CurrentPlayer)
	err = <-done
	if err != nil && err != io.EOF {
		//
	}
	time.Sleep(250 * time.Millisecond)
	players = RemoveFromArray(players, GuildID)
	QueueCheck(s, GuildID, MessageChannelID, ChannelID, videoID, vc)
	return
}

func QueueCheck(s *discordgo.Session, GuildID string, MessageChannelID string, ChannelID string, videoID string, vc *discordgo.VoiceConnection) {
	i := 0
	k := 0
	exist := false
	for i = 0; i < len(AllQueues); i++ {
		if AllQueues[i].GuildID == GuildID {
			for k = 0; k < len(AllQueues[i].Queue); k++ {
				if AllQueues[i].Queue[k].URL == videoID {
					exist = true
					break
				}
			}
			break
		}
	}
	if exist {
		AllQueues[i].Queue = append(AllQueues[i].Queue[:k], AllQueues[i].Queue[k+1:]...)
		if len(AllQueues[i].Queue) != 0 {
			Params := Parameters{
				InVoice: true,
				vc:      vc,
			}
			go Playing(s, GuildID, MessageChannelID, ChannelID, AllQueues[i].Queue[0].URL, Params)
		} else {
			_ = vc.Disconnect()
		}
	} else {
		_ = vc.Disconnect()
	}
}

func QueueManagement(s *discordgo.Session, GuildID string, MessageChannelID string, ChannelID string, videoID string) {
	i := 0
	exist := false
	for i = 0; i < len(AllQueues); i++ {
		if AllQueues[i].GuildID == GuildID {
			exist = true
			break
		}
	}
	if exist {
		Song := Queue{
			URL: videoID,
		}
		videoInfo, err := ytdl.GetVideoInfo(`https://www.youtube.com/watch?v=` + videoID)
		if err != nil {
			// Handle the error
		}
		Song.Name = videoInfo.Title
		Song.Duration = videoInfo.Duration.String()
		AllQueues[i].Queue = append(AllQueues[i].Queue, Song)
	} else {
		Song := Queue{
			URL: videoID,
		}
		videoInfo, err := ytdl.GetVideoInfo(`https://www.youtube.com/watch?v=` + videoID)
		if err != nil {
			// Handle the error
		}
		Song.Name = videoInfo.Title
		Song.Duration = videoInfo.Duration.String()
		var songs []Queue
		songs = append(songs, Song)
		Queue := Queues{
			GuildID: GuildID,
			Queue:   songs,
		}
		AllQueues = append(AllQueues, Queue)
	}
	Params := Parameters{
		InVoice: false,
	}
	Playing(s, GuildID, MessageChannelID, ChannelID, videoID, Params)
}

func AddToQueue(Session *discordgo.Session, GuildID string, MessageChannelID string, ChannelID string, URL string) {
	NotId := strings.Split(URL, "&")
	NotId = strings.Split(NotId[0], "=")
	videoID := NotId[1] // Converts the URL to only an ID
	i := 0
	exist := false
	for i = 0; i < len(AllQueues); i++ { // Will search for the guild and see if there's already an existing queue
		if AllQueues[i].GuildID == GuildID {
			exist = true
			break
		}
	}
	if exist {
		Song := Queue{
			URL: videoID,
		}
		videoInfo, err := ytdl.GetVideoInfo(`https://www.youtube.com/watch?v=` + videoID)
		if err != nil {
			// Handle the error
		}
		Song.Name = videoInfo.Title
		Song.Duration = videoInfo.Duration.String()
		AllQueues[i].Queue = append(AllQueues[i].Queue, Song)

		Session.ChannelMessageSend(MessageChannelID, ":arrow_down: Adding to queue : "+videoInfo.Title)
	} else {
		Song := Queue{
			URL: videoID,
		}
		videoInfo, err := ytdl.GetVideoInfo(`https://www.youtube.com/watch?v=` + videoID)
		if err != nil {
			// Handle the error
		}
		Song.Name = videoInfo.Title
		Song.Duration = videoInfo.Duration.String()
		var songs []Queue
		songs = append(songs, Song)
		Queue := Queues{
			GuildID: GuildID,
			Queue:   songs,
		}
		AllQueues = append(AllQueues, Queue)
		Params := Parameters{
			InVoice: false,
		}
		go Playing(Session, GuildID, MessageChannelID, ChannelID, videoID, Params)
	}
}

func PlayYoutubeLink(Session *discordgo.Session, GuildID string, MessageChannelID string, ChannelID string, URL string) {
	i := 0
	for i = 0; i < len(players); i++ {
		if players[i].id == GuildID {
			AddToQueue(Session, GuildID, MessageChannelID, ChannelID, URL)
			return
		}
	}
	NotId := strings.Split(URL, "&")
	NotId = strings.Split(NotId[0], "=")
	videoID := NotId[1]
	go QueueManagement(Session, GuildID, MessageChannelID, ChannelID, videoID)
}

func StopPlaying(GuildID string, ChannelID string, Session *discordgo.Session) {
	AllQueues = RemoveQueue(AllQueues, GuildID)
	i := 0
	exist := false
	for i = 0; i < len(players); i++ {
		if players[i].id == GuildID {
			exist = true
			break
		}
	}
	if exist {
		Session.ChannelMessageSend(ChannelID, ":no_entry: Removed all songs from the queue and stopped playing.")
		players[i].stream.Stop()
	}
}

func SkipPlaying(GuildID string, ChannelID string, Session *discordgo.Session) {
	i := 0
	exist := false
	for i = 0; i < len(players); i++ {
		if players[i].id == GuildID {
			exist = true
			break
		}
	}
	if exist {
		Session.ChannelMessageSend(ChannelID, ":arrow_forward: Skipped to the next song.")
		players[i].stream.Stop()
	}
}

func PausePlaying(GuildID string, ChannelID string, Session *discordgo.Session) {
	i := 0
	exist := false
	for i = 0; i < len(players); i++ {
		if players[i].id == GuildID {
			exist = true
			break
		}
	}
	if exist {
		Session.ChannelMessageSend(ChannelID, ":pause_button: Paused the current song.")
		players[i].stream.SetPaused(true)
	}
}

func ResumePlaying(GuildID string, ChannelID string, Session *discordgo.Session) {
	i := 0
	exist := false
	for i = 0; i < len(players); i++ {
		if players[i].id == GuildID {
			exist = true
			break
		}
	}
	if exist {
		Session.ChannelMessageSend(ChannelID, ":musical_note: Resuming the current song.")
		players[i].stream.SetPaused(false)
	}
}
