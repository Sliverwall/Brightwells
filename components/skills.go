package components

const SkillsComponentID = "SkillsComponent"

type SkillsComponent struct {
	// Skill section
	CurrentHelath int // keeps track of health that can be damaged
	Health        int
	Attack        int
	Strength      int
	Defence       int
	Range         int
	Prayer        int
	Magic         int
	Woodcutting   int
}
