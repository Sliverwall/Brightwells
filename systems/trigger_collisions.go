package systems

import (
	"Brightwells/entities"
)

type TriggerCollisionSystem struct {
	foodRespawnSystem *FoodRespawnSystem
	moveCollideSystem *MoveCollideSystem
}

// Update processes all collision events
func (tcs *TriggerCollisionSystem) Update(entitySlice []*entities.Entity, collisions map[int][]int) {
	for entityID, collidingEntityIDs := range collisions {
		entity := entities.GetEntityByID(entitySlice, entityID)

		// // Specific trigger: Food respawn
		// tcs.foodRespawnSystem.Update(entity)

		// Handle other entity collisions (e.g., movement rollback)
		tcs.moveCollideSystem.Update(entitySlice, entity, collidingEntityIDs)

	}
}
