package components

import "Brightwells/config"

const CollisionBoxID = "CollisionBox"

type CollisionBox struct {
	*PositionComponent
	*CollisionComponent
}

// BoundingBox returns the bounds of the collision box in tile coordinates
func (cb *CollisionBox) BoundingBox() (float64, float64, float64, float64) {
	// Convert position to tile-based coordinates
	tileX := cb.PositionComponent.TileX
	tileY := cb.PositionComponent.TileY
	width := cb.CollisionComponent.Width / config.TileSize
	height := cb.CollisionComponent.Height / config.TileSize

	// Define the bounding box based on tile coordinates
	x1 := tileX
	y1 := tileY
	x2 := x1 + width
	y2 := y1 + height

	return x1, y1, x2, y2
}
