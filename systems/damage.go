package systems

import (
	"Brightwells/entities"
)

type DamageSystem struct {
}

// DamageSystem Update takes a map of entities tagged as being attacked, then calculates damage dealt
func (ds *DamageSystem) Update(entitySlice []*entities.Entity) {
	return
	// for entityID := range entitySlice {
	// 	entity := entities.GetEntityByID(entitySlice, entityID)

	// 	entitySkills := entity.GetComponent(components.SkillsComponentID).(*components.SkillsComponent)

	// 	entityAttack := entitySkills.Attack

	// 	// need loop map whith all entities that are attacking.
	// 	var baseHitChance float64 = 0.5

	// 	var hitChance float64 = 1 - (baseHitChance * (targetDefence / entityAttack))

	// }
}
