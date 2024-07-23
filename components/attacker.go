package components

const AttackerComponentID = "AttackerComponent"

// Component to tag unit as able to attack
type AttackerComponent struct {
	IsAttacking bool
	Target      int // entity ID of thing being attacked
}
