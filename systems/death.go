package systems

import (
	"Brightwells/components"
	"Brightwells/entities"
)

type DeathSystem struct{}

// DeathSystem's update system handles killing entities.
func (ds *DeathSystem) Update(entitySlice []*entities.Entity) {

	for _, entity := range entitySlice {
		// Skill component required to die, as it holds health stat
		if !entity.HasComponent(components.SkillsComponentID) {
			continue
		}

		skills := entity.GetComponent(components.SkillsComponentID).(*components.SkillsComponent)

		// entity has died
		if skills.CurrentHealth <= 0 {
			// reset currentHealth back to max
			skills.CurrentHealth = skills.Health

			// Kill entity
			entity.KillEntity()

		}
	}
}
