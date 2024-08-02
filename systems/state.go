package systems

import (
	"Brightwells/components"
	"Brightwells/entities"
)

// ------------------------------ STATE MACHINE SYSTEMS -------------------------------
type StateSystem struct{}

func (ss *StateSystem) Update(entitySlice []*entities.Entity) {
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
				ss.HandleAttacking(entity, entitySlice)
			case components.StateGather:
				ss.HandleGathering(entity, entitySlice)
			}
		}
	}
}
