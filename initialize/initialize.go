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

	// Create player with position and velocity
	player := entities.NewPlayer(100, 100, 2, 1, playerSprite)

	// Return the list containing the single player entity
	return []*entities.Entity{player}
}

// InitializeSystems creates and initializes systems
func InitializeSystems() (*systems.MovementSystem, *systems.DrawSystem, *systems.CollisionSystem, *systems.UserInputSystem) {
	return &systems.MovementSystem{}, &systems.DrawSystem{}, &systems.CollisionSystem{}, &systems.UserInputSystem{}
}
