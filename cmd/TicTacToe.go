package commands

import (
	"strconv"
	"strings"
	"time"

	"DiscGo.discordgo/config"
	"github.com/bwmarrin/discordgo"
)

type tttGrid struct {
	Line1   []int
	Line2   []int
	Line3   []int
	Message string
}

type tttGame struct {
	CreatedAt   time.Time
	ChannelID   string
	Player1     string
	Player2     string
	Turn        int
	Grid        tttGrid
	Started     bool
	Win         bool
	WinID       string
	TurnMessage *discordgo.Message
	GridMessage *discordgo.Message
}

var AllTTTGames []tttGame

func RemoveTTTGame(s []tttGame, Player1 string, Player2 string) []tttGame {
	i := 0
	exist := false
	for i = 0; i < len(s); i++ {
		if s[i].Player1 == Player1 && s[i].Player2 == Player2 {
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

func Timer(s *discordgo.Session, ChannelID string, Player1 string) {
	time.Sleep(20 * time.Second)
	i := 0
	exist := false
	for i = 0; i < len(AllTTTGames); i++ {
		if AllTTTGames[i].ChannelID == ChannelID && AllTTTGames[i].Player1 == Player1 {
			exist = true
			break
		}
	}
	if exist && AllTTTGames[i].Player2 == "" {
		AllTTTGames[i] = AllTTTGames[len(AllTTTGames)-1]
		AllTTTGames = AllTTTGames[:len(AllTTTGames)-1]
		s.ChannelMessageSend(ChannelID, "Too slow, the game has been deleted.")
	} else {
		return
	}
}

func GameTimer(s *discordgo.Session, ChannelID string, Game tttGame) {
	time.Sleep(5 * time.Minute)
	i := 0
	exist := false
	for i = 0; i < len(AllTTTGames); i++ {
		if AllTTTGames[i].Player1 == Game.Player1 && AllTTTGames[i].Player2 == Game.Player2 {
			exist = true
			break
		}
	}
	if exist {
		AllTTTGames[i] = AllTTTGames[len(AllTTTGames)-1]
		AllTTTGames = AllTTTGames[:len(AllTTTGames)-1]
		s.ChannelMessageSend(ChannelID, "Too slow, the game has been deleted.")
	} else {
		return
	}
}

func CheckForWin(Grid tttGrid, number int) (bool, bool) {
	if Grid.Line1[0] == number && Grid.Line1[1] == number && Grid.Line1[2] == number {
		return true, false
	}
	if Grid.Line2[0] == number && Grid.Line2[1] == number && Grid.Line2[2] == number {
		return true, false
	}
	if Grid.Line3[0] == number && Grid.Line3[1] == number && Grid.Line3[2] == number {
		return true, false
	}

	if Grid.Line1[0] == number && Grid.Line2[0] == number && Grid.Line3[0] == number {
		return true, false
	}
	if Grid.Line1[1] == number && Grid.Line2[1] == number && Grid.Line3[1] == number {
		return true, false
	}
	if Grid.Line1[2] == number && Grid.Line2[2] == number && Grid.Line3[2] == number {
		return true, false
	}

	if Grid.Line1[0] == number && Grid.Line2[1] == number && Grid.Line3[2] == number {
		return true, false
	}
	if Grid.Line1[2] == number && Grid.Line2[1] == number && Grid.Line3[0] == number {
		return true, false
	}
	if CheckForTie(Grid) {
		return false, true
	}
	return false, false
}

func CheckForTie(Grid tttGrid) bool {
	for i := 0; i <= 3; i++ {
		if Grid.Line1[i] == 0 {
			return false
		}
		if Grid.Line2[i] == 0 {
			return false
		}
		if Grid.Line3[i] == 0 {
			return false
		}
	}
	return true
}

func TCommandStart(s *discordgo.Session, m *discordgo.MessageCreate) {
	AvatarURL := m.Author.AvatarURL("512")
	embed := &discordgo.MessageEmbed{
		Title:       "Tic Tac Toe",
		Description: "To start playing you need a second player, he can join by typing g!t p2",
		Footer: &discordgo.MessageEmbedFooter{
			IconURL: AvatarURL,
			Text:    m.Author.Username,
		},
		Color: 0xFFDD00,
	}
	_, err := s.ChannelMessageSendEmbed(m.ChannelID, embed)
	if err != nil {
		return
	}
	err = s.ChannelMessageDelete(m.ChannelID, m.ID)
	if err != nil {
		return
	}
	NewGame := tttGame{
		CreatedAt: time.Now(),
		Player1:   m.Author.ID,
		ChannelID: m.ChannelID,
		Win:       false,
		Started:   false,
	}
	AllTTTGames = append(AllTTTGames, NewGame)
	go Timer(s, m.ChannelID, m.Author.ID)
}

func TCommandP2(s *discordgo.Session, m *discordgo.MessageCreate) {
	i := 0
	exist := false
	for i = 0; i < len(AllTTTGames); i++ {
		if AllTTTGames[i].ChannelID == m.ChannelID && !AllTTTGames[i].Started && AllTTTGames[i].Player1 != m.Author.ID {
			exist = true
			break
		}
	}
	if exist && AllTTTGames[i].Player1 == m.Author.ID {
		s.ChannelMessageSend(m.ChannelID, ":x: You cannot play against yourself.")
		return
	} else if exist && AllTTTGames[i].Started == true {
		s.ChannelMessageSend(m.ChannelID, ":x: The game has already started.")
		return
	} else if exist {
		AllTTTGames[i].Player2 = m.Author.ID
		AllTTTGames[i].Started = true
		AllTTTGames[i].Turn = 1
		AllTTTGames[i].Grid = tttGrid{
			Line1: []int{0, 0, 0},
			Line2: []int{0, 0, 0},
			Line3: []int{0, 0, 0},
			Message: ":one: :two: :three:\n" +
				":four: :five: :six:\n" +
				":seven: :eight: :nine:\n",
		}
		AllTTTGames[i].TurnMessage, _ = s.ChannelMessageSend(m.ChannelID, "It is **Player 1**'s turn !")
		AllTTTGames[i].GridMessage, _ = s.ChannelMessageSend(m.ChannelID, AllTTTGames[i].Grid.Message)
		go GameTimer(s, m.ChannelID, AllTTTGames[i])
		return
	}
}

func TCommandGame(s *discordgo.Session, m *discordgo.MessageCreate) {
	message := strings.Replace(m.Content, config.Prefix+"t ", "", 1)
	command := strings.Split(message, " ")[0]
	i := 0
	exist := false
	for i = 0; i < len(AllTTTGames); i++ {
		if (AllTTTGames[i].Player1 == m.Author.ID || AllTTTGames[i].Player2 == m.Author.ID) && AllTTTGames[i].Started {
			exist = true
			break
		}
	}
	if exist {
		NumString := []string{"", ":one:", ":two:", ":three:", ":four:", ":five:", ":six:", ":seven:", ":eight:", ":nine:"}
		number, _ := strconv.Atoi(command)
		if number > 0 && number < 10 {
			if AllTTTGames[i].Turn == 1 && AllTTTGames[i].Player1 == m.Author.ID {
				UpdateGrid := false
				if number > 6 {
					number -= 7
					if AllTTTGames[i].Grid.Line1[number] == 0 {
						UpdateGrid = true
						AllTTTGames[i].Grid.Line1[number] = 1
					}
				} else if number > 3 {
					number -= 4
					if AllTTTGames[i].Grid.Line2[number] == 0 {
						UpdateGrid = true
						AllTTTGames[i].Grid.Line2[number] = 1
					}
				} else {
					number -= 1
					if AllTTTGames[i].Grid.Line3[number] == 0 {
						UpdateGrid = true
						AllTTTGames[i].Grid.Line3[number] = 1
					}
				}
				if UpdateGrid {
					Win, Tie := CheckForWin(AllTTTGames[i].Grid, 1)
					number, _ := strconv.Atoi(command)
					AllTTTGames[i].Turn = 2
					NewContent := strings.Replace(AllTTTGames[i].TurnMessage.Content, "1", "2", 1)
					Edit := &discordgo.MessageEdit{
						Content: &NewContent,
						ID:      AllTTTGames[i].TurnMessage.ID,
						Channel: AllTTTGames[i].TurnMessage.ChannelID,
					}
					s.ChannelMessageEditComplex(Edit)

					NewContent = strings.Replace(AllTTTGames[i].GridMessage.Content, NumString[number], ":x:", 1)
					Edit = &discordgo.MessageEdit{
						Content: &NewContent,
						ID:      AllTTTGames[i].GridMessage.ID,
						Channel: AllTTTGames[i].GridMessage.ChannelID,
					}
					AllTTTGames[i].GridMessage, _  = s.ChannelMessageEditComplex(Edit)
					if Win {
						RemoveTTTGame(AllTTTGames, AllTTTGames[i].Player1, AllTTTGames[i].Player2)
						s.ChannelMessageSend(AllTTTGames[i].GridMessage.ChannelID, "**Player 1** won the game !")
					}
					if Tie {
						RemoveTTTGame(AllTTTGames, AllTTTGames[i].Player1, AllTTTGames[i].Player2)
						s.ChannelMessageSend(AllTTTGames[i].GridMessage.ChannelID, "No one won the game !")
					}
					err := s.ChannelMessageDelete(m.ChannelID, m.ID)
					if err != nil {
						return
					}
				}
			} else if AllTTTGames[i].Turn == 2 && AllTTTGames[i].Player2 == m.Author.ID {
				UpdateGrid := false
				if number > 6 {
					number -= 7
					if AllTTTGames[i].Grid.Line1[number] == 0 {
						UpdateGrid = true
						AllTTTGames[i].Grid.Line1[number] = 2
					}
				} else if number > 3 {
					number -= 4
					if AllTTTGames[i].Grid.Line2[number] == 0 {
						UpdateGrid = true
						AllTTTGames[i].Grid.Line2[number] = 2
					}
				} else {
					number -= 1
					if AllTTTGames[i].Grid.Line3[number] == 0 {
						UpdateGrid = true
						AllTTTGames[i].Grid.Line3[number] = 2
					}
				}
				if UpdateGrid {
					Win, Tie := CheckForWin(AllTTTGames[i].Grid, 2)
					number, _ := strconv.Atoi(command)
					AllTTTGames[i].Turn = 1
					NewContent := strings.Replace(AllTTTGames[i].TurnMessage.Content, "2", "1", 1)
					Edit := &discordgo.MessageEdit{
						Content: &NewContent,
						ID:      AllTTTGames[i].TurnMessage.ID,
						Channel: AllTTTGames[i].TurnMessage.ChannelID,
					}
					s.ChannelMessageEditComplex(Edit)

					NewContent = strings.Replace(AllTTTGames[i].GridMessage.Content, NumString[number], ":o:", 1)
					Edit = &discordgo.MessageEdit{
						Content: &NewContent,
						ID:      AllTTTGames[i].GridMessage.ID,
						Channel: AllTTTGames[i].GridMessage.ChannelID,
					}
					AllTTTGames[i].GridMessage, _ = s.ChannelMessageEditComplex(Edit)
					if Win {
						RemoveTTTGame(AllTTTGames, AllTTTGames[i].Player1, AllTTTGames[i].Player2)
						s.ChannelMessageSend(AllTTTGames[i].GridMessage.ChannelID, "**Player 2** won the game !")
					}
					if Tie {
						RemoveTTTGame(AllTTTGames, AllTTTGames[i].Player1, AllTTTGames[i].Player2)
						s.ChannelMessageSend(AllTTTGames[i].GridMessage.ChannelID, "No one won the game !")
					}
					err := s.ChannelMessageDelete(m.ChannelID, m.ID)
					if err != nil {
						return
					}
				}
			}
		}
	}
}
