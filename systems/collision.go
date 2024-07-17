package systems

import (
	"Brightwells/components"
	"Brightwells/entities"
	"fmt"
)

type CollisionSystem struct {
}

func (cs *CollisionSystem) Update(entities []*entities.Entity) bool {
	for _, entity := range entities {
		if entity.HasComponent(components.CollisionComponentID) && entity.HasComponent(components.PositionComponentID) {
			collision := entity.GetComponent(components.CollisionComponentID).(*components.CollisionComponent)
			position := entity.GetComponent(components.PositionComponentID).(*components.PositionComponent)
			// Perform collision detection
			for _, other := range entities {
				if entity == other || !other.HasComponent(components.CollisionComponentID) || !other.HasComponent(components.PositionComponentID) {
					continue
				}
				otherCollision := other.GetComponent(components.CollisionComponentID).(*components.CollisionComponent)
				otherPosition := other.GetComponent(components.PositionComponentID).(*components.PositionComponent)
				if collision.CollidesWith(position, otherCollision, otherPosition.X, otherPosition.Y) {
					fmt.Print("Collision Detected")
					return true
				}
			}
		}
	}
	return false
}
