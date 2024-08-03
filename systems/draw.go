package systems

import (
	"Brightwells/components"
	"Brightwells/config"
	"Brightwells/entities"
	"image"
	"math"
	"sort"

	"github.com/hajimehoshi/ebiten/v2"
)

// ------------------------------ DRAW SYSTEMS -------------------------------
type DrawSystem struct{}

func (ds *DrawSystem) Update(backgroundTiles []*entities.Entity, entitySlice []*entities.Entity, screen *ebiten.Image, player *entities.Entity) {

	// Draw background tiles first
	ds.drawEntities(backgroundTiles, screen, player)

	// Sort entities by layer
	sort.Slice(entitySlice, func(i, j int) bool {
		return entitySlice[i].RenderLayer < entitySlice[j].RenderLayer
	})

	// Draw entities
	ds.drawEntities(entitySlice, screen, player)
}

func (ds *DrawSystem) drawEntities(entitySlice []*entities.Entity, screen *ebiten.Image, player *entities.Entity) {
	for _, entity := range entitySlice {
		if entity.HasComponent(components.PositionComponentID) && entity.HasComponent(components.SpriteComponentID) {
			// Get entity compontents
			position := entity.GetComponent(components.PositionComponentID).(*components.PositionComponent)
			sprite := entity.GetComponent(components.SpriteComponentID).(*components.SpriteComponent)

			// Get camera compontents
			camera := player.GetComponent(components.CameraComponentID).(*components.CameraComponent)

			// Calculate the actual position on the screen
			actualX := math.Floor(position.TileX*config.TileSize - camera.X)
			actualY := math.Floor(position.TileY*config.TileSize - camera.Y)
			// Set the translation of the drawImage
			opts := ebiten.DrawImageOptions{}
			opts.GeoM.Translate(actualX, actualY)

			// Use layer to determine if entity is background tile
			if entity.RenderLayer == 0 {
				// If background til, do not subdivide image.
				screen.DrawImage(sprite.Image, &opts)
			} else {
				// Draw the image, ensuring the correct sub-image is selected
				subImageRect := image.Rect(
					sprite.X,
					sprite.Y,
					sprite.X1,
					sprite.Y1,
				)
				subImage := sprite.Image.SubImage(subImageRect).(*ebiten.Image)
				screen.DrawImage(subImage, &opts)
			}
		}
	}
}

// ------------------------------ CAMERA SYSTEMS -------------------------------
func UpdateCamera(player *entities.Entity) {

	if player != nil {
		playerPosition := player.GetComponent(components.PositionComponentID).(*components.PositionComponent)
		playerCamera := player.GetComponent(components.CameraComponentID).(*components.CameraComponent)

		playerCamera.X = (playerPosition.TileX*config.TileSize - float64(config.RESOLUTION_WIDTH)/2) + 10 // slight x-axis offset
		playerCamera.Y = (playerPosition.TileY*config.TileSize - float64(config.RESOLUTION_HEIGHT)/2) + 10
	}
}
