package entities

import (
	"Brightwells/components"
	"Brightwells/config"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

// Id = 1
func NewPlayer(posX, posY float64, sprite *ebiten.Image, layer int) *Entity {
	// Create base entity
	entity := NewEntity(layer)

	// Mark as player
	entity.AddComponent(components.PlayerComponentID, &components.PlayerComponent{})

	// Have Camera set to player
	entity.AddComponent(components.CameraComponentID, &components.CameraComponent{
		X: posX,
		Y: posY,
	})

	// Physics components
	entity.AddComponent(components.PositionComponentID, &components.PositionComponent{
		X:     math.Floor(posX * config.TileSize),
		Y:     math.Floor(posY * config.TileSize),
		TileX: posX,
		TileY: posY,
	})
	entity.AddComponent(components.VelocityComponentID, &components.VelocityComponent{VX: 0, VY: 0})
	entity.AddComponent(components.CollisionComponentID, &components.CollisionComponent{})
	entity.AddComponent(components.CollisionBoxID, &components.CollisionBox{})

	// Visual compontents
	entity.AddComponent(components.SpriteComponentID, &components.SpriteComponent{
		Image: sprite,
		X:     0, Y: 0, X1: 16, Y1: 16, // Initial sub-image coordinates
	})

	// Combat components
	entity.AddComponent(components.AttackerComponentID, &components.AttackerComponent{
		IsAttacking: false,
		Target:      -1,
	})
	entity.AddComponent(components.DestinationComponentID, &components.DestinationComponent{X: posX, Y: posY})

	// Skill compontents
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
