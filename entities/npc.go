package entities

import (
	"Brightwells/components"
	"Brightwells/config"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func NewNPC(posX, posY, velX, velY float64, sprite *ebiten.Image, layer int) *Entity {
	entity := NewEntity(layer) // Ensure you have a NewEntity() function that initializes a new Entity.

	entity.AddComponent(components.PositionComponentID, &components.PositionComponent{X: posX, Y: posY})
	entity.AddComponent(components.PositionComponentID, &components.PositionComponent{
		X:     posX,
		Y:     posY,
		TileX: math.Round(posX / config.TileSize),
		TileY: math.Round(posY / config.TileSize),
	})
	entity.AddComponent(components.SpriteComponentID, &components.SpriteComponent{
		Image: sprite,
		X:     0, Y: 0, X1: 16, Y1: 16, // Initial sub-image coordinates
	})
	entity.AddComponent(components.CollisionComponentID, &components.CollisionComponent{
		Width:  config.TileSize / 2,
		Height: config.TileSize / 2,
	})
	entity.AddComponent(components.CollisionBoxID, &components.CollisionBox{
		PositionComponent:  entity.GetComponent(components.PositionComponentID).(*components.PositionComponent),
		CollisionComponent: entity.GetComponent(components.CollisionComponentID).(*components.CollisionComponent),
	})
	return entity
}
