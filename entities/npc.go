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
		DesX:  posX,
		DesY:  posY,
	})
	// Spawn Point
	entity.AddComponent(components.SpawnPointComponentID, &components.SpawnPointComponent{
		TileX:            posX,
		TileY:            posY,
		RespawnTime:      10,
		RespawnTimeCount: 0,
	})

	// Idle Point
	entity.AddComponent(components.IdlePositionComponentID, &components.IdlePositionComponent{
		TileX: posX,
		TileY: posY,
	})
	entity.AddComponent(components.VelocityComponentID, &components.VelocityComponent{
		VX: 0,
		VY: 0,
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

	// Give state
	entity.AddComponent(components.StateComponentID, &components.StateComponent{
		CurrentState: 0,
		NextState:    1,
	})

	// Set attack response to attack back
	entity.AddComponent(components.AttackedResponseComponentID, &components.AttackedResponseComponent{
		Type: components.AttackBack,
	})

	entity.AddComponent(components.AttackerComponentID, &components.AttackerComponent{
		Target: -1,
	})

	entity.AddComponent(components.DamageComponentID, &components.DamageComponent{})
	entity.AddComponent(components.SkillsComponentID, &components.SkillsComponent{
		CurrentHealth: 5,
		Health:        5,
		Melee:         5,
		Defense:       5,
		Devotion:      5,
	})
	return entity
}
