package systems

import (
	"Brightwells/components"
	"Brightwells/entities"
	"log"
)

// ------------------------------ DAMAGE SYSTEMS -------------------------------

// State for handling attacking.
func (ss *StateSystem) HandleAttacking(attacker *entities.Entity, entitySlice []*entities.Entity) {
	if attacker.HasComponent(components.AttackerComponentID) {
		attackComponent := attacker.GetComponent(components.AttackerComponentID).(*components.AttackerComponent)
		if attackComponent.Target != -1 {
			targetID := attackComponent.Target
			for _, target := range entitySlice {
				if target.ID == targetID && target.HasComponent(components.DamageComponentID) && target.ID != attacker.ID {
					// Grab desition and position compontent from attacker and target to keep adjusting destination position
					targetPosition := target.GetComponent(components.PositionComponentID).(*components.PositionComponent)
					attackerPosition := attacker.GetComponent(components.PositionComponentID).(*components.PositionComponent)
					// Keep updating attacker position to follow target if both are moving
					attackerPosition.DesX, attackerPosition.DesY = targetPosition.TileX, targetPosition.TileY
					if IsWithinOneTile(attacker, target) {
						// Reset attacker's destination tile to current tile after reaching target
						attackerPosition.DesX, attackerPosition.DesY = attackerPosition.TileX, attackerPosition.TileY
						// Retrieve the SkillsComponent of the target
						targetSkills := target.GetComponent(components.SkillsComponentID).(*components.SkillsComponent)

						// Calculate damage
						damage := 1

						// Deal damage to the target
						targetSkills.CurrentHealth -= damage
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
// DeathSystem's update system handles killing entities.
func UpdateDeath(entitySlice []*entities.Entity) {

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

			// Check if entity has a spawn point
			if entity.HasComponent(components.SpawnPointComponentID) {
				// Set position to spawn positions
				position := entity.GetComponent(components.PositionComponentID).(*components.PositionComponent)
				spawnPoint := entity.GetComponent(components.SpawnPointComponentID).(*components.SpawnPointComponent)

				position.TileX, position.TileY = spawnPoint.TileX, spawnPoint.TileY

				// Set Destination to spawn point as well
				position.DesX, position.DesY = spawnPoint.TileX, spawnPoint.TileY

				break // Leave loop before deleting entity
			}
			// Kill entity
			entity.KillEntity()

		}
	}
}
