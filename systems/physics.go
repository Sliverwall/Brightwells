package systems

import (
	"Brightwells/components"
	"Brightwells/entities"
)

// ------------------------------ COLLISION SYSTEMS -------------------------------
type CollisionSystem struct {
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

// ----- TRIGGER COLLISIONS--------
type TriggerCollisionSystem struct {
	FoodRespawnSystem *FoodRespawnSystem
	CollisionSystem   *CollisionSystem
}

// TriggerCollisionSystem Update handles firing collision affects
func (tcs *TriggerCollisionSystem) Update(entitySlice []*entities.Entity) {
	// Update entitySlice if needed
	collisions := tcs.CollisionSystem.CheckTileCollisions(entitySlice)

	// Handle collision functions here

	tcs.FoodRespawnSystem.FoodCollide(entitySlice, collisions) // Dropped food effect
}

// ------------------------------ MOVEMENT SYSTEMS -------------------------------
type MovementSystem struct {
}

func (ms *MovementSystem) Update(entitySlice []*entities.Entity) {

	for _, entity := range entitySlice {
		if entity.HasComponent(components.PositionComponentID) && entity.HasComponent(components.VelocityComponentID) {
			// Update position based on velocity
			position := entity.GetComponent(components.PositionComponentID).(*components.PositionComponent)
			velocity := entity.GetComponent(components.VelocityComponentID).(*components.VelocityComponent)

			// Calculate movement for this frame
			futureTileX := position.TileX + velocity.VX
			futureTileY := position.TileY + velocity.VY

			if !entity.HasComponent(components.CollisionBoxID) {
				// No need to check collision boxes, Move to the next tile
				position.TileX = futureTileX
				position.TileY = futureTileY
				return
			}

			// Check for collisions at the future position
			if IsTileOccupiedByCollidableEntity(futureTileX, futureTileY, entitySlice) {
				// Prevent movement
				continue
			} else {
				// Move to the next tile
				position.TileX = futureTileX
				position.TileY = futureTileY
			}
		}
	}
}
