package systems

import (
	"Brightwells/components"
	"Brightwells/entities"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type DrawSystem struct{}

func (ds *DrawSystem) Update(entities []*entities.Entity, screen *ebiten.Image) {
	for _, entity := range entities {
		if entity.HasComponent(components.PositionComponentID) && entity.HasComponent(components.SpriteComponentID) {
			position := entity.GetComponent(components.PositionComponentID).(*components.PositionComponent)
			sprite := entity.GetComponent(components.SpriteComponentID).(*components.SpriteComponent)

			// set the translation of the drawImage
			opts := ebiten.DrawImageOptions{}
			opts.GeoM.Translate(float64(position.X), float64(position.Y))

			screen.DrawImage(
				sprite.Image.SubImage(
					image.Rect(sprite.X, sprite.Y, sprite.X1, sprite.Y1),
				).(*ebiten.Image),
				&opts,
			)

			// clear geom position for next sprite
			opts.GeoM.Reset()
		}
	}
}
