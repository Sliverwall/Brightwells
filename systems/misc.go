package systems

import (
	"Brightwells/components"
	"Brightwells/entities"
	"math"
)

// Module to hold useful functions not tied to any one system

func CheckTileForEntity(tileX, tileY float64, entitySlice []*entities.Entity) int {
	for _, entity := range entitySlice {
		if !entity.HasComponent(components.PositionComponentID) {
			continue
		}
		position := entity.GetComponent(components.PositionComponentID).(*components.PositionComponent)
		// Check if entity is on the tile being checked
		if position.TileX == tileX && position.TileY == tileY {
			// Grab the id then exit function
			return entity.ID
		}
	}
	// Return -1 if no entity is on tile
	return -1
}

// IsTileOccupiedByColidableEntity method to check if a tile is occupied by a collidable entity
func IsTileOccupiedByCollidableEntity(tileX, tileY float64, entitySlice []*entities.Entity) bool {
	for _, entity := range entitySlice {
		if !entity.HasComponent(components.CollisionBoxID) || !entity.HasComponent(components.PositionComponentID) {
			continue
		}
		position := entity.GetComponent(components.PositionComponentID).(*components.PositionComponent)
		if position.TileX == tileX && position.TileY == tileY {
			return true
		}
	}
	return false
}

// IsWithinOneTile checks if two entities are within one tile of eachother
func IsWithinOneTile(entity1, entity2 *entities.Entity) bool {
	position1 := entity1.GetComponent(components.PositionComponentID).(*components.PositionComponent)
	position2 := entity2.GetComponent(components.PositionComponentID).(*components.PositionComponent)

	deltaX := math.Abs(position1.TileX - position2.TileX)
	deltaY := math.Abs(position1.TileY - position2.TileY)

	// Check if the entities are within one tile of each other in any direction
	return deltaX <= 1 && deltaY <= 1
}
