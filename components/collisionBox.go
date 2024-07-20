package components

const CollisionBoxID = "CollisionBox"

type CollisionBox struct {
	*PositionComponent
	*CollisionComponent
}

// BoundingBox returns the bounds of the collision box
// Add switch case to chane bounding box based on direction of entity.
func (cb *CollisionBox) BoundingBox(direction int) (float64, float64, float64, float64) {
	x1 := cb.PositionComponent.X
	y1 := cb.PositionComponent.Y + cb.CollisionComponent.Height/2 // Start from the bottom half
	x2 := x1 + cb.CollisionComponent.Width
	y2 := y1 + cb.CollisionComponent.Height // Full height of the collision box
	return x1, y1, x2, y2
}
