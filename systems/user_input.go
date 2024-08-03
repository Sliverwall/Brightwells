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

func (uis *UserInputSystem) Update(entity *entities.Entity, entitySlice []*entities.Entity) {
	// Get needed compontents
	// sprite := entity.GetComponent(components.SpriteComponentID).(*components.SpriteComponent)
	camera := entity.GetComponent(components.CameraComponentID).(*components.CameraComponent)
	attacker := entity.GetComponent(components.AttackerComponentID).(*components.AttackerComponent)
	gather := entity.GetComponent(components.GatherComponentID).(*components.GatherComponent)
	position := entity.GetComponent(components.PositionComponentID).(*components.PositionComponent)

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
				SetNextState(entity, components.StateAttacking)
				attacker.Target = checkEntityID
			} else if targetEntity.HasComponent(components.ResourceNodeComponentID) { // Check if entity is gatherable
				SetNextState(entity, components.StateGather)
				gather.Target = checkEntityID
			}

			log.Print(checkEntityID)
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
