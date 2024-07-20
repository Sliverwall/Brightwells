package systems

import (
	"Brightwells/components"
	"Brightwells/entities"

	"github.com/hajimehoshi/ebiten/v2"
)

type MovementSystem struct {
	MoveCollideSystem *MoveCollideSystem
}

func (ms *MovementSystem) Update(entitySlice []*entities.Entity) {

	var fps float64
	if ebiten.ActualFPS() == 0 {
		fps = 60.0
	} else {
		fps = ebiten.ActualFPS()
	}
	dt := 1.0 / fps // Assume 60 FPS for movement calculations; adjust as needed

	for _, entity := range entitySlice {
		if entity.HasComponent(components.PositionComponentID) && entity.HasComponent(components.VelocityComponentID) {
			// Save the original position before moving
			ms.MoveCollideSystem.SaveOriginalPosition(entity)

			// Update position based on velocity
			position := entity.GetComponent(components.PositionComponentID).(*components.PositionComponent)
			velocity := entity.GetComponent(components.VelocityComponentID).(*components.VelocityComponent)

			// Calculate movement for this frame
			position.X += velocity.VX * dt
			position.Y += velocity.VY * dt
		}
	}
}
