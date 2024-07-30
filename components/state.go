package components

import (
	"time"
)

// Define possible states
const (
	StateIdle      = "Idle"
	StateMoving    = "Moving"
	StateAttacking = "Attacking"
	StateDead      = "Dead"
)

type StateComponent struct {
	CurrentState string
	NextState    string
	StateChanged time.Time
}

func NewStateComponent(initialState string) *StateComponent {
	return &StateComponent{
		CurrentState: initialState,
		NextState:    initialState,
		StateChanged: time.Now(),
	}
}
