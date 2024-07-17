package main

import (
	"Brightwells/entities"
	"Brightwells/initialize"
	"Brightwells/systems"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	HEIGHT int = 240
	WIDTH  int = 320
)

type Game struct {
	entities        []*entities.Entity
	movementSystem  *systems.MovementSystem
	drawSystem      *systems.DrawSystem
	collisionSystem *systems.CollisionSystem
	userInputSystem *systems.UserInputSystem
}

func (g *Game) Update() error {
	g.movementSystem.Update(g.entities)
	g.collisionSystem.Update(g.entities)
	g.userInputSystem.Update(g.entities)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Clear the screen with black color
	screen.Fill(color.Black)
	// Draw the entities using the draw system
	g.drawSystem.Update(g.entities, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WIDTH, HEIGHT
}

func main() {
	entities := initialize.InitializeEntities()
	movementSystem, drawSystem, _, userInputSystem := initialize.InitializeSystems()

	game := &Game{
		entities:        entities,
		movementSystem:  movementSystem,
		drawSystem:      drawSystem,
		userInputSystem: userInputSystem,
	}
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Green Square Moving Back and Forth")
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
