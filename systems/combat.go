package systems

import (
	"Brightwells/components"
	"Brightwells/entities"
)

// ------------------------------ DAMAGE SYSTEMS -------------------------------
type DamageSystem struct{}

func (ds *DamageSystem) Update(entitySlice []*entities.Entity) {
	for _, attacker := range entitySlice {
		if attacker.HasComponent(components.AttackerComponentID) && attacker.HasComponent(components.SkillsComponentID) {
			attackComponent := attacker.GetComponent(components.AttackerComponentID).(*components.AttackerComponent)
			if attackComponent.Target != -1 {
				targetID := attackComponent.Target
				for _, target := range entitySlice {
					if target.ID == targetID && target.HasComponent(components.SkillsComponentID) && target.HasComponent(components.DamageComponentID) && target.ID != attacker.ID {
						// Grab desition and position compontent from attacker and target to keep adjusting destination position
						targetPosition := target.GetComponent(components.PositionComponentID).(*components.PositionComponent)
						attackerDestination := attacker.GetComponent(components.DestinationComponentID).(*components.DestinationComponent)
						// Keep updating attacker position to follow target if both are moving
						attackerDestination.X, attackerDestination.Y = targetPosition.TileX, targetPosition.TileY
						if IsWithinOneTile(attacker, target) {
							// Reset attacker's destination tile to current tile after reaching target
							attackerPosition := attacker.GetComponent(components.PositionComponentID).(*components.PositionComponent)

							attackerDestination.X, attackerDestination.Y = attackerPosition.TileX, attackerPosition.TileY
							// flag attacker as attacking
							attackComponent.IsAttacking = true
							// Retrieve the SkillsComponent of the target
							targetSkills := target.GetComponent(components.SkillsComponentID).(*components.SkillsComponent)

							// Calculate damage
							damage := 1

							// Deal damage to the target
							targetSkills.CurrentHealth -= damage

							// Ensure health does not go below zero
							if targetSkills.CurrentHealth < 0 {
								targetSkills.CurrentHealth = 0
								// Reset attack status for the attacker
								attackComponent.IsAttacking = false
								attackComponent.Target = -1
							}

							// Print damage dealt for debugging
							println("Entity", attacker.ID, "attacked Entity", target.ID, "for", damage, "damage!")
						}
					}
				}
			}
		}
	}
}

// ------------------------------ DEATH SYSTEMS -------------------------------
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

			// Add loot system later

			// Kill entity
			entity.KillEntity()

		}
	}
}
