package systems

import (
	"Brightwells/config"
	"Brightwells/entities"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type TileSystem struct {
	GameMap    [][]int
	TileImages map[int]*ebiten.Image
}

func (ts *TileSystem) InitializeTiles() []*entities.Entity {
	var entitySlice []*entities.Entity

	for y, row := range ts.GameMap {
		for x, tile := range row {
			img, ok := ts.TileImages[tile]
			if !ok {
				// Handle case where tile type is not found
				continue
			}

			posX := float64(x) * config.TileSize
			posY := float64(y) * config.TileSize

			var entity *entities.Entity
			var layer int

			switch tile {
			case -1:
				// Initialize player entity
				layer = 1
				entity = entities.NewPlayer(posX, posY, 0, 0, img, layer)
			case 1:
				layer = 2
				entity = entities.NewNPC(posX, posY, -1, 0, img, layer)
			case 2:
				// Initialize regular tile entity
				layer = 1
				entity = entities.NewApple(posX, posY, img, layer)
			case 0:
				layer = -1
				entity = entities.NewTileEntity(posX, posY, img, layer)
			}

			entitySlice = append(entitySlice, entity)
		}
	}
	return entitySlice
}

func LoadTiles() map[int]*ebiten.Image {
	playerSprite, _, err := ebitenutil.NewImageFromFile("assets/images/eggBoy.png")
	if err != nil {
		log.Fatal(err)
	}

	npcSprite, _, err := ebitenutil.NewImageFromFile("assets/images/npc1.png")
	if err != nil {
		log.Fatal(err)
	}

	appleSprite, _, err := ebitenutil.NewImageFromFile("assets/images/apple.png")
	if err != nil {
		log.Fatal(err)
	}

	logicSprite, _, err := ebitenutil.NewImageFromFile("assets/images/TilesetLogic.png")
	if err != nil {
		log.Fatal(err)
	}
	// Map tile types to their corresponding images
	tileImages := map[int]*ebiten.Image{
		-1: playerSprite,
		1:  npcSprite,
		2:  appleSprite,
		0:  logicSprite,
		// Add more tile types and their images here
	}

	return tileImages
}
