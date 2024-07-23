package entities

import (
	"Brightwells/components"
	"Brightwells/config"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func NewPlayer(posX, posY, velX, velY float64, sprite *ebiten.Image, layer int) *Entity {
	entity := NewEntity(layer) // Ensure you have a NewEntity() function that initializes a new Entity.

	entity.AddComponent(components.PlayerComponentID, &components.PlayerComponent{})
	entity.AddComponent(components.PositionComponentID, &components.PositionComponent{
		X:     posX,
		Y:     posY,
		TileX: math.Round(posX / config.TileSize),
		TileY: math.Round(posY / config.TileSize),
	})
	entity.AddComponent(components.VelocityComponentID, &components.VelocityComponent{VX: velX, VY: velY})
	entity.AddComponent(components.SpriteComponentID, &components.SpriteComponent{
		Image: sprite,
		X:     0, Y: 0, X1: 16, Y1: 16, // Initial sub-image coordinates
	})
	entity.AddComponent(components.CollisionComponentID, &components.CollisionComponent{})
	entity.AddComponent(components.CollisionBoxID, &components.CollisionBox{})
	entity.AddComponent(components.DestinationComponentID, &components.DestinationComponent{X: posX, Y: posY})

	// Combat components
	entity.AddComponent(components.AttackerComponentID, &components.AttackerComponent{
		IsAttacking: false,
		Target:      -1,
	})
	entity.AddComponent(components.DamageComponentID, &components.DamageComponent{})
	entity.AddComponent(components.SkillsComponentID, &components.SkillsComponent{
		CurrentHealth: 15,
		Health:        15,
		Attack:        5,
		Strength:      5,
		Defence:       5,
	})
	return entity
}
