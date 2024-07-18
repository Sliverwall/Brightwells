package systems

import (
	"Brightwells/components"
	et "Brightwells/entities"
	"log"
	"math/rand"
)

type FoodRespawnSystem struct {
	collisionSystem *CollisionSystem
}

func (frs *FoodRespawnSystem) Update(entities []*et.Entity) {
	collisions := frs.collisionSystem.CheckEntityCollisions(entities)
	for foodID := range collisions {
		food := et.GetEntityByID(entities, foodID)
		if food.HasComponent(components.FoodComponentID) {
			position := food.GetComponent(components.PositionComponentID).(*components.PositionComponent)
			log.Println("Ate apple")
			position.X = float64(rand.Intn(100))
			position.Y = float64(rand.Intn(100))
			break
		}
	}
}
