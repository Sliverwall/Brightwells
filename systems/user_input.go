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
			// Get needed compontents
			// sprite := entity.GetComponent(components.SpriteComponentID).(*components.SpriteComponent)
			camera := entity.GetComponent(components.CameraComponentID).(*components.CameraComponent)
			position := entity.GetComponent(components.PositionComponentID).(*components.PositionComponent)
			velocity := entity.GetComponent(components.VelocityComponentID).(*components.VelocityComponent)
			destination := entity.GetComponent(components.DestinationComponentID).(*components.DestinationComponent)
			attacker := entity.GetComponent(components.AttackerComponentID).(*components.AttackerComponent)

			// Log for debugging
			if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
				// Get tile data on what was clicked
				checkX, checkY := ebiten.CursorPosition()
				checkX += int(camera.X)
				checkY += int(camera.Y)
				checkTileX := math.Floor(float64(checkX) / config.TileSize)
				checkTileY := math.Floor(float64(checkY) / config.TileSize)

				// Use tile data to grab entity ID on targeted tile
				checkEntityID := CheckTileForEntity(checkTileX, checkTileY, entitySlice)

				// check if there is no entity
				if checkEntityID != -1 {
					targetEntity := entities.GetEntityByID(entitySlice, checkEntityID)

					// Use target's ID to get entities compontents
					targetEntityPosition := targetEntity.GetComponent(components.PositionComponentID).(*components.PositionComponent)

					// Set player's target id to entity clicked
					attacker.Target = checkEntityID

					// Set coords to start moving to be target's current tile
					destX, destY := targetEntityPosition.TileX, targetEntityPosition.TileY

					destination.X = destX
					destination.Y = destY
					log.Print(checkEntityID)
				}
			}

			// ----------Left click START---------
			if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
				// Check if player is attacking, then turn off target and attacking
				if attacker.IsAttacking {
					attacker.Target = -1
					attacker.IsAttacking = false
				}
				// Capture x,y vector clicked on
				x, y := ebiten.CursorPosition()

				// Adjust for camera offset
				x += int(camera.X)
				y += int(camera.Y)

				// Mark destination to move to
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
