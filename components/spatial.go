package components

// PositionComponentID is the identifier for the PositionComponent
const PositionComponentID = "PositionComponent"

type PositionComponent struct {
	TileX float64
	TileY float64
	DesX  float64
	DesY  float64
}

// VelocityComponentID is the identifier for the VelocityComponent
const VelocityComponentID = "VelocityComponent"

type VelocityComponent struct {
	VX, VY float64
}

// SpawnPointComponentID is the identifier for the SpawnPointComponent
const SpawnPointComponentID = "SpawnPointComponent"

type SpawnPointComponent struct {
	TileX, TileY float64
}

// IdlePositionComponentID is the identifier for the IdlePositionComponent
const IdlePositionComponentID = "IdlePositionComponent"

type IdlePositionComponent struct {
	TileX, TileY float64
}
