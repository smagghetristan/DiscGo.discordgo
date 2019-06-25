package commands

import (
	"math/rand"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

func ScoreCalc(Field [10][10]string, i int, j int) int {
	Score := 0
	for k := i - 1; k < i+2; k++ {
		for l := j - 1; l < j+2; l++ {
			if k >= 0 && l >= 0 {
				if k <= 9 && l <= 9 {
					if Field[k][l] == "x" {
						Score++
					}
				}
			}
		}
	}
	return Score
}

func MineSweeperCreate(session *discordgo.Session, m *discordgo.MessageCreate) {
	Numbers := []string{
		"||:zero:||",
		"||:one:||",
		"||:two:||",
		"||:three:||",
		"||:four:||",
		"||:five:||",
		"||:six:||",
		"||:seven:||",
		"||:eight:||",
	}
	var Field [10][10]string
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			dice := rand.Intn(100)
			if dice < 8 {
				Field[i][j] = "x"
			} else {
				Field[i][j] = "."
			}
		}
	}
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if Field[i][j] == "." {
				Field[i][j] = strconv.Itoa(ScoreCalc(Field, i, j))
			}
		}
	}
	Message := ""
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if Field[i][j] != "x" {
				Doing, _ := strconv.Atoi(Field[i][j])
				Message += Numbers[Doing]
			} else {
				Message += "||:bomb:||"
			}
		}
		Message += "\n"
	}
	session.ChannelMessageSend(m.ChannelID, "Here is your minesweeper grid :")
	session.ChannelMessageSend(m.ChannelID, Message)
}
