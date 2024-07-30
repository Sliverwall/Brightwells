package entities

import (
	"Brightwells/components"

	"github.com/hajimehoshi/ebiten/v2"
)

func NewTileEntity(posX, posY float64, sprite *ebiten.Image, layer int) *Entity {
	entity := NewEntity(layer)
	entity.AddComponent(components.PositionComponentID, &components.PositionComponent{
		TileX: posX,
		TileY: posY,
	})
	entity.AddComponent(components.SpriteComponentID, &components.SpriteComponent{
		Image: sprite,
		X:     0, Y: 0, X1: 16, Y1: 16, // Initial sub-image coordinates
	})
	return entity
}
