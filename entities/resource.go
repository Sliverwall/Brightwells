package entities

import (
	"Brightwells/components"

	"github.com/hajimehoshi/ebiten/v2"
)

// id = 3
func NewTree(posX, posY float64, sprite *ebiten.Image, layer int) *Entity {
	entity := NewEntity(layer) // Ensure you have a NewEntity() function that initializes a new Entity.

	entity.AddComponent(components.CollisionComponentID, &components.CollisionComponent{})
	entity.AddComponent(components.CollisionBoxID, &components.CollisionBox{})
	// Physics components
	entity.AddComponent(components.PositionComponentID, &components.PositionComponent{
		TileX: posX,
		TileY: posY,
	})
	entity.AddComponent(components.SpriteComponentID, &components.SpriteComponent{
		Image: sprite,
		X:     0, Y: 0, X1: 32, Y1: 32, // Initial sub-image coordinates
	})
	entity.AddComponent(components.CollisionComponentID, &components.CollisionComponent{})
	entity.AddComponent(components.ResourceNodeComponentID, &components.ResourceNodeComponent{
		Type:             "Lumbering",
		SkillsNeeded:     map[string]int{"Lumbering": 1},
		Active:           true,
		DrainedChance:    0.10,
		RespawnTime:      10,
		RespawnTimeCount: 0,
	})

	return entity
}
