package components

import "Brightwells/config"

const CollisionBoxID = "CollisionBox"

type CollisionBox struct {
	*PositionComponent
	*CollisionComponent
}

// BoundingBox returns the bounds of the collision box in tile coordinates
// Add switch case to change bounding box based on direction of entity.
func (cb *CollisionBox) BoundingBox(direction int) (float64, float64, float64, float64) {
	// Convert position to tile-based coordinates
	tileX := cb.PositionComponent.TileX
	tileY := cb.PositionComponent.TileY
	width := cb.CollisionComponent.Width / config.TileSize
	height := cb.CollisionComponent.Height / config.TileSize

	switch direction {
	case 0, 1, 3, 4: // right, left, down, idle
		// Moving right, left, down, or idle: the bounding box is as expected
		x1 := tileX
		y1 := tileY
		x2 := x1 + width
		y2 := y1 + height
		return x1, y1, x2, y2
	case 2: // up
		// Moving up: adjust the y1 position by subtracting the height
		x1 := tileX
		y1 := tileY - height
		x2 := x1 + width
		y2 := y1 + height
		return x1, y1, x2, y2
	}
	return 0, 0, 0, 0
}
