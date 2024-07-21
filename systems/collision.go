package systems

import (
	"Brightwells/components"
	"Brightwells/entities"
	"log"
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

// WillCollide checks if moving to the specified tile will cause a collision with any other entity based on bounding boxes
func (cs *CollisionSystem) WillCollide(futureTileX, futureTileY float64, movingEntity *entities.Entity, entitySlice []*entities.Entity) bool {
	for _, entity := range entitySlice {
		// Check if entity has a collision box and position, and if it's not the moving entity being assessed
		if !entity.HasComponent(components.CollisionBoxID) || !entity.HasComponent(components.PositionComponentID) || entity.ID == movingEntity.ID {
			continue
		}

		// Get the bounding boxes for both the moving entity and the other entity
		box1 := movingEntity.GetComponent(components.CollisionBoxID).(*components.CollisionBox)
		box2 := entity.GetComponent(components.CollisionBoxID).(*components.CollisionBox)

		x1, y1, x2, y2 := box1.BoundingBox()
		ox1, oy1, ox2, oy2 := box2.BoundingBox()

		// Debug information to verify positions and bounding boxes
		log.Printf("Checking collision for entity %d at future tile (%f, %f)", movingEntity.ID, futureTileX, futureTileY)
		log.Printf("Moving entity bounding box: (%f, %f, %f, %f)", x1, y1, x2, y2)
		log.Printf("Other entity bounding box: (%f, %f, %f, %f)", ox1, oy1, ox2, oy2)

		// Check if the bounding boxes overlap
		if !(x1 > ox2 || x2 < ox1 || y1 > oy2 || y2 < oy1) {
			log.Printf("Collision detected with entity %d", entity.ID)
			return true
		}
	}
	return false
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
