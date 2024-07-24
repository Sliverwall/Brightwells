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

type DrawSystem struct{}

func (ds *DrawSystem) Update(backgroundTiles []*entities.Entity, entitySlice []*entities.Entity, screen *ebiten.Image) {

	var camera *components.CameraComponent
	// Find the camera entity
	for _, entity := range entitySlice {
		if entity.HasComponent(components.CameraComponentID) {
			camera = entity.GetComponent(components.CameraComponentID).(*components.CameraComponent)
			break
		}
	}
	// Draw background tiles first
	ds.drawEntities(backgroundTiles, screen, camera, -1)

	// Sort entities by layer
	sort.Slice(entitySlice, func(i, j int) bool {
		return entitySlice[i].RenderLayer < entitySlice[j].RenderLayer
	})

	// Draw entities
	ds.drawEntities(entitySlice, screen, camera, 0)
}

func (ds *DrawSystem) drawEntities(entitySlice []*entities.Entity, screen *ebiten.Image, camera *components.CameraComponent, layer int) {
	for _, entity := range entitySlice {
		if entity.HasComponent(components.PositionComponentID) && entity.HasComponent(components.SpriteComponentID) && (entity.RenderLayer == layer || layer == -1) {
			position := entity.GetComponent(components.PositionComponentID).(*components.PositionComponent)
			sprite := entity.GetComponent(components.SpriteComponentID).(*components.SpriteComponent)

			// Calculate the actual position on the screen
			actualX := math.Floor(position.TileX*config.TileSize) - camera.X
			actualY := math.Floor(position.TileY*config.TileSize) - camera.Y

			// Set the translation of the drawImage
			opts := ebiten.DrawImageOptions{}
			opts.GeoM.Translate(float64(actualX), actualY)

			// Draw the image, ensuring the correct sub-image is selected
			subImageRect := image.Rect(
				sprite.X,
				sprite.Y,
				sprite.X+sprite.X1,
				sprite.Y+sprite.Y1,
			)
			subImage := sprite.Image.SubImage(subImageRect).(*ebiten.Image)

			screen.DrawImage(subImage, &opts)
		}
	}
}
