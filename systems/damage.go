package systems

import (
	"Brightwells/components"
	"Brightwells/entities"
	"Brightwells/state"
	"time"
)

type DamageSystem struct {
	WorldInstance state.World
}

func (ds *DamageSystem) Update(entitySlice []*entities.Entity) {
	currentTime := time.Now()
	if currentTime.Sub(ds.WorldInstance.LastTick) < ds.WorldInstance.UpdateInterval {
		return // Not enough time has passed, skip update
	}
	ds.WorldInstance.LastTick = currentTime

	for _, attacker := range entitySlice {
		if attacker.HasComponent(components.AttackerComponentID) && attacker.HasComponent(components.SkillsComponentID) {
			attackComponent := attacker.GetComponent(components.AttackerComponentID).(*components.AttackerComponent)
			if attackComponent.Target != -1 {
				targetID := attackComponent.Target
				for _, target := range entitySlice {
					if target.ID == targetID && target.HasComponent(components.SkillsComponentID) && target.ID != attacker.ID {
						if IsWithinOneTile(attacker, target) {
							// Reset attacker's destination tile to current tile after reaching target
							attackerDestination := attacker.GetComponent(components.DestinationComponentID).(*components.DestinationComponent)
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
