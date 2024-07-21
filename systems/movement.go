package systems

import (
	"Brightwells/components"
	"Brightwells/config"
	"Brightwells/entities"
	"math"
)

type MovementSystem struct {
	CollisionSystem *CollisionSystem
}

func (ms *MovementSystem) Update(entitySlice []*entities.Entity) {
	for _, entity := range entitySlice {
		// var fps float64
		// if ebiten.ActualFPS() == 0 {
		// 	fps = 60.0
		// } else {
		// 	fps = ebiten.ActualFPS()
		// }
		// dt := 1.0 / fps // Assume 60 FPS for movement calculations; adjust as needed
		if entity.HasComponent(components.PositionComponentID) && entity.HasComponent(components.VelocityComponentID) {
			// Update position based on velocity
			position := entity.GetComponent(components.PositionComponentID).(*components.PositionComponent)
			velocity := entity.GetComponent(components.VelocityComponentID).(*components.VelocityComponent)

			// Calculate movement for this frame
			futureTileX := position.TileX + velocity.VX
			futureTileY := position.TileY + velocity.VY
			futrePositionX := math.Round(position.TileX / config.TileSize)
			futrePositionY := math.Round(position.TileY / config.TileSize)

			if !entity.HasComponent(components.CollisionBoxID) {
				// No need to check collision boxes, Move to the next tile
				position.TileX = futureTileX
				position.TileY = futureTileY
				position.X = futrePositionX
				position.Y = futrePositionY
				return
			}

			// Check for collisions at the future position
			if ms.CollisionSystem.IsTileOccupiedByCollidableEntity(futureTileX, futureTileY, entitySlice) {
				// Prevent movement
				velocity.VX, velocity.VY = 0.0, 0.0
			} else {
				// Move to the next tile
				position.TileX = futureTileX
				position.TileY = futureTileY
				position.X = futrePositionX
				position.Y = futrePositionY
			}
		}
	}
}
