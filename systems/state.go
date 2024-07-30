package systems

import (
	"Brightwells/components"
	"Brightwells/entities"
	"time"
)

type StateSystem struct{}

func (ss *StateSystem) Update(entitySlice []*entities.Entity) {
	for _, entity := range entitySlice {
		if entity.HasComponent(components.StateComponentID) {
			stateComponent := entity.GetComponent(components.StateComponentID).(*components.StateComponent)

			if stateComponent.CurrentState != stateComponent.NextState {
				// Perform any cleanup for the old state here

				// Transition to the next state
				stateComponent.CurrentState = stateComponent.NextState
				stateComponent.StateChanged = time.Now()

				// Perform any initialization for the new state here
				println("Entity", entity.ID, "changed state to", stateComponent.CurrentState)
			}

			// Handle the current state logic
			switch stateComponent.CurrentState {
			case components.StateIdle:
				ss.handleIdleState(entity)
			case components.StateMoving:
				ss.handleMovingState(entity)
			case components.StateAttacking:
				ss.handleAttackingState(entity)
			case components.StateDead:
				ss.handleDeadState(entity)
			}
		}
	}
}

func (ss *StateSystem) handleIdleState(entity *entities.Entity) {
	// Idle state logic here
}

func (ss *StateSystem) handleMovingState(entity *entities.Entity) {
	// Moving state logic here
}

func (ss *StateSystem) handleAttackingState(entity *entities.Entity) {
	// Attacking state logic here
}

func (ss *StateSystem) handleDeadState(entity *entities.Entity) {
	// Dead state logic here
}
