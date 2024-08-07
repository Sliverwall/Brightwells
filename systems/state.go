package systems

import (
	"Brightwells/components"
	"Brightwells/entities"
	"log"
)

// ------------------------------ STATE MACHINE SYSTEMS -------------------------------
type StateSystem struct{}

func UpdateState(entitySlice []*entities.Entity) {
	for _, entity := range entitySlice {
		if entity.HasComponent(components.StateComponentID) {
			stateComponent := entity.GetComponent(components.StateComponentID).(*components.StateComponent)

			if stateComponent.CurrentState != stateComponent.NextState {
				// Perform any cleanup for the old state here

				// Transition to the next state
				stateComponent.CurrentState = stateComponent.NextState

				// Perform any initialization for the new state here
				println("Entity", entity.ID, "changed state to", stateComponent.CurrentState)
			}

			// Handle the current state logic
			switch stateComponent.CurrentState {
			case components.StateIdle:
				if entity.HasComponent(components.IdlePositionComponentID) {
					idlePosition := entity.GetComponent(components.IdlePositionComponentID).(*components.IdlePositionComponent)
					position := entity.GetComponent(components.PositionComponentID).(*components.PositionComponent)
					position.DesX, position.DesY = idlePosition.TileX, idlePosition.TileY
				}
				continue
			case components.StateAttacking:
				HandleAttacking(entity, entitySlice)
			case components.StateGather:
				HandleGathering(entity, entitySlice)
			case components.StateDead:
				HandleDeath(entity)
			case components.StateWalkHere:
				continue
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
