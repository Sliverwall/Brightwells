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

	// Draw background tiles first
	ds.drawEntities(backgroundTiles, screen)

	// Sort entities by layer
	sort.Slice(entitySlice, func(i, j int) bool {
		return entitySlice[i].RenderLayer < entitySlice[j].RenderLayer
	})

	// Draw entities
	ds.drawEntities(entitySlice, screen)
}

func (ds *DrawSystem) drawEntities(entitySlice []*entities.Entity, screen *ebiten.Image) {
	for _, entity := range entitySlice {
		if entity.HasComponent(components.PositionComponentID) && entity.HasComponent(components.SpriteComponentID) {
			position := entity.GetComponent(components.PositionComponentID).(*components.PositionComponent)
			sprite := entity.GetComponent(components.SpriteComponentID).(*components.SpriteComponent)

			// Calculate the actual position on the screen
			actualX := math.Floor(position.TileX * config.TileSize)
			actualY := math.Floor(position.TileY * config.TileSize)
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
