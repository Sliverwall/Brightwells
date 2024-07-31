package components

const ResourceNodeComponentID = "ResourceNodeComponent"

type ResourceNodeComponent struct {
	// Mark type of resource node. Lumbering, fishing, etc...
	Type string
	// Map skill and level requirements skillName:level
	SkillsNeeded map[string]int
	// Active
	Active bool
	// Chance to set non-active
	DrainedChance float32
}

// Tag entity as able to gather
const GatherComponentID = "GatherComponent"

type GatherComponent struct {
	Target int // entity ID of thing being gathered
}
