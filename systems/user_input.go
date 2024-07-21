package systems

import (
	"Brightwells/components"
	"Brightwells/config"
	"Brightwells/entities"
	"log"
	"math"

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
			speed := config.TileSize
			if ebiten.IsKeyPressed(ebiten.KeyUp) {
				sprite.X, sprite.Y, sprite.X1, sprite.Y1 = 16, 16, 16, 16
				velocity.VY = -speed
			} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
				sprite.X, sprite.Y, sprite.X1, sprite.Y1 = 0, 0, 16, 16
				velocity.VY = speed
			} else {
				velocity.VY = 0
			}

			if ebiten.IsKeyPressed(ebiten.KeyLeft) && !ebiten.IsKeyPressed(ebiten.KeyUp) && !ebiten.IsKeyPressed(ebiten.KeyDown) {
				sprite.X, sprite.Y, sprite.X1, sprite.Y1 = 32, 32, 16, 16
				velocity.VX = -speed
			} else if ebiten.IsKeyPressed(ebiten.KeyRight) && !ebiten.IsKeyPressed(ebiten.KeyUp) && !ebiten.IsKeyPressed(ebiten.KeyDown) {
				sprite.X, sprite.Y, sprite.X1, sprite.Y1 = 48, 48, 16, 16
				velocity.VX = speed
			} else {
				velocity.VX = 0
			}

			// Log for debugging
			if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
				x, y := ebiten.CursorPosition()
				destX := math.Round(float64(x) / config.TileSize)
				destY := math.Round(float64(y) / config.TileSize)
				log.Print("clicked tileX,tileY/X,Y", destX, destY, "/", x, y)

			}

		}
	}
}
