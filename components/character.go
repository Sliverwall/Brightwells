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
	Health_exp    int

	CurrentMelee int
	Melee        int
	Melee_exp    int

	CurrentDefense int
	Defense        int
	Defense_exp    int

	CurrentArchery int
	Archery        int
	Archery_exp    int

	CurrentDevotion int
	Devotion        int
	Devotion_exp    int

	CurrentMagicka int
	Magicka        int
	Magicka_exp    int

	CurrentLumbering int
	Lumbering        int
	Lumbering_exp    int

	CurrentAlchemy int
	Alchemy        int
	Alchemy_exp    int

	CurrentHunter int
	Hunter        int
	Hunter_exp    int
}
