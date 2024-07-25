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

// RenderComponentID is the identifier for the RenderComponent
const RenderComponentID = "RenderComponent"

type RenderComponent struct {
	Layer int // Higher layers are drawn on top of lower layers
}

// ShapeComponentID is the identifier for the ShapeComponent
const ShapeComponentID = "ShapeComponent"
const ShapeColorComponentID = "ShapeColorComponentID"

type ShapeComponent struct {
	Height, Width float64
	Geometry      string
}

type ShapeColorComponent struct {
	// RGBA units for shape's color
	R, G, B, A uint8
}
