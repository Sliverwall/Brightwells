package systems

import (
	"Brightwells/components"
	"Brightwells/config"
	"Brightwells/entities"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type UserInputSystem struct{}

func (uis *UserInputSystem) Update(entitySlice []*entities.Entity) {
	for _, entity := range entitySlice {
		// Checks for players that have position and velocity
		if entity.HasComponent(components.PlayerComponentID) {
			sprite := entity.GetComponent(components.SpriteComponentID).(*components.SpriteComponent)
			position := entity.GetComponent(components.PositionComponentID).(*components.PositionComponent)
			velocity := entity.GetComponent(components.VelocityComponentID).(*components.VelocityComponent)
			destination := entity.GetComponent(components.DestinationComponentID).(*components.DestinationComponent)
			sprite.X, sprite.Y, sprite.X1, sprite.Y1 = 0, 0, 16, 16

			// Log for debugging
			if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
				log.Print("Destination X/Y: ", destination.X, destination.Y, " Current Pos X/Y: ", position.TileX, position.TileY)
			}

			// ----------Left click START---------
			if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
				x, y := ebiten.CursorPosition()
				destX := float64(x) / config.TileSize
				destY := float64(y) / config.TileSize
				destination.X = math.Floor(destX)
				destination.Y = math.Floor(destY)
				log.Print("Clicked tileX,tileY: ", destination.X, destination.Y)
			}
			// Control movement
			speed := 1.0
			// Pathfinding
			if position.TileX < destination.X {
				velocity.VX = speed
				velocity.VY = 0
			} else if position.TileX > destination.X {
				velocity.VX = -speed
				velocity.VY = 0
			} else if position.TileY < destination.Y {
				velocity.VY = speed
				velocity.VX = 0
			} else if position.TileY > destination.Y {
				velocity.VY = -speed
				velocity.VX = 0
			} else {
				velocity.VX = 0
				velocity.VY = 0
			}
			// ----------Left click END---------
		}
	}
}
