package systems

import (
	"Brightwells/components"
	"Brightwells/entities"

	"github.com/hajimehoshi/ebiten/v2"
)

type UserInputSystem struct{}

func (uis *UserInputSystem) Update(entitySlice []*entities.Entity) {
	for _, entity := range entitySlice {
		// Checks for players that have position and velocity
		if entity.HasComponent(components.PlayerComponentID) {
			sprite := entity.GetComponent(components.SpriteComponentID).(*components.SpriteComponent)
			velocity := entity.GetComponent(components.VelocityComponentID).(*components.VelocityComponent)

			// Control movement
			speed := 64.0 // Speed is the size of one tile

			// Reset velocity
			velocity.VX = 0
			velocity.VY = 0

			if ebiten.IsKeyPressed(ebiten.KeyUp) {
				sprite.X, sprite.Y, sprite.X1, sprite.Y1 = 16, 16, 16, 16
				velocity.VY = -speed
			} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
				sprite.X, sprite.Y, sprite.X1, sprite.Y1 = 0, 0, 16, 16
				velocity.VY = speed
			}

			if ebiten.IsKeyPressed(ebiten.KeyLeft) && !ebiten.IsKeyPressed(ebiten.KeyUp) && !ebiten.IsKeyPressed(ebiten.KeyDown) {
				sprite.X, sprite.Y, sprite.X1, sprite.Y1 = 32, 32, 16, 16
				velocity.VX = -speed
			} else if ebiten.IsKeyPressed(ebiten.KeyRight) && !ebiten.IsKeyPressed(ebiten.KeyUp) && !ebiten.IsKeyPressed(ebiten.KeyDown) {
				sprite.X, sprite.Y, sprite.X1, sprite.Y1 = 48, 48, 16, 16
				velocity.VX = speed
			}
		}
	}
}
