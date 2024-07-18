package main

import (
	"Brightwells/config"
	"Brightwells/entities"
	"Brightwells/initialize"
	"Brightwells/systems"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	// game ECS init
	entities          []*entities.Entity
	movementSystem    *systems.MovementSystem
	drawSystem        *systems.DrawSystem
	userInputSystem   *systems.UserInputSystem
	foodRespawnSystem *systems.FoodRespawnSystem
}

func (g *Game) Update() error {
	g.foodRespawnSystem.Update(g.entities)
	g.movementSystem.Update(g.entities)
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
	windowWidth := config.WindowSize.Width
	windowHight := config.WindowSize.Height
	return windowWidth, windowHight
}

func main() {

	windowWidth := config.WindowSize.Width
	windowHight := config.WindowSize.Height
	entities := initialize.InitializeEntities()

	// Init systems
	movementSystem := &systems.MovementSystem{}
	drawSystem := &systems.DrawSystem{}
	userInputSystem := &systems.UserInputSystem{}
	foodRespawnSystem := &systems.FoodRespawnSystem{}

	game := &Game{
		entities:          entities,
		movementSystem:    movementSystem,
		drawSystem:        drawSystem,
		userInputSystem:   userInputSystem,
		foodRespawnSystem: foodRespawnSystem,
	}
	ebiten.SetWindowSize(windowWidth, windowHight)
	ebiten.SetWindowTitle("Green Square Moving Back and Forth")
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
