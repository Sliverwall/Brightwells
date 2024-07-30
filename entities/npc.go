package entities

import (
	"Brightwells/components"

	"github.com/hajimehoshi/ebiten/v2"
)

// Id = 2
func NewMonsterGirl(posX, posY float64, sprite *ebiten.Image, layer int) *Entity {
	entity := NewEntity(layer) // Ensure you have a NewEntity() function that initializes a new Entity.

	// Spatial compontents
	entity.AddComponent(components.PositionComponentID, &components.PositionComponent{
		TileX: posX,
		TileY: posY,
	})
	entity.AddComponent(components.VelocityComponentID, &components.VelocityComponent{
		VX: 0,
		VY: 1.0,
	})

	// Visual components
	entity.AddComponent(components.SpriteComponentID, &components.SpriteComponent{
		Image: sprite,
		X:     0, Y: 0, X1: 16, Y1: 16, // Initial sub-image coordinates
	})

	// Collision compontents
	entity.AddComponent(components.CollisionComponentID, &components.CollisionComponent{})
	entity.AddComponent(components.CollisionBoxID, &components.CollisionBox{})

	// Combat compontents
	entity.AddComponent(components.AttackerComponentID, &components.AttackerComponent{
		IsAttacking: false,
		Target:      -1,
	})
	entity.AddComponent(components.DamageComponentID, &components.DamageComponent{})
	entity.AddComponent(components.SkillsComponentID, &components.SkillsComponent{
		CurrentHealth: 15,
		Health:        15,
		Melee:         5,
		Defense:       5,
		Devotion:      5,
	})
	return entity
}
