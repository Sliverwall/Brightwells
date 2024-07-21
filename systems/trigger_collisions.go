package systems

import (
	"Brightwells/entities"
)

type TriggerCollisionSystem struct {
	FoodRespawnSystem *FoodRespawnSystem
	CollisionSystem   *CollisionSystem
}

func (tcs *TriggerCollisionSystem) Update(entitySlice []*entities.Entity) {
	// Update entitySlice if needed
	collisions := tcs.CollisionSystem.CheckTileCollisions(entitySlice)

	// Handle other triggers if needed
	tcs.FoodRespawnSystem.FoodCollide(entitySlice, collisions)
}
