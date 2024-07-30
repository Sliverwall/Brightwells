package systems

import (
	"Brightwells/components"
	"Brightwells/entities"
	"log"
	"math/rand"
)

// ------------------------------ ResourceNode SYSTEMS -------------------------------
type ResourceNodeSystem struct {
}

func (ls *ResourceNodeSystem) Update(entitySlice []*entities.Entity) {
	for _, gather := range entitySlice {
		if gather.HasComponent(components.GatherComponentID) {
			// Gather gather compontent
			gatherComponent := gather.GetComponent(components.GatherComponentID).(*components.GatherComponent)
			// Check if target is not set to none (-1)
			if gatherComponent.Target != -1 {
				targetID := gatherComponent.Target
				for _, target := range entitySlice {
					if target.ID == targetID && target.HasComponent(components.ResourceNodeComponentID) && target.ID != gather.ID {
						// Grab desition and position compontent from attacker and target to keep adjusting destination position
						gatherDestination := gather.GetComponent(components.DestinationComponentID).(*components.DestinationComponent)
						targetPosition := target.GetComponent(components.PositionComponentID).(*components.PositionComponent)
						// Keep updating gather position to follow target if both are moving
						gatherDestination.X, gatherDestination.Y = targetPosition.TileX, targetPosition.TileY
						// Check if near node before begining checks
						if IsWithinOneTile(gather, target) {
							// Reset gather's destination tile to current tile after reaching target
							gatherPosition := gather.GetComponent(components.PositionComponentID).(*components.PositionComponent)
							gatherDestination.X, gatherDestination.Y = gatherPosition.TileX, gatherPosition.TileY

							// Check active tag
							resourceCompontent := target.GetComponent(components.ResourceNodeComponentID).(*components.ResourceNodeComponent)
							if !resourceCompontent.Active {
								log.Println("Resource depleted...")
								gatherComponent.IsGathering = false
								gatherComponent.Target = -1
								break
							}
							// flag gather as Gathering
							gatherComponent.IsGathering = true

							roll := rand.Float32()
							log.Println("Roll chance: ", roll)
							// Ensure health does not go below zero
							if roll <= resourceCompontent.DrainedChance {
								// Set active flag to false
								resourceCompontent.Active = false
								// Reset gather status for the gather
								gatherComponent.IsGathering = false
								gatherComponent.Target = -1
								// Leave loop
								break
							}
						}
					}
				}
			}
		}
	}
}