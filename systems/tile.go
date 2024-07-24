package systems

import (
	"Brightwells/config"
	"Brightwells/entities"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type TileSystem struct {
	BackgroundMap [][]int
	ForegroundMap [][]int
	TileImages    map[int]*ebiten.Image
}

func (ts *TileSystem) InitializeTiles() ([]*entities.Entity, []*entities.Entity) {
	var backgroundTiles []*entities.Entity
	var entitySlice []*entities.Entity

	// Initialize background tiles
	for y, row := range ts.BackgroundMap {
		for x, tile := range row {
			img, ok := ts.TileImages[tile]
			if !ok {
				continue
			}
			posX := math.Floor(float64(x) * config.TileSize)
			posY := math.Floor(float64(y) * config.TileSize)
			layer := 1
			entity := entities.NewTileEntity(posX, posY, img, layer)
			backgroundTiles = append(backgroundTiles, entity)
		}
	}

	// Initialize foreground entities
	for y, row := range ts.ForegroundMap {
		for x, tile := range row {
			img, ok := ts.TileImages[tile]
			if !ok {
				continue
			}
			posX := float64(x) * config.TileSize
			posY := float64(y) * config.TileSize
			var entity *entities.Entity
			layer := 2

			switch tile {
			case -1:
				entity = entities.NewPlayer(posX, posY, 0, 0, img, layer)
			case 1:
				entity = entities.NewNPC(posX, posY, 0, 0, img, layer)
			case 2:
				entity = entities.NewApple(posX, posY, img, layer)
			default:
				continue
			}
			entitySlice = append(entitySlice, entity)
		}
	}
	return backgroundTiles, entitySlice
}

func LoadTiles() map[int]*ebiten.Image {
	playerSprite, _, err := ebitenutil.NewImageFromFile("assets/images/eggBoy.png")
	if err != nil {
		log.Fatal(err)
	}

	npcSprite, _, err := ebitenutil.NewImageFromFile("assets/images/caveGirl.png")
	if err != nil {
		log.Fatal(err)
	}

	appleSprite, _, err := ebitenutil.NewImageFromFile("assets/images/apple.png")
	if err != nil {
		log.Fatal(err)
	}

	orangeFloorSprite, _, err := ebitenutil.NewImageFromFile("assets/images/TilesetField.png")
	if err != nil {
		log.Fatal(err)
	}
	// Map tile types to their corresponding images
	tileImages := map[int]*ebiten.Image{
		-1: playerSprite,
		1:  npcSprite,
		2:  appleSprite,
		0:  orangeFloorSprite,
		// Add more tile types and their images here
	}

	return tileImages
}
