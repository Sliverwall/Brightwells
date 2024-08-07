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

var (
	// Trigger variables for passing information to main
	RightClickTriggerOptions []int
	RightClickX, RightClickY []int
	// Index offset for action text
	indexOffset int = 10
)

type UserInputSystem struct {
}

func (uis *UserInputSystem) Update(entity *entities.Entity, entitySlice []*entities.Entity) {
	// Get needed compontents
	camera := entity.GetComponent(components.CameraComponentID).(*components.CameraComponent)
	position := entity.GetComponent(components.PositionComponentID).(*components.PositionComponent)
	inventory := entity.GetComponent(components.InventoryComponentID).(*components.InventoryComponent)

	// ----------I click START---------

	if inpututil.IsKeyJustPressed(ebiten.KeyI) {
		log.Println(inventory.Items)
	}

	// ----------RIGHT click START---------
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		// Get tile data on what was clicked
		checkX, checkY := ebiten.CursorPosition()
		checkX += int(camera.X)
		checkY += int(camera.Y)
		checkTileX := math.Floor(float64(checkX) / config.TileSize)
		checkTileY := math.Floor(float64(checkY) / config.TileSize)

		// Use tile data to grab entity ID on targeted tile
		checkEntityID := CheckTileForEntity(checkTileX, checkTileY, entitySlice)

		// Check if there is no entity
		if checkEntityID != -1 && checkEntityID != entity.ID {
			// use the id to grab entity data
			targetEntity := entities.GetEntityByID(entitySlice, checkEntityID)

			// check if it has right click data
			if targetEntity.HasComponent(components.RightClickComponentID) {
				// grab right click data
				rightClickOptions := targetEntity.GetComponent(components.RightClickComponentID).(*components.RightClickComponent)
				// Show the context menu at the mouse position
				log.Println(rightClickOptions.Actions)

				// trigger draw menu
				RightClickTriggerOptions = rightClickOptions.Actions

				for index := range RightClickTriggerOptions {
					RightClickX = append(RightClickX, checkX-int(camera.X))                       // Index and X position
					RightClickY = append(RightClickY, (checkY-int(camera.Y))+(index*indexOffset)) // Index offset to move action down 1
				}

			}
		}
	}

	// ----------Left click START---------
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		// Capture x,y vector clicked on
		x, y := ebiten.CursorPosition()

		// Check if right click option is not nill
		if RightClickTriggerOptions != nil {

			// Set right click options to nil
			RightClickTriggerOptions = nil
		}

		// Set state to Walk Here
		SetNextState(entity, components.StateWalkHere)

		// Adjust for camera offset
		x += int(camera.X)
		y += int(camera.Y)

		// Check the tile for entities
		checkTileX := math.Floor(float64(x) / config.TileSize)
		checkTileY := math.Floor(float64(y) / config.TileSize)

		// Mark destination to move to
		destX := checkTileX
		destY := checkTileY
		position.DesX = math.Floor(destX)
		position.DesY = math.Floor(destY)
		log.Print("Clicked tileX,tileY: ", position.DesX, position.DesY)
	}
	// ----------Left click END---------
}
