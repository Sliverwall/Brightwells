package components

const AttackerComponentID = "AttackerComponent"

// Component to tag unit as able to attack
type AttackerComponent struct {
	Target int // entity ID of thing being attacked
}

const DamageComponentID = "DamageComponent"

type DamageComponent struct {
	// Component to tag unit as damageable
	AttackerID int // entity ID of thing attacking
}

// Define possible MonsterTypes
const (
	Human = iota
	Undead
	Dragon
)

// MonsterTypeComponentID is the identifier for the MonsterTypeComponent
const MonsterTypeComponentID = "MonsterTypeComponent"

type MonsterTypeComponent struct {
	Type int
}
