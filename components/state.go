package components

// Define possible states
const (
	StateIdle      = 0
	StateAttacking = 1
	StateGather    = 2
)

// StateComponentID is the identifier for the StateComponent
const StateComponentID = "StateComponent"

type StateComponent struct {
	CurrentState int
	NextState    int
}
