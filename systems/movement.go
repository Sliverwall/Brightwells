package systems

import (
	"Brightwells/components"
	"Brightwells/config"
	"Brightwells/entities"
	"math"
	"time"
)

type MovementSystem struct {
	CollisionSystem *CollisionSystem
	LastUpdateTime  time.Time
	UpdateInterval  time.Duration
}

func (ms *MovementSystem) Update(entitySlice []*entities.Entity) {
	currentTime := time.Now()
	if currentTime.Sub(ms.LastUpdateTime) < ms.UpdateInterval {
		return // Not enough time has passed, skip update
	}
	ms.LastUpdateTime = currentTime

	for _, entity := range entitySlice {
		if entity.HasComponent(components.PositionComponentID) && entity.HasComponent(components.VelocityComponentID) {
			// Update position based on velocity
			position := entity.GetComponent(components.PositionComponentID).(*components.PositionComponent)
			velocity := entity.GetComponent(components.VelocityComponentID).(*components.VelocityComponent)

			// Calculate movement for this frame
			futureTileX := position.TileX + velocity.VX
			futureTileY := position.TileY + velocity.VY
			futurePositionX := math.Floor(futureTileX * config.TileSize)
			futurePositionY := math.Floor(futureTileY * config.TileSize)

			if !entity.HasComponent(components.CollisionBoxID) {
				// No need to check collision boxes, Move to the next tile
				position.TileX = futureTileX
				position.TileY = futureTileY
				position.X = futurePositionX
				position.Y = futurePositionY
				return
			}

			// Check for collisions at the future position
			if ms.CollisionSystem.IsTileOccupiedByCollidableEntity(futureTileX, futureTileY, entitySlice) {
				// Prevent movement
				continue
			} else {
				// Move to the next tile
				position.TileX = futureTileX
				position.TileY = futureTileY
				position.X = futurePositionX
				position.Y = futurePositionY
			}
		}
	}
}
