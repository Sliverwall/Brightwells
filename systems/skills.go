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
				// Get skill component before going into next loop
				// skillComponent := gather.GetComponent(components.SkillsComponentID).(*components.SkillsComponent)
				targetID := gatherComponent.Target
				for _, target := range entitySlice {
					if target.ID == targetID && target.HasComponent(components.ResourceNodeComponentID) && target.ID != gather.ID {
						// Check if near node before begining checks
						if IsWithinOneTile(gather, target) {
							// Reset gather's destination tile to current tile after reaching target
							gatherDestination := gather.GetComponent(components.DestinationComponentID).(*components.DestinationComponent)
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
