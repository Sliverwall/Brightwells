package systems

import (
	"Brightwells/components"
	"Brightwells/entities"
	"log"
)

// ------------------------------ Inventory SYSTEMS -------------------------------
type InventorySystem struct{}

func (is *InventorySystem) Update(player *entities.Entity) {
	log.Panicln("Inventory")
}

// ------------------------------ FOOD SYSTEMS -------------------------------
func (tcs *TriggerCollisionSystem) FoodCollide(entitySlice []*entities.Entity, collisions map[int][]int) {
	for entityID := range collisions {
		entity := entities.GetEntityByID(entitySlice, entityID)
		if entity.HasComponent(components.FoodComponentID) {
			log.Println("Ate apple")
			entity.KillEntity()

		}
	}
}
