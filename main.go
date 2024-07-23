package main

import (
	"Brightwells/config"
	"Brightwells/entities"
	"Brightwells/state"
	"Brightwells/systems"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	// game ECS init
	backgroundTiles []*entities.Entity
	entitySlice     []*entities.Entity

	collisionSystem        *systems.CollisionSystem
	triggerCollisionSystem *systems.TriggerCollisionSystem

	movementSystem  *systems.MovementSystem
	drawSystem      *systems.DrawSystem
	userInputSystem *systems.UserInputSystem
	deathSystem     systems.DeathSystem
	damageSystem    *systems.DamageSystem
}

func (g *Game) Update() error {

	// update existing entities
	g.entitySlice = entities.GetExistEntitySlice(g.entitySlice)
	g.damageSystem.Update(g.entitySlice)
	g.deathSystem.Update(g.entitySlice)

	g.triggerCollisionSystem.Update(g.entitySlice)

	g.movementSystem.Update(g.entitySlice)
	g.userInputSystem.Update(g.entitySlice)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Clear the screen with black color
	screen.Fill(color.Opaque)

	// Draw the background tiles and entities using the draw system
	g.drawSystem.Update(g.backgroundTiles, g.entitySlice, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	resolutionWidth := config.RESOLUTION_WIDTH
	resolutionHeight := config.RESOLUTION_HEIGHT

	return resolutionWidth, resolutionHeight
}

func main() {

	// Set world state
	worldInstance := state.WorldInstance
	tileImages := systems.LoadTiles()

	// Background tile map
	backgroundMap := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	// Foreground entity map
	foregroundMap := [][]int{
		{-1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 2, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	tileSystem := &systems.TileSystem{
		BackgroundMap: backgroundMap,
		ForegroundMap: foregroundMap,
		TileImages:    tileImages,
	}

	// Initialize background tiles and entities
	backgroundTiles, entitySlice := tileSystem.InitializeTiles()
	windowWidth := config.WINDOW_WIDTH
	windowHeight := config.WINDOW_HEIGHT

	// Init systems
	collisionSystem := &systems.CollisionSystem{
		GameMap: foregroundMap,
	}
	foodRespawnSystem := &systems.FoodRespawnSystem{}

	triggerCollisionSystem := &systems.TriggerCollisionSystem{
		FoodRespawnSystem: foodRespawnSystem,
		CollisionSystem:   collisionSystem,
	}

	movementSystem := &systems.MovementSystem{
		CollisionSystem: collisionSystem,
		WorldInstance:   *worldInstance,
	}
	damageSystem := &systems.DamageSystem{
		WorldInstance: *worldInstance,
	}
	drawSystem := &systems.DrawSystem{}
	userInputSystem := &systems.UserInputSystem{}

	game := &Game{
		backgroundTiles:        backgroundTiles,
		entitySlice:            entitySlice,
		triggerCollisionSystem: triggerCollisionSystem,
		collisionSystem:        collisionSystem,
		movementSystem:         movementSystem,
		drawSystem:             drawSystem,
		userInputSystem:        userInputSystem,
		damageSystem:           damageSystem,
	}
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Brightwells")

	// Set a fixed frame rate of 60 FPS

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
