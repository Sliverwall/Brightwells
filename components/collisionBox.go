package components

const CollisionBoxID = "CollisionBox"

type CollisionBox struct {
	*PositionComponent
}

// BoundingBox returns the bounds of the collision box in tile coordinates
func (cb *CollisionBox) BoundingBox() (float64, float64, float64, float64) {
	// Convert position to tile-based coordinates
	tileX := cb.PositionComponent.TileX
	tileY := cb.PositionComponent.TileY
	width := 1.0  // Adjust as needed
	height := 1.0 // Adjust as needed

	// Define the bounding box based on tile coordinates
	x1 := tileX
	y1 := tileY
	x2 := tileX + width
	y2 := tileY + height

	return x1, y1, x2, y2
}
