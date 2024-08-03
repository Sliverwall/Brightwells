package systems

import (
	"Brightwells/data"
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

func (ls *LoginSystem) Update() bool {
	// Capture keyboard input to form the username string
	ls.handleInput()

	// Check if Enter is pressed to submit the form
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		// Load Login table
		playerData := data.SQL_query(data.Select_all_Player)

		for _, player := range playerData {
			if player[2] == ls.username {
				return true
			}
		}
	}

	return false
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
