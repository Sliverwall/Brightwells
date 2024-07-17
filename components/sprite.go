package components

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// SpriteComponentID is the identifier for the SpriteComponent
const SpriteComponentID = "SpriteComponent"

// SpriteComponent represents a component that holds the sprite image and its sub-image coordinates.
type SpriteComponent struct {
	Image        *ebiten.Image
	X, Y, X1, Y1 int // Coordinates for the sub-image within the spritesheet
}
