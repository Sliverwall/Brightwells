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

type ShapeComponent struct {
	Height, Width float64
	Geometry      string
}

const ShapeColorComponentID = "ShapeColorComponentID"

type ShapeColorComponent struct {
	// RGBA units for shape's color
	R, G, B, A uint8
}

// enum for right click options
const (
	ClickWalkHere = iota
	ClickAttack
	ClickGather
)
const RightClickComponentID = "RightClickComponentID"

type RightClickComponent struct {
	Actions []int // Actions like "Walk here", "Attack", "Gather"
}
