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
	entitySlice            []*entities.Entity
	collisionSystem        *systems.CollisionSystem
	triggerCollisionSystem *systems.TriggerCollisionSystem

	movementSystem  *systems.MovementSystem
	drawSystem      *systems.DrawSystem
	userInputSystem *systems.UserInputSystem
}

func (g *Game) Update() error {

	collisions := g.collisionSystem.Update(g.entitySlice)
	g.triggerCollisionSystem.Update(g.entitySlice, collisions)

	g.movementSystem.Update(g.entitySlice)
	g.userInputSystem.Update(g.entitySlice)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Clear the screen with black color
	screen.Fill(color.Black)
	// Draw the entities using the draw system
	g.drawSystem.Update(g.entitySlice, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	windowWidth := config.WindowSize.Width
	windowHight := config.WindowSize.Height
	return windowWidth, windowHight
}

func main() {

	windowWidth := config.WindowSize.Width
	windowHight := config.WindowSize.Height
	entitySlice := initialize.InitializeEntities()

	// Init systems
	collisionSystem := &systems.CollisionSystem{}
	triggerCollisionSystem := &systems.TriggerCollisionSystem{}

	movementSystem := &systems.MovementSystem{}
	drawSystem := &systems.DrawSystem{}
	userInputSystem := &systems.UserInputSystem{}

	game := &Game{
		entitySlice:            entitySlice,
		triggerCollisionSystem: triggerCollisionSystem,
		collisionSystem:        collisionSystem,
		movementSystem:         movementSystem,
		drawSystem:             drawSystem,
		userInputSystem:        userInputSystem,
	}
	ebiten.SetWindowSize(windowWidth, windowHight)
	ebiten.SetWindowTitle("Green Square Moving Back and Forth")
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
