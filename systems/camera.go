package systems

import (
	"Brightwells/components"
	"Brightwells/config"
	"Brightwells/entities"
)

type CameraSystem struct{}

func (cs *CameraSystem) Update(entitySlice []*entities.Entity) {
	var player *entities.Entity
	var camera *entities.Entity

	for _, entity := range entitySlice {
		if entity.HasComponent(components.PlayerComponentID) {
			player = entity
		}
		if entity.HasComponent(components.CameraComponentID) {
			camera = entity
		}
	}

	if player != nil && camera != nil {
		playerPosition := player.GetComponent(components.PositionComponentID).(*components.PositionComponent)
		cameraComponent := camera.GetComponent(components.CameraComponentID).(*components.CameraComponent)

		cameraComponent.X = playerPosition.TileX*config.TileSize - cameraComponent.Width/2
		cameraComponent.Y = playerPosition.TileY*config.TileSize - cameraComponent.Height/2
	}
}
