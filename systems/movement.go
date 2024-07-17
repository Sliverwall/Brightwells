package systems

import (
	"Brightwells/components"
	"Brightwells/entities"
)

type MovementSystem struct {
	collisionSystem *CollisionSystem
}

func (ms *MovementSystem) Update(entities []*entities.Entity) {
	for _, entity := range entities {
		if entity.HasComponent(components.PositionComponentID) && entity.HasComponent(components.VelocityComponentID) {
			position := entity.GetComponent(components.PositionComponentID).(*components.PositionComponent)
			velocity := entity.GetComponent(components.VelocityComponentID).(*components.VelocityComponent)
			position.X += velocity.VX
			position.Y += velocity.VY

			// Check for collisions if the collision system is enabled
			if ms.collisionSystem.Update(entities) {
				velocity.VX = -velocity.VX
				velocity.VY = -velocity.VY
			}

		}
	}
}
