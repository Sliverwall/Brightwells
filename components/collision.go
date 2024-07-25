package components

const CollisionComponentID = "CollisionComponent"

type CollisionComponent struct {
	// flag is detect collisions
}

const CollisionBoxID = "CollisionBox"

type CollisionBox struct {
	*PositionComponent
}
