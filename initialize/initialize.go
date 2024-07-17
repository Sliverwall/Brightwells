package initialize

import (
	"Brightwells/entities"
	"Brightwells/systems"
	"log"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// InitializeEntities creates and initializes entities with their components
func InitializeEntities() []*entities.Entity {
	// Load player sprite
	playerSprite, _, err := ebitenutil.NewImageFromFile("assets/images/eggBoy.png")
	if err != nil {
		log.Fatal(err)
	}

	npcSprite, _, err := ebitenutil.NewImageFromFile("assets/images/eggBoy.png")
	if err != nil {
		log.Fatal(err)
	}

	// Create entities
	npc1 := entities.NewNPC(50, 50, 0, 0, npcSprite)
	npc2 := entities.NewNPC(200, 100, 0, 0, npcSprite)

	player := entities.NewPlayer(0, 0, 0, 0, playerSprite)

	// Return the list containing all entities
	return []*entities.Entity{player, npc1, npc2}
}

// InitializeSystems creates and initializes systems
func InitializeSystems() (*systems.MovementSystem, *systems.DrawSystem, *systems.UserInputSystem) {
	return &systems.MovementSystem{}, &systems.DrawSystem{}, &systems.UserInputSystem{}
}
