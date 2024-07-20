package components

const CollisionBoxID = "CollisionBox"

type CollisionBox struct {
	*PositionComponent
	*CollisionComponent
}
