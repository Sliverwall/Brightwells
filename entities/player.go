package entities

import (
	"Brightwells/components"

	"github.com/hajimehoshi/ebiten/v2"
)

func NewPlayer(posX, posY, velX, velY float64, sprite *ebiten.Image) *Entity {
	entity := NewEntity() // Ensure you have a NewEntity() function that initializes a new Entity.

	entity.AddComponent(components.PlayerComponentID, &components.PlayerComponent{})
	entity.AddComponent(components.PositionComponentID, &components.PositionComponent{X: posX, Y: posY})
	entity.AddComponent(components.VelocityComponentID, &components.VelocityComponent{VX: velX, VY: velY})
	entity.AddComponent(components.SpriteComponentID, &components.SpriteComponent{
		Image: sprite,
		X:     0, Y: 0, X1: 16, Y1: 16, // Initial sub-image coordinates
	})
	return entity
}
