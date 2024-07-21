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

		direction := 4 // default direction (idle)
		if entity.HasComponent(components.VelocityComponentID) {
			velocity := entity.GetComponent(components.VelocityComponentID).(*components.VelocityComponent)
			direction = cs.checkDirection(velocity.VX, velocity.VY)
		}

		box1 := entity.GetComponent(components.CollisionBoxID).(*components.CollisionBox)
		tileX1, tileY1, tileX2, tileY2 := box1.BoundingBox(direction)

		for _, otherEntity := range entitySlice {
			if entity.ID == otherEntity.ID {
				continue
			}

			if !otherEntity.HasComponent(components.CollisionComponentID) || !otherEntity.HasComponent(components.PositionComponentID) {
				continue
			}

			otherDirection := 4 // default direction (idle)
			if otherEntity.HasComponent(components.VelocityComponentID) {
				otherVelocity := otherEntity.GetComponent(components.VelocityComponentID).(*components.VelocityComponent)
				otherDirection = cs.checkDirection(otherVelocity.VX, otherVelocity.VY)
			}

			box2 := otherEntity.GetComponent(components.CollisionBoxID).(*components.CollisionBox)
			otherTileX1, otherTileY1, otherTileX2, otherTileY2 := box2.BoundingBox(otherDirection)

			if cs.isOverlapping(tileX1, tileY1, tileX2, tileY2, otherTileX1, otherTileY1, otherTileX2, otherTileY2) {
				collisions[entity.ID] = append(collisions[entity.ID], otherEntity.ID)
			}
		}
	}

	return collisions
}

// isOverlapping checks if two sets of bounding box tiles overlap
func (cs *CollisionSystem) isOverlapping(x1, y1, x2, y2, ox1, oy1, ox2, oy2 float64) bool {
	return !(x1 > ox2 || x2 < ox1 || y1 > oy2 || y2 < oy1)
}

// checkDirection returns the direction based on the velocity components
func (cs *CollisionSystem) checkDirection(vx, vy float64) int {
	switch {
	case vx > 0:
		return 0 // right
	case vx < 0:
		return 1 // left
	case vy < 0:
		return 2 // up
	case vy > 0:
		return 3 // down
	default:
		return 4 // idle
	}
}
