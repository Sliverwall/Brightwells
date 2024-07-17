package systems

import (
	"Brightwells/components"
	"Brightwells/entities"

	"github.com/hajimehoshi/ebiten/v2"
)

type UserInputSystem struct{}

func (uis *UserInputSystem) Update(entities []*entities.Entity) {
	for _, entity := range entities {
		// Checks for players that have position and velocity
		if entity.HasComponent(components.PlayerComponentID) {
			sprite := entity.GetComponent(components.SpriteComponentID).(*components.SpriteComponent)
			velocity := entity.GetComponent(components.VelocityComponentID).(*components.VelocityComponent)

			// Control movement
			if ebiten.IsKeyPressed(ebiten.KeyUp) {
				sprite.X, sprite.Y, sprite.X1, sprite.Y1 = 16, 16, 32, 32
				velocity.VY = -2
			} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
				sprite.X, sprite.Y, sprite.X1, sprite.Y1 = 0, 0, 16, 16
				velocity.VY = 2
			} else {
				velocity.VY = 0
			}

			if ebiten.IsKeyPressed(ebiten.KeyLeft) && !ebiten.IsKeyPressed(ebiten.KeyUp) && !ebiten.IsKeyPressed(ebiten.KeyDown) {
				sprite.X, sprite.Y, sprite.X1, sprite.Y1 = 32, 32, 48, 48
				velocity.VX = -2
			} else if ebiten.IsKeyPressed(ebiten.KeyRight) && !ebiten.IsKeyPressed(ebiten.KeyUp) && !ebiten.IsKeyPressed(ebiten.KeyDown) {
				sprite.X, sprite.Y, sprite.X1, sprite.Y1 = 48, 48, 64, 64
				velocity.VX = 2
			} else {
				velocity.VX = 0
			}

			// Switch Sprites

		}
	}
}
