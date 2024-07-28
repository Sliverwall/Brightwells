package entities

import (
	"Brightwells/components"
	"Brightwells/config"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func NewApple(posX, posY float64, sprite *ebiten.Image, layer int) *Entity {
	entity := NewEntity(layer) // Ensure you have a NewEntity() function that initializes a new Entity.

	entity.AddComponent(components.FoodComponentID, &components.FoodComponent{})
	entity.AddComponent(components.PositionComponentID, &components.PositionComponent{
		X:     math.Floor(posX * config.TileSize),
		Y:     math.Floor(posY * config.TileSize),
		TileX: posX,
		TileY: posY,
	})
	entity.AddComponent(components.SpriteComponentID, &components.SpriteComponent{
		Image: sprite,
		X:     0, Y: 0, X1: 16, Y1: 16, // Initial sub-image coordinates
	})
	entity.AddComponent(components.CollisionComponentID, &components.CollisionComponent{})

	return entity
}
