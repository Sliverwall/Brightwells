package components

const CollisionComponentID = "CollisionComponent"

type CollisionComponent struct {
	Width  float64
	Height float64
}

func (cw *CollisionComponent) CollidesWith(currentPosition *PositionComponent, other *CollisionComponent, otherX, otherY float64) bool {
	// Calculate absolute positions
	left1 := currentPosition.X
	right1 := currentPosition.X + cw.Width
	top1 := currentPosition.Y
	bottom1 := currentPosition.Y + cw.Height

	left2 := otherX
	right2 := otherX + other.Width
	top2 := otherY
	bottom2 := otherY + other.Height

	// Check for collision
	return right1 > left2 &&
		left1 < right2 &&
		bottom1 > top2 &&
		top1 < bottom2
}
