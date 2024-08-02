package entities

import (
	"Brightwells/components"

	"github.com/hajimehoshi/ebiten/v2"
)

// Id = 1
func NewPlayer(posX, posY float64, sprite *ebiten.Image, layer int) *Entity {
	// Create base entity
	entity := NewEntity(layer)

	// Mark as player
	entity.AddComponent(components.PlayerComponentID, &components.PlayerComponent{})

	// Give state
	entity.AddComponent(components.StateComponentID, &components.StateComponent{
		CurrentState: 0,
		NextState:    0,
	})
	// Have Camera set to player
	entity.AddComponent(components.CameraComponentID, &components.CameraComponent{
		X: posX,
		Y: posY,
	})

	// Physics components
	entity.AddComponent(components.PositionComponentID, &components.PositionComponent{
		TileX: posX,
		TileY: posY,
		DesX:  posX,
		DesY:  posY,
	})

	entity.AddComponent(components.VelocityComponentID, &components.VelocityComponent{VX: 0, VY: 0})
	entity.AddComponent(components.CollisionComponentID, &components.CollisionComponent{})
	entity.AddComponent(components.CollisionBoxID, &components.CollisionBox{})

	// Spawn Point
	entity.AddComponent(components.SpawnPointComponentID, &components.SpawnPointComponent{
		TileX: 0,
		TileY: 0,
	})
	// Visual compontents
	entity.AddComponent(components.SpriteComponentID, &components.SpriteComponent{
		Image: sprite,
		X:     0, Y: 0, X1: 16, Y1: 16, // Initial sub-image coordinates
	})

	// Combat components
	entity.AddComponent(components.AttackerComponentID, &components.AttackerComponent{
		Target: -1,
	})

	// Skill compontents
	entity.AddComponent(components.DamageComponentID, &components.DamageComponent{})
	entity.AddComponent(components.SkillsComponentID, &components.SkillsComponent{
		CurrentHealth: 15,
		Health:        15,
		Melee:         5,
		Defense:       5,
		Devotion:      5,
	})

	// Resource components
	entity.AddComponent(components.GatherComponentID, &components.GatherComponent{Target: -1})

	return entity

}
