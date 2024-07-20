package main

import (
	"Brightwells/config"
	"Brightwells/entities"
	"Brightwells/systems"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	// game ECS init
	entitySlice []*entities.Entity

	collisionSystem        *systems.CollisionSystem
	triggerCollisionSystem *systems.TriggerCollisionSystem

	movementSystem  *systems.MovementSystem
	drawSystem      *systems.DrawSystem
	userInputSystem *systems.UserInputSystem
}

func (g *Game) Update() error {

	// update existing entities
	g.entitySlice = entities.GetExistEntitySlice(g.entitySlice)
	g.triggerCollisionSystem.Update(g.entitySlice)

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

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	resolutionWidth := config.ResolutionSize.Width
	resolutionHeight := config.ResolutionSize.Height

	return resolutionWidth, resolutionHeight
}

func main() {
	tileImages := systems.LoadTiles()

	gameMap := [][]int{
		{-1, 0, 0, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 2, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	tileSystem := &systems.TileSystem{
		GameMap:    gameMap,
		TileImages: tileImages,
	}

	entitySlice := tileSystem.InitializeTiles()
	windowWidth := config.WindowSize.Width
	windowHeight := config.WindowSize.Height

	// Init systems
	collisionSystem := &systems.CollisionSystem{
		GameMap: gameMap,
	}
	foodRespawnSystem := &systems.FoodRespawnSystem{}
	moveCollideSystem := systems.NewMoveCollideSystem()

	triggerCollisionSystem := &systems.TriggerCollisionSystem{
		FoodRespawnSystem: foodRespawnSystem,
		MoveCollideSystem: moveCollideSystem,
		CollisionSystem:   collisionSystem,
	}

	movementSystem := &systems.MovementSystem{
		MoveCollideSystem: moveCollideSystem,
	}
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
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Green Square Moving Back and Forth")

	// Set a fixed frame rate of 60 FPS

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
