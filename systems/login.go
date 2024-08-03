package systems

import (
	"Brightwells/data"
	"fmt"
	"log"
	"unicode"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type LoginSystem struct {
	username string
}

func NewLoginSystem() *LoginSystem {
	return &LoginSystem{}
}

func (ls *LoginSystem) Update() (string, bool) {
	// Capture keyboard input to form the username string
	ls.handleInput()

	// Check if Enter is pressed to submit the form
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		// Load Login table
		playerData := data.SQL_query(data.Select_all_Player)

		for _, player := range playerData {
			if player[2] == ls.username {
				// Take player data then map onto entity table
				log.Println(player[2], player[3], player[4])
				x := float64(player[3].(int64))
				y := float64(player[4].(int64))
				log.Println(x, ",", y)
				updateCurrentPlayerData := fmt.Sprintf("UPDATE ENTITY SET x = %f, y = %f WHERE name = 'player'", x, y)
				data.SQL_exec(updateCurrentPlayerData)
				return ls.username, true
			}
		}
	}

	return ls.username, false
}

func (ls *LoginSystem) handleInput() {
	// Check for backspace to delete characters
	if ebiten.IsKeyPressed(ebiten.KeyBackspace) && len(ls.username) > 0 {
		ls.username = ls.username[:len(ls.username)-1]
	}

	// Append typed characters to the username
	for _, char := range ebiten.InputChars() {
		// Filter only printable characters
		if unicode.IsPrint(char) {
			ls.username += string(char)
		}
	}
}

func (ls *LoginSystem) Draw(screen *ebiten.Image) {
	// Draw login prompt and the current username
	text := "Login Screen\nEnter Username: " + ls.username + "\nPress Enter to submit"
	ebitenutil.DebugPrint(screen, text)
}
