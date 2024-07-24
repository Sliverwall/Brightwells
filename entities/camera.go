package entities

import (
	"Brightwells/components"
	"Brightwells/config"

	"github.com/hajimehoshi/ebiten/v2"
)

func NewCamera() *Entity {
	entity := NewEntity(-1) // Ensure you have a NewEntity() function that initializes a new Entity.

	// Spatial compontents
	entity.AddComponent(components.CameraComponentID, &components.CameraComponent{
		Width:    float64(config.RESOLUTION_WIDTH),
		Height:   float64(config.RESOLUTION_HEIGHT),
		Viewport: ebiten.NewImage(config.RESOLUTION_WIDTH, config.RESOLUTION_HEIGHT),
	})

	return entity
}
