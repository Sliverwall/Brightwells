package systems

import (
	"Brightwells/components"
	"Brightwells/entities"
	"log"
	"math/rand"
)

type FoodRespawnSystem struct {
}

func (frs *FoodRespawnSystem) Update(entity *entities.Entity) {

	if entity.HasComponent(components.FoodComponentID) {
		position := entity.GetComponent(components.PositionComponentID).(*components.PositionComponent)
		log.Println("Ate apple")
		position.X = float64(rand.Intn(100))
		position.Y = float64(rand.Intn(100))
	}

}
