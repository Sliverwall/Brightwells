package entities

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Initialize entites using unique ID

func SpawnEntity(npcID int, posX, posY float64, img *ebiten.Image, layer int) *Entity {

	// Use NewEntity functions to spawn entites
	var entity *Entity
	switch npcID {
	case 1:
		entity = NewPlayer(posX, posY, img, layer)
	case 2:
		entity = NewMonsterGirl(posX, posY, img, layer)
	}

	return entity
}

// Map sprite_id to sprite
func LoadSprites() map[int]*ebiten.Image {
	player_default, _, err := ebitenutil.NewImageFromFile("assets/images/eggBoy.png")
	if err != nil {
		log.Fatal(err)
	}

	monsterGirl, _, err := ebitenutil.NewImageFromFile("assets/images/caveGirl.png")
	if err != nil {
		log.Fatal(err)
	}

	appleSprite, _, err := ebitenutil.NewImageFromFile("assets/images/apple.png")
	if err != nil {
		log.Fatal(err)
	}

	// Map sprite_id to their corresponding images
	spriteImages := map[int]*ebiten.Image{
		1: player_default,
		2: monsterGirl,
		3: appleSprite,
		// Add more tile types and their images here
	}

	return spriteImages
}
