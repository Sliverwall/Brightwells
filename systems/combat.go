package systems

import (
	"Brightwells/components"
	"Brightwells/entities"
	"log"
)

// ------------------------------ DAMAGE SYSTEMS -------------------------------

// State for handling attacking.
func HandleAttacking(attacker *entities.Entity, entitySlice []*entities.Entity) {
	if attacker.HasComponent(components.AttackerComponentID) {
		attackComponent := attacker.GetComponent(components.AttackerComponentID).(*components.AttackerComponent)
		if attackComponent.Target != -1 {
			targetID := attackComponent.Target
			for _, target := range entitySlice {
				if target.ID == targetID && target.HasComponent(components.DamageComponentID) && target.ID != attacker.ID {
					// Check if target has died before proceeding
					targetState := target.GetComponent(components.StateComponentID).(*components.StateComponent)

					if targetState.CurrentState == components.StateDead {
						// Reset attacker to idle
						attackComponent.Target = -1
						SetNextState(attacker, components.StateIdle)
						break
					}

					// Grab desition and position compontent from attacker and target to keep adjusting destination position
					targetPosition := target.GetComponent(components.PositionComponentID).(*components.PositionComponent)
					attackerPosition := attacker.GetComponent(components.PositionComponentID).(*components.PositionComponent)
					// Keep updating attacker position to follow target if both are moving
					attackerPosition.DesX, attackerPosition.DesY = targetPosition.TileX, targetPosition.TileY
					if IsWithinOneTile(attacker, target) {
						// Reset attacker's destination tile to current tile after reaching target
						attackerPosition.DesX, attackerPosition.DesY = attackerPosition.TileX, attackerPosition.TileY

						// Calculate damage
						damage := 1
						// Retrieve the SkillsComponent of the target
						targetSkills := target.GetComponent(components.SkillsComponentID).(*components.SkillsComponent)
						// Deal damage to the target
						targetSkills.CurrentHealth -= damage
						log.Println(damage, " Damage occured to ", target.ID)
						// Check if entity should fight back
						if target.HasComponent(components.AttackedResponseComponentID) && targetSkills.CurrentHealth > 0 {
							targetAttackResponse := target.GetComponent(components.AttackedResponseComponentID).(*components.AttackedResponseComponent)
							// If set to AttackBack, grab and assign attacker ID to target's target
							if targetAttackResponse.Type == components.AttackBack {
								targetAttacking := target.GetComponent(components.AttackerComponentID).(*components.AttackerComponent)
								targetAttacking.Target = attacker.ID
								SetNextState(target, components.StateAttacking)
							}
						}

						// Print damage dealt for debugging
						println("Entity", attacker.ID, "attacked Entity", target.ID, "for", damage, "damage!")
					}
					if IsWithinManyTile(attacker, target, 5) {
						log.Println("Distance 1000 tiles")
						// Check if target is far, if so become idle
						attackComponent.Target = -1
						SetNextState(attacker, components.StateIdle)
					}
				}
			}
		}
	}
}

// ------------------------------ DEATH SYSTEMS -------------------------------
// UpdateCheckZeroHP update system handles killing entities if hp set to 0.
func UpdateCheckZeroHP(entitySlice []*entities.Entity) {
	for _, entity := range entitySlice {
		if entity.HasComponent(components.SkillsComponentID) && entity.HasComponent(components.StateComponentID) {
			skills := entity.GetComponent(components.SkillsComponentID).(*components.SkillsComponent)
			// entity has died
			if skills.CurrentHealth <= 0 {
				// DEBUG
				// log.Println("Triggerd HP:", skills.CurrentHealth)
				// Set state to dead
				SetNextState(entity, components.StateDead)
			}
		}
	}
}

// Handle Death system handles killing entities.
func HandleDeath(entity *entities.Entity) {
	// ALL CLEAN UP MUST BE DONE IN HANDLE DEATH

	// Reset attacker target
	if entity.HasComponent(components.AttackerComponentID) {
		attackerTarget := entity.GetComponent(components.AttackerComponentID).(*components.AttackerComponent)
		attackerTarget.Target = -1
	}
	// Reset skils back to max
	if entity.HasComponent(components.SkillsComponentID) {
		skills := entity.GetComponent(components.SkillsComponentID).(*components.SkillsComponent)
		skills.CurrentHealth = skills.Health
	}

	// Add loot system later

	if entity.Exist {
		// After respawn checks, kill the entity then set to idle
		entity.KillEntity()
		SetNextState(entity, components.StateIdle)
	}
}
