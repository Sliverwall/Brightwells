package systems

import (
	"Brightwells/components"
	"Brightwells/entities"
)

type CollisionSystem struct {
	GameMap [][]int
}

// CheckTileCollisions processes all entities and returns a map of collisions
func (cs *CollisionSystem) CheckTileCollisions(entitySlice []*entities.Entity) map[int][]int {
	collisions := make(map[int][]int)

	for _, entity := range entitySlice {
		if !entity.HasComponent(components.CollisionComponentID) || !entity.HasComponent(components.PositionComponentID) {
			continue
		}

		// Get the collision tile for the entity
		position := entity.GetComponent(components.PositionComponentID).(*components.PositionComponent)

		for _, otherEntity := range entitySlice {
			if entity.ID == otherEntity.ID {
				continue
			}

			if !otherEntity.HasComponent(components.CollisionComponentID) || !otherEntity.HasComponent(components.PositionComponentID) {
				continue
			}

			// Get the collision box for the other entity
			otherPosition := otherEntity.GetComponent(components.PositionComponentID).(*components.PositionComponent)

			// Check if the bounding boxes overlap
			if cs.isOnSameTile(position.TileX, position.TileY, otherPosition.TileX, otherPosition.TileY) {
				collisions[entity.ID] = append(collisions[entity.ID], otherEntity.ID)
			}
		}
	}

	return collisions
}

// isOnSameTile checks if two entities are on the same tile
func (cs *CollisionSystem) isOnSameTile(X1, Y1, X2, Y2 float64) bool {
	return (X1 == X2 && Y1 == Y2)
}

// CollisionSystem method to check if a tile is occupied by a collidable entity
func (cs *CollisionSystem) IsTileOccupiedByCollidableEntity(tileX, tileY float64, entitySlice []*entities.Entity) bool {
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
