package systems

import (
	"Brightwells/components"
	"Brightwells/entities"
	"image"
	"sort"

	"github.com/hajimehoshi/ebiten/v2"
)

type DrawSystem struct{}

func (ds *DrawSystem) Update(entitySlice []*entities.Entity, screen *ebiten.Image) {

	// Sort entities by layer
	sort.Slice(entitySlice, func(i, j int) bool {
		return entitySlice[i].RenderLayer < entitySlice[j].RenderLayer
	})
	for _, entity := range entitySlice {
		if entity.HasComponent(components.PositionComponentID) && entity.HasComponent(components.SpriteComponentID) {
			position := entity.GetComponent(components.PositionComponentID).(*components.PositionComponent)
			sprite := entity.GetComponent(components.SpriteComponentID).(*components.SpriteComponent)

			// Calculate the actual position on the screen
			actualX := position.X
			actualY := position.Y

			// Set the translation of the drawImage
			opts := ebiten.DrawImageOptions{}
			opts.GeoM.Translate(float64(actualX), float64(actualY))

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
