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
			destination := entity.GetComponent(components.DestinationComponentID).(*components.DestinationComponent)
			state := entity.GetComponent(components.StateComponentID).(*components.StateComponent)
			attacker := entity.GetComponent(components.AttackerComponentID).(*components.AttackerComponent)
			gather := entity.GetComponent(components.GatherComponentID).(*components.GatherComponent)

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
				if checkEntityID != -1 && checkEntityID != entity.ID {
					targetEntity := entities.GetEntityByID(entitySlice, checkEntityID)
					// Set player's target id to entity clicked
					if targetEntity.HasComponent(components.DamageComponentID) { // Check if entity is attackable
						state.NextState = 1 // Set to attacking
						attacker.Target = checkEntityID
					} else if targetEntity.HasComponent(components.ResourceNodeComponentID) { // Check if entity is gatherable
						state.NextState = 2 // Set to gathering
						gather.Target = checkEntityID
					}

					log.Print(checkEntityID)
				}
			}

			// ----------Left click START---------
			if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
				// Set state to Idle
				state.NextState = 0

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
			// ----------Left click END---------
		}
	}
}
