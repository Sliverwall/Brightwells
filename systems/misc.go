package systems

import (
	"Brightwells/components"
	"Brightwells/entities"
	"math"
)

// Module to hold useful functions not tied to any one system

func CheckTileForEntity(tileX, tileY float64, entitySlice []*entities.Entity) int {
	for _, entity := range entitySlice {
		if !entity.HasComponent(components.PositionComponentID) {
			continue
		}
		position := entity.GetComponent(components.PositionComponentID).(*components.PositionComponent)
		// Check if entity is on the tile being checked
		if position.TileX == tileX && position.TileY == tileY {
			// Grab the id then exit function
			return entity.ID
		}
	}
	// Return -1 if no entity is on tile
	return -1
}

// IsTileOccupiedByColidableEntity method to check if a tile is occupied by a collidable entity
func IsTileOccupiedByCollidableEntity(tileX, tileY float64, entitySlice []*entities.Entity) bool {
	for _, entity := range entitySlice {
		if !entity.HasComponent(components.CollisionBoxID) || !entity.HasComponent(components.PositionComponentID) {
			continue
		}
		position := entity.GetComponent(components.PositionComponentID).(*components.PositionComponent)
		if position.TileX == tileX && position.TileY == tileY {
			return true
		}
	}
	return false
}

// IsWithinOneTile checks if two entities are within one tile of eachother
func IsWithinOneTile(entity1, entity2 *entities.Entity) bool {
	position1 := entity1.GetComponent(components.PositionComponentID).(*components.PositionComponent)
	position2 := entity2.GetComponent(components.PositionComponentID).(*components.PositionComponent)

	deltaX := math.Abs(position1.TileX - position2.TileX)
	deltaY := math.Abs(position1.TileY - position2.TileY)

	// Check if the entities are within one tile of each other in any direction
	return deltaX <= 1 && deltaY <= 1
}

// IsWithinManyTile checks if two entities are within x tile of eachother
func IsWithinManyTile(entity1, entity2 *entities.Entity, limit float64) bool {
	position1 := entity1.GetComponent(components.PositionComponentID).(*components.PositionComponent)
	position2 := entity2.GetComponent(components.PositionComponentID).(*components.PositionComponent)

	deltaX := math.Abs(position1.TileX - position2.TileX)
	deltaY := math.Abs(position1.TileY - position2.TileY)

	// Check if the entities are within one tile of each other in any direction
	return deltaX >= limit && deltaY >= limit
}

// ActivateRightClick handles right click input action for num pad user input
func ActivateRightClick(option, checkID int, entity *entities.Entity, entitySlice []*entities.Entity, attaker *components.AttackerComponent, gather *components.GatherComponent, position *components.PositionComponent) {
	switch option {
	case components.StateAttacking:
		attaker.Target = checkID
		SetNextState(entity, components.StateAttacking)
	case components.StateGather:
		gather.Target = checkID
		SetNextState(entity, components.StateGather)
	case components.StateWalkHere:
		SetNextState(entity, components.StateWalkHere)
		// use the id to grab entity data
		targetEntity := entities.GetEntityByID(entitySlice, checkID)
		targetPosition := targetEntity.GetComponent(components.PositionComponentID).(*components.PositionComponent)

		// Set desitination to the target
		position.DesX = targetPosition.TileX
		position.DesY = targetPosition.TileY

	}
	// remove options so print goes away
	RightClickTriggerOptions = nil
}
