package components

// Define possible AttackedResponses Types
const (
	AttackBack = iota
)

// AttackedResponseComponentID is the identifier for the AttackedResponseComponent
const AttackedResponseComponentID = "AttackedResponseComponent"

type AttackedResponseComponent struct {
	Type int
}
