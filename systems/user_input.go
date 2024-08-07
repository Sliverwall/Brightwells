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
	// Save right clicked entity id
	RightClickCheckEntityID int
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
	attacker := entity.GetComponent(components.AttackerComponentID).(*components.AttackerComponent)
	gather := entity.GetComponent(components.GatherComponentID).(*components.GatherComponent)

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
		RightClickCheckEntityID = CheckTileForEntity(checkTileX, checkTileY, entitySlice)

		// Check if there is no entity
		if RightClickCheckEntityID != -1 && RightClickCheckEntityID != entity.ID {
			log.Println(RightClickCheckEntityID)
			// use the id to grab entity data
			targetEntity := entities.GetEntityByID(entitySlice, RightClickCheckEntityID)

			// check if it has right click data
			if targetEntity.HasComponent(components.RightClickComponentID) {
				// Clean up options if previously used
				if RightClickX != nil || RightClickY != nil || RightClickTriggerOptions != nil {
					// Mark as nil so GC releases memory
					RightClickX, RightClickY, RightClickTriggerOptions = nil, nil, nil
				}
				// grab right click data
				rightClickOptions := targetEntity.GetComponent(components.RightClickComponentID).(*components.RightClickComponent)
				// Show the context menu at the mouse position
				log.Println(rightClickOptions.Actions)

				// trigger draw menu
				RightClickTriggerOptions = rightClickOptions.Actions

				for index := range RightClickTriggerOptions {
					RightClickX = append(RightClickX, checkX-int(camera.X)+(indexOffset))         // Index and X position
					RightClickY = append(RightClickY, (checkY-int(camera.Y))+(index*indexOffset)) // Index offset to move action down 1
				}

			}
		}

	}

	// ----------Left click START---------
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		// Adjust for camera offset
		x += int(camera.X)
		y += int(camera.Y)
		// Check the tile for entities
		checkTileX := math.Floor(float64(x) / config.TileSize)
		checkTileY := math.Floor(float64(y) / config.TileSize)

		// Use tile data to grab entity ID on targeted tile
		checkEntityID := CheckTileForEntity(checkTileX, checkTileY, entitySlice)
		// Check if there is no entity
		if checkEntityID != -1 && checkEntityID != entity.ID {
			log.Println(checkEntityID)
			// use the id to grab entity data
			targetEntity := entities.GetEntityByID(entitySlice, checkEntityID)
			// check if it has right click data
			if targetEntity.HasComponent(components.RightClickComponentID) {
				// grab right click data
				rightClickOptions := targetEntity.GetComponent(components.RightClickComponentID).(*components.RightClickComponent)
				// Select first rightClickOption as default leftclick
				switch rightClickOptions.Actions[0] {
				case components.StateAttacking:
					attacker.Target = checkEntityID
					SetNextState(entity, components.StateAttacking)
				case components.StateGather:
					gather.Target = checkEntityID
					SetNextState(entity, components.StateGather)
				}
			}
		} else {

			// Clear any right click optins
			RightClickX, RightClickY, RightClickTriggerOptions = nil, nil, nil
			// Set state to Walk Here
			SetNextState(entity, components.StateIdle)

			// Mark destination to move to
			destX := checkTileX
			destY := checkTileY
			position.DesX = math.Floor(destX)
			position.DesY = math.Floor(destY)
			log.Print("Clicked tileX,tileY: ", position.DesX, position.DesY)
		}

	}
	// ----------Left click END---------

	// ----------NUM KEYS START---------

	// Num keys when right click options are avaiable
	if RightClickTriggerOptions != nil {
		// First click option
		if inpututil.IsKeyJustPressed(ebiten.Key1) {
			// Pass in [0] for first option
			ActivateRightClick(RightClickTriggerOptions[0],
				RightClickCheckEntityID,
				entity, entitySlice,
				attacker, gather, position)
		} else if inpututil.IsKeyJustPressed(ebiten.Key2) {
			// Pass in [1] for second option
			ActivateRightClick(RightClickTriggerOptions[1],
				RightClickCheckEntityID,
				entity, entitySlice,
				attacker, gather, position)
		}
	}
	// ----------NUM KEYS END---------
}
