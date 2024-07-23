package systems

import (
	"Brightwells/components"
	"Brightwells/entities"
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
