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
	for _, player := range entitySlice {
		if player.HasComponent(components.PlayerComponentID) {
			// Gather gather compontent
			gatherComponent := player.GetComponent(components.GatherComponentID).(*components.GatherComponent)
			// Check if target is not set to none (-1)
			if gatherComponent.Target != -1 {
				// Get skill component before going into next loop
				// skillComponent := player.GetComponent(components.SkillsComponentID).(*components.SkillsComponent)
				targetID := gatherComponent.Target
				for _, target := range entitySlice {
					if target.ID == targetID && target.HasComponent(components.ResourceNodeComponentID) && target.ID != player.ID {
						// Check if near node before begining checks
						if IsWithinOneTile(player, target) {
							// Reset player's destination tile to current tile after reaching target
							playerDestination := player.GetComponent(components.DestinationComponentID).(*components.DestinationComponent)
							playerPosition := player.GetComponent(components.PositionComponentID).(*components.PositionComponent)

							playerDestination.X, playerDestination.Y = playerPosition.TileX, playerPosition.TileY

							// Check active tag
							resourceCompontent := target.GetComponent(components.ResourceNodeComponentID).(*components.ResourceNodeComponent)
							if !resourceCompontent.Active {
								log.Println("Resource depleted...")
								gatherComponent.IsGathering = false
								gatherComponent.Target = -1
								break
							}
							// flag player as Gathering
							gatherComponent.IsGathering = true

							roll := rand.Float32()
							log.Println("Roll chance: ", roll)
							// Ensure health does not go below zero
							if roll <= resourceCompontent.DrainedChance {
								// Set active flag to false
								resourceCompontent.Active = false
								// Reset player status for the player
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
