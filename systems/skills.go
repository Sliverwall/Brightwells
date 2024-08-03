package systems

import (
	"Brightwells/components"
	"Brightwells/entities"
	"log"
	"math/rand"
)

// ------------------------------ ResourceNode SYSTEMS -------------------------------

func HandleGathering(gather *entities.Entity, entitySlice []*entities.Entity) {
	if gather.HasComponent(components.GatherComponentID) {
		// Gather gather compontent
		gatherComponent := gather.GetComponent(components.GatherComponentID).(*components.GatherComponent)
		// Check if target is not set to none (-1)
		if gatherComponent.Target != -1 {
			targetID := gatherComponent.Target
			for _, target := range entitySlice {
				if target.ID == targetID && target.HasComponent(components.ResourceNodeComponentID) && target.ID != gather.ID {
					// Grab desition and position compontent from attacker and target to keep adjusting destination position
					targetPosition := target.GetComponent(components.PositionComponentID).(*components.PositionComponent)
					// Keep updating gather position to follow target if both are moving
					gatherPosition := gather.GetComponent(components.PositionComponentID).(*components.PositionComponent)
					gatherPosition.DesX, gatherPosition.DesY = targetPosition.TileX, targetPosition.TileY
					// Check if near node before begining checks
					if IsWithinOneTile(gather, target) {
						// Reset gather's destination tile to current tile after reaching target
						gatherPosition.DesX, gatherPosition.DesY = gatherPosition.TileX, gatherPosition.TileY

						// Check active tag
						resourceCompontent := target.GetComponent(components.ResourceNodeComponentID).(*components.ResourceNodeComponent)
						if !resourceCompontent.Active {
							log.Println("Resource depleted...")
							gatherComponent.Target = -1
							// Set next state to idle
							SetNextState(gather, components.StateIdle)
							break
						}
						roll := rand.Float32()
						log.Println("Roll chance: ", roll)
						// Ensure health does not go below zero
						if roll <= resourceCompontent.DrainedChance {
							// Set active flag to false
							resourceCompontent.Active = false
							// Reset gather status for the gather
							gatherComponent.Target = -1
							// Leave loop
							// Set next state to idle
							SetNextState(gather, components.StateIdle)
							break
						}
					}
				}
			}
		}
	}
}

func SetNextState(entity *entities.Entity, state int) {
	// Set next step
	if entity.HasComponent(components.StateComponentID) {
		entity.GetComponent(components.StateComponentID).(*components.StateComponent).NextState = state
	} else {
		log.Println(entity.ID, " Has no state")
	}
}
