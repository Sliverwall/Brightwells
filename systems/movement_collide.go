package systems

import (
	"Brightwells/components"
	"Brightwells/entities"
)

type MoveCollideSystem struct {
	originalPositions map[int]components.PositionComponent
}

// NewMoveCollideSystem initializes a MoveCollideSystem with a map to store original positions.
func NewMoveCollideSystem() *MoveCollideSystem {
	return &MoveCollideSystem{
		originalPositions: make(map[int]components.PositionComponent),
	}
}

// SaveOriginalPosition stores the original position of an entity before movement.
func (mcs *MoveCollideSystem) SaveOriginalPosition(entity *entities.Entity) {
	if entity.HasComponent(components.PositionComponentID) {
		position := entity.GetComponent(components.PositionComponentID).(*components.PositionComponent)
		mcs.originalPositions[entity.ID] = *position
	}
}

// RevertPosition restores the position of an entity to its original position.
func (mcs *MoveCollideSystem) RevertPosition(entity *entities.Entity) {
	if originalPos, ok := mcs.originalPositions[entity.ID]; ok {
		if position := entity.GetComponent(components.PositionComponentID); position != nil {
			*position.(*components.PositionComponent) = originalPos
		}
	}
}

// HandleCollisions processes collision results and reverts positions as necessary.
func (mcs *MoveCollideSystem) HandleCollisions(entitySlice []*entities.Entity, collisions map[int][]int) {
	for entityID, collidingEntities := range collisions {
		entity := entities.GetEntityByID(entitySlice, entityID)
		if entity != nil {
			// Revert the position of the current entity
			mcs.RevertPosition(entity)
		}

		for _, collidingID := range collidingEntities {
			collidingEntity := entities.GetEntityByID(entitySlice, collidingID)
			if collidingEntity != nil {
				// Revert the position of the colliding entity
				mcs.RevertPosition(collidingEntity)
			}
		}
	}
}
