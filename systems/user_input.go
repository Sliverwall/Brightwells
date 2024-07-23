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
				checkX, checkY := ebiten.CursorPosition()
				checkTileX := math.Floor(float64(checkX) / config.TileSize)
				checkTileY := math.Floor(float64(checkY) / config.TileSize)

				checkEntityID := CheckTileForEntity(checkTileX, checkTileY, entitySlice)

				attacker := entity.GetComponent(components.AttackerComponentID).(*components.AttackerComponent)
				attacker.IsAttacking = true
				attacker.Target = checkEntityID
				log.Print(checkEntityID)
			}

			// ----------Left click START---------
			if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
				x, y := ebiten.CursorPosition()
				destX := math.Floor(float64(x) / config.TileSize)
				destY := math.Floor(float64(y) / config.TileSize)
				destination.X = math.Floor(destX)
				destination.Y = math.Floor(destY)
				log.Print("Clicked tileX,tileY: ", destination.X, destination.Y)
			}
			// Control movement
			speed := 1.0
			// Pathfinding
			// X-axis
			if position.TileX < destination.X {
				velocity.VX = speed
			} else if position.TileX > destination.X {
				velocity.VX = -speed
			} else if position.TileX == destination.X {
				velocity.VX = 0
			}
			// Y-axis
			if position.TileY < destination.Y {
				velocity.VY = speed
			} else if position.TileY > destination.Y {
				velocity.VY = -speed
			} else if position.TileY == destination.Y {
				velocity.VY = 0
			}
			// ----------Left click END---------
		}
	}
}
