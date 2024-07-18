package systems

import (
	"Brightwells/components"
	"Brightwells/config"
	"Brightwells/entities"
)

type CollisionSystem struct{}

func (cs *CollisionSystem) CheckBoundaryCollision(entity *entities.Entity) bool {
	windowWidth := config.WindowSize.Width
	windowHeight := config.WindowSize.Height

	if entity.HasComponent(components.CollisionComponentID) && entity.HasComponent(components.PositionComponentID) {
		collision := entity.GetComponent(components.CollisionComponentID).(*components.CollisionComponent)
		position := entity.GetComponent(components.PositionComponentID).(*components.PositionComponent)

		// Check collision with screen edges
		if position.X < 0 {
			position.X = 0
			return true
		} else if position.X+collision.Width > float64(windowWidth) {
			position.X = float64(windowWidth) - collision.Width
			return true
		}

		if position.Y < 0 {
			position.Y = 0
			return true
		} else if position.Y+collision.Height > float64(windowHeight) {
			position.Y = float64(windowHeight) - collision.Height
			return true
		}
	}
	return false
}

func (cs *CollisionSystem) CheckEntityCollisions(entities []*entities.Entity) map[int][]int {
	collisions := make(map[int][]int)
	for i, entity1 := range entities {
		if entity1.HasComponent(components.CollisionComponentID) && entity1.HasComponent(components.PositionComponentID) {
			for j, entity2 := range entities {
				if i != j && entity2.HasComponent(components.CollisionComponentID) && entity2.HasComponent(components.PositionComponentID) {
					if cs.CheckEntityCollision(entity1, entity2) {
						if _, ok := collisions[entity1.ID]; !ok {
							collisions[entity1.ID] = []int{}
						}
						collisions[entity1.ID] = append(collisions[entity1.ID], entity2.ID)
					}
				}
			}
		}
	}
	return collisions
}

func (cs *CollisionSystem) CheckEntityCollision(entity1, entity2 *entities.Entity) bool {
	coll1 := entity1.GetComponent(components.CollisionComponentID).(*components.CollisionComponent)
	pos1 := entity1.GetComponent(components.PositionComponentID).(*components.PositionComponent)
	coll2 := entity2.GetComponent(components.CollisionComponentID).(*components.CollisionComponent)
	pos2 := entity2.GetComponent(components.PositionComponentID).(*components.PositionComponent)

	return pos1.X < pos2.X+coll2.Width &&
		pos1.X+coll1.Width > pos2.X &&
		pos1.Y < pos2.Y+coll2.Height &&
		pos1.Y+coll1.Height > pos2.Y
}
