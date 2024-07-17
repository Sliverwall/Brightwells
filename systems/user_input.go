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
		if entity.HasComponent(components.PlayerComponentID) &&
			entity.HasComponent(components.PositionComponentID) &&
			entity.HasComponent(components.VelocityComponentID) {
			velocity := entity.GetComponent(components.VelocityComponentID).(*components.VelocityComponent)

			if ebiten.IsKeyPressed(ebiten.KeyUp) {
				velocity.VY = -2
			} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
				velocity.VY = 2
			} else {
				velocity.VY = 0
			}

			if ebiten.IsKeyPressed(ebiten.KeyLeft) {
				velocity.VX = -2
			} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
				velocity.VX = 2
			} else {
				velocity.VX = 0
			}
		}
	}
}
