package systems

import (
	"Brightwells/components"
	"Brightwells/entities"
	"log"
)

type FoodRespawnSystem struct {
}

func (frs *FoodRespawnSystem) FoodCollide(entitySlice []*entities.Entity, collisions map[int][]int) {

	for entityID := range collisions {
		entity := entities.GetEntityByID(entitySlice, entityID)
		if entity.HasComponent(components.FoodComponentID) {
			log.Println("Ate apple")
			entity.KillEntity()

		}
	}
}
