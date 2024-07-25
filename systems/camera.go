package systems

import (
	"Brightwells/components"
	"Brightwells/config"
	"Brightwells/entities"
)

type CameraSystem struct{}

func (cs *CameraSystem) Update(player *entities.Entity) {

	if player != nil {
		playerPosition := player.GetComponent(components.PositionComponentID).(*components.PositionComponent)
		playerCamera := player.GetComponent(components.CameraComponentID).(*components.CameraComponent)

		playerCamera.X = playerPosition.TileX*config.TileSize - float64(config.RESOLUTION_WIDTH)/2
		playerCamera.Y = playerPosition.TileY*config.TileSize - float64(config.RESOLUTION_HEIGHT)/2
	}
}
