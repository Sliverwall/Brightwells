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

		box1 := entity.GetComponent(components.CollisionBoxID).(*components.CollisionBox)

		for _, otherEntity := range entitySlice {
			if entity.ID == otherEntity.ID {
				continue
			}

			if !otherEntity.HasComponent(components.CollisionComponentID) || !otherEntity.HasComponent(components.PositionComponentID) {
				continue
			}

			box2 := otherEntity.GetComponent(components.CollisionBoxID).(*components.CollisionBox)

			if cs.isOverlapping(box1, box2) {
				collisions[entity.ID] = append(collisions[entity.ID], otherEntity.ID)
			}
		}
	}

	return collisions
}

// isOverlapping checks if two bounding boxes overlap
func (cs *CollisionSystem) isOverlapping(box1, box2 *components.CollisionBox) bool {
	return !(box1.PositionComponent.X > box2.PositionComponent.X+box2.CollisionComponent.Width ||
		box1.PositionComponent.X+box1.CollisionComponent.Width < box2.PositionComponent.X ||
		box1.PositionComponent.Y > box2.PositionComponent.Y+box2.CollisionComponent.Height ||
		box1.PositionComponent.Y+box1.CollisionComponent.Height < box2.PositionComponent.Y)
}
