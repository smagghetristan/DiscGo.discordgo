package commands

import (
	"math/rand"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

//Will calculate the score of a square
func ScoreCalc(Field [10][10]string, i int, j int) int {
	Score := 0
	for k := i - 1; k < i+2; k++ {
		for l := j - 1; l < j+2; l++ {
			//Prevents any glitch on the corners and on the firsts and lasts columns / lines
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

//Function is called when the user types the command
func MineSweeperCreate(session *discordgo.Session, m *discordgo.MessageCreate) {
	//Will be used to convert an int to a text for discord
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
	//Generate the bombs, 8/100 chances
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
	//Calculates the scores of each squares and stores them
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if Field[i][j] == "." {
				Field[i][j] = strconv.Itoa(ScoreCalc(Field, i, j))
			}
		}
	}
	Message := ""
	//Prepare the message and sends it afterwhile
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
