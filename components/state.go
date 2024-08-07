package components

// Define possible states
const (
	StateIdle = iota
	StateAttacking
	StateGather
	StateDead
	StateWalkHere
)

// StateComponentID is the identifier for the StateComponent
const StateComponentID = "StateComponent"

type StateComponent struct {
	CurrentState int
	NextState    int
}
