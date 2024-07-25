package components

const CameraComponentID = "CameraComponentID"

type CameraComponent struct {
	X, Y float64
}

// PlayerComponentID is the identifier for the PlayerComponent
const PlayerComponentID = "PlayerComponent"

type PlayerComponent struct {
	// flags entity as player
}

const SkillsComponentID = "SkillsComponent"

type SkillsComponent struct {
	// Skill section
	CurrentHealth int // keeps track of health that can be damaged
	Health        int
	Attack        int
	Strength      int
	Defence       int
	Range         int
	Prayer        int
	Magic         int
	Woodcutting   int
}
