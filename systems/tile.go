package systems

import (
	"Brightwells/config"
	"Brightwells/entities"
	"encoding/csv"
	"image"
	"log"
	"os"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// ------------------------------ TILE SYSTEMS -------------------------------

type TileSystem struct {
	BackgroundMap [][]int
	ForegroundMap [][]interface{}
	SpriteImages  map[int]*ebiten.Image
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
			posX := float64(x)
			posY := float64(y)
			// Set 0 to be background layer. Draw system will not subdivides theses
			layer := 0

			// Create entity
			entity := entities.NewTileEntity(posX, posY, img, layer)

			// Append new tile entity to backgroundTiles set
			backgroundTiles = append(backgroundTiles, entity)
		}
	}

	// Initialize foreground entities
	for _, row := range ts.ForegroundMap {
		// Grab needed features from tuple. Type assertion as grabbing from interface{}
		npc_id := int(row[0].(int64))
		x := float64(row[2].(int64))
		y := float64(row[3].(int64))
		sprite_id := int(row[4].(int64))
		layer := int(row[5].(int64))

		img, ok := ts.SpriteImages[sprite_id]
		if !ok {
			continue
		}

		// Construct entity based on npc_id
		entity := entities.SpawnEntity(npc_id, x, y, img, layer)

		// Place entity in world
		entitySlice = append(entitySlice, entity)
	}
	return backgroundTiles, entitySlice
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
