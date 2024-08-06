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
	RightClickTriggerOptions []int
	RightClickX, RightClickY int
)

type UserInputSystem struct {
}

func (uis *UserInputSystem) Update(entity *entities.Entity, entitySlice []*entities.Entity) {
	// Get needed compontents
	camera := entity.GetComponent(components.CameraComponentID).(*components.CameraComponent)
	// attacker := entity.GetComponent(components.AttackerComponentID).(*components.AttackerComponent)
	// gather := entity.GetComponent(components.GatherComponentID).(*components.GatherComponent)
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
				RightClickX, RightClickY = checkX-int(camera.X), checkY-int(camera.Y)

			}
		}
	}

	// ----------Left click START---------
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		// Set state to Idle
		SetNextState(entity, components.StateIdle)

		// Capture x,y vector clicked on
		x, y := ebiten.CursorPosition()

		// Adjust for camera offset
		x += int(camera.X)
		y += int(camera.Y)

		// Mark destination to move to
		destX := math.Floor(float64(x) / config.TileSize)
		destY := math.Floor(float64(y) / config.TileSize)
		position.DesX = math.Floor(destX)
		position.DesY = math.Floor(destY)
		log.Print("Clicked tileX,tileY: ", position.DesX, position.DesY)
	}
	// ----------Left click END---------
}
