package systems

import (
	"Brightwells/components"
	"Brightwells/entities"
)

type MoveCollideSystem struct {
}

func (mcs *MoveCollideSystem) Update(entitySlice []*entities.Entity, entity *entities.Entity, collidingEntityIDs []int) {

	// Handle other entity collisions (e.g., movement rollback)
	if entity.HasComponent(components.PositionComponentID) && entity.HasComponent(components.VelocityComponentID) {
		position := entity.GetComponent(components.PositionComponentID).(*components.PositionComponent)
		velocity := entity.GetComponent(components.VelocityComponentID).(*components.VelocityComponent)

		// Save current position for potential rollback
		oldX, oldY := position.X, position.Y

		// Process the collisions
		for _, collidingEntityID := range collidingEntityIDs {
			collidingEntity := entities.GetEntityByID(entitySlice, collidingEntityID)
			if collidingEntity.HasComponent(components.CollisionComponentID) {
				// Rollback to previous position if collision detected
				position.X = oldX
				position.Y = oldY
				velocity.VX = -velocity.VX
				velocity.VY = -velocity.VY
			}
		}
	}

}
