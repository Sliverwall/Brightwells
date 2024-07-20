package entities

import (
	"Brightwells/components"

	"github.com/hajimehoshi/ebiten/v2"
)

func NewApple(posX, posY float64, sprite *ebiten.Image, layer int) *Entity {
	entity := NewEntity(layer) // Ensure you have a NewEntity() function that initializes a new Entity.

	entity.AddComponent(components.FoodComponentID, &components.FoodComponent{})
	entity.AddComponent(components.PositionComponentID, &components.PositionComponent{X: posX, Y: posY})
	entity.AddComponent(components.SpriteComponentID, &components.SpriteComponent{
		Image: sprite,
		X:     0, Y: 0, X1: 16, Y1: 16, // Initial sub-image coordinates
	})
	entity.AddComponent(components.CollisionComponentID, &components.CollisionComponent{
		Width:  16,
		Height: 16,
	})
	entity.AddComponent(components.CollisionBoxID, &components.CollisionBox{
		PositionComponent:  entity.GetComponent(components.PositionComponentID).(*components.PositionComponent),
		CollisionComponent: entity.GetComponent(components.CollisionComponentID).(*components.CollisionComponent),
	})

	return entity
}
