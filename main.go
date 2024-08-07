package main

import (
	"Brightwells/config"
	"Brightwells/entities"
	"Brightwells/systems"
	"image/color"
	"time"

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
	screen.Fill(color.Opaque)
	// Draw the entities using the draw system
	g.drawSystem.Update(g.entitySlice, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	resolutionWidth := config.RESOLUTION_WIDTH
	resolutionHeight := config.RESOLUTION_HEIGHT

	return resolutionWidth, resolutionHeight
}

func main() {

	updateInterval := 300 * time.Millisecond // Update tick rate
	tileImages := systems.LoadTiles()

	gameMap := [][]int{
		{-1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
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
	windowWidth := config.WINDOW_WIDTH
	windowHeight := config.WINDOW_HEIGHT

	// Init systems
	collisionSystem := &systems.CollisionSystem{
		GameMap: gameMap,
	}
	foodRespawnSystem := &systems.FoodRespawnSystem{}

	triggerCollisionSystem := &systems.TriggerCollisionSystem{
		FoodRespawnSystem: foodRespawnSystem,
		CollisionSystem:   collisionSystem,
	}

	movementSystem := &systems.MovementSystem{
		CollisionSystem: collisionSystem,
		LastUpdateTime:  time.Now(),
		UpdateInterval:  updateInterval,
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
