package components

// PositionComponentID is the identifier for the PositionComponent
const PositionComponentID = "PositionComponent"

type PositionComponent struct {
	TileX float64
	TileY float64
}

const DestinationComponentID = "DestinationComponent"

type DestinationComponent struct {
	X float64
	Y float64
}

// VelocityComponentID is the identifier for the VelocityComponent
const VelocityComponentID = "VelocityComponent"

type VelocityComponent struct {
	VX, VY float64
}
