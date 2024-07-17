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

			// Save current position for potential rollback
			oldX, oldY := position.X, position.Y

			// Update position based on velocity
			position.X += velocity.VX
			position.Y += velocity.VY

			// Check for collisions using the collision system
			if ms.collisionSystem.Update(entities) {
				// If collision detected, rollback to previous position
				position.X = oldX
				position.Y = oldY

				// Optionally adjust velocity or perform other collision response actions
				velocity.VX = 0
				velocity.VY = 0
			}
		}
	}
}
