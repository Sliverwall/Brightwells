package systems

import (
	"Brightwells/config"
	"Brightwells/entities"
	"encoding/csv"
	"image"
	"log"
	"math"
	"os"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// ------------------------------ TILE SYSTEMS -------------------------------

type TileSystem struct {
	BackgroundMap [][]int
	ForegroundMap [][]int
	TileImages    map[int]*ebiten.Image
}

func (ts *TileSystem) InitializeTiles() ([]*entities.Entity, []*entities.Entity) {
	var backgroundTiles []*entities.Entity
	var entitySlice []*entities.Entity

	backgroundImage, _, err := ebitenutil.NewImageFromFile("assets/images/TilesetFloor.png")
	if err != nil {
		log.Fatal(err)
	}
	// Initialize background tiles
	for y, row := range ts.BackgroundMap {
		for x, tileId := range row {
			// Get image from tileSet
			srcX, srcY, srcX1, srcY1 := GetSubImage(tileId, 22)

			// Subdivide tileset
			img := backgroundImage.SubImage(image.Rect(srcX, srcY, srcX1, srcY1)).(*ebiten.Image)

			// Get tile's position on map
			posX := math.Floor(float64(x) * config.TileSize)
			posY := math.Floor(float64(y) * config.TileSize)
			// Set 0 to be background layer. Draw system will not subdivides theses
			layer := 0

			// Create entity
			entity := entities.NewTileEntity(posX, posY, img, layer)

			// Append new tile entity to backgroundTiles set
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
				entity = entities.NewPlayer(posX, posY, img, layer)
			case 1:
				entity = entities.NewNPC(posX, posY, img, layer)
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
		18: orangeFloorSprite,
		// Add more tile types and their images here
	}

	return tileImages
}

func ReadMap(filePath string) [][]int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
		return nil
	}

	var result [][]int

	for _, record := range records {
		var intRecord []int
		for _, value := range record {
			intValue, err := strconv.Atoi(value)
			if err != nil {
				log.Fatal(err)
				return nil
			}
			intRecord = append(intRecord, intValue)
		}
		result = append(result, intRecord)
	}

	return result

}

// ------------------------------ PLACE ENTITY SYSTEMS -------------------------------

// ------------------------------ SPRITE FUNCTIONS -------------------------------

// GetSprite returns the Sprites subImage cords given the row and column in the tileset
func GetSprite(row, column int) (int, int, int, int) {
	// Declare cord variables

	x := (column * int(config.TileSize))
	y := (row * int(config.TileSize))
	x1 := x + int(config.TileSize)
	y1 := y + int(config.TileSize)

	return x, y, x1, y1
}

// GetSubImage returns the subImage cords given a tileId and the tileSets tilesPerRow
func GetSubImage(tileId, tilesPerRow int) (int, int, int, int) {
	tileRow := tileId / tilesPerRow
	tileColumn := tileId % tilesPerRow

	x, y, x1, y1 := GetSprite(tileRow, tileColumn)
	return x, y, x1, y1
}
