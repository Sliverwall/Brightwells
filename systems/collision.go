package systems

import (
	"Brightwells/components"
	"Brightwells/config"
	"Brightwells/entities"
)

type CollisionSystem struct {
}

func (cs *CollisionSystem) Update(entities []*entities.Entity) bool {
	// Grab screen dims from config
	windowWidth := config.WindowSize.Width
	windowHight := config.WindowSize.Height

	playerCollisionDetected := false
	// Loop through all entities to check for collisions with the player
	for _, entity := range entities {
		if entity.HasComponent(components.CollisionComponentID) && entity.HasComponent(components.PositionComponentID) {
			collision := entity.GetComponent(components.CollisionComponentID).(*components.CollisionComponent)
			position := entity.GetComponent(components.PositionComponentID).(*components.PositionComponent)

			// Check collision with screen edges
			if position.X < 0 {
				position.X = 0
				playerCollisionDetected = true
			} else if position.X+collision.Width > float64(windowWidth) {
				position.X = float64(windowHight) - collision.Width
				playerCollisionDetected = true
			}

			if position.Y < 0 {
				position.Y = 0
				playerCollisionDetected = true
			} else if position.Y+collision.Height > float64(windowHight) {
				position.Y = float64(windowWidth) - collision.Height
				playerCollisionDetected = true
			}

			// Check collision with other entities
			for _, other := range entities {
				if entity == other || !other.HasComponent(components.CollisionComponentID) {
					continue
				}
				otherCollision := other.GetComponent(components.CollisionComponentID).(*components.CollisionComponent)
				otherPosition := other.GetComponent(components.PositionComponentID).(*components.PositionComponent)

				if Collides(position, collision, otherPosition, otherCollision) {
					// Handle collision logic here, e.g., set flags, perform actions, etc.
					playerCollisionDetected = true
				}
			}
		}
	}

	return playerCollisionDetected
}

// Collides checks if two entities collide
func Collides(pos1 *components.PositionComponent, coll1 *components.CollisionComponent, pos2 *components.PositionComponent, coll2 *components.CollisionComponent) bool {
	return pos1.X < pos2.X+coll2.Width &&
		pos1.X+coll1.Width > pos2.X &&
		pos1.Y < pos2.Y+coll2.Height &&
		pos1.Y+coll1.Height > pos2.Y
}
