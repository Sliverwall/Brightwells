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
	backgroundTiles []*entities.Entity
	entitySlice     []*entities.Entity
	player          *entities.Entity

	collisionSystem        *systems.CollisionSystem
	triggerCollisionSystem *systems.TriggerCollisionSystem

	movementSystem  *systems.MovementSystem
	drawSystem      *systems.DrawSystem
	userInputSystem *systems.UserInputSystem
	deathSystem     systems.DeathSystem
	damageSystem    *systems.DamageSystem
	cameraSystem    *systems.CameraSystem
	tickManager     *systems.TickManager
}

func (g *Game) Update() error {

	// Update entites
	g.entitySlice = entities.GetExistEntitySlice(g.entitySlice)
	// Always update user input for responsiveness
	g.userInputSystem.Update(g.entitySlice)

	// Get player entity (consider caching this if it doesn't change often)
	if g.player == nil {
		g.player = entities.GetPlayerEntity(g.entitySlice)
	}

	// Check if it's time for a tick update
	if g.tickManager.ShouldUpdate() {
		// Update systems that should be tick-based
		g.movementSystem.Update(g.entitySlice)
		g.damageSystem.Update(g.entitySlice)
		g.triggerCollisionSystem.Update(g.entitySlice)
		g.deathSystem.Update(g.entitySlice)
	}

	// Always update camera for smooth following
	g.cameraSystem.Update(g.player)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Clear the screen
	screen.Clear()
	// Set primer background
	screen.Fill(color.Opaque)

	// Draw the background tiles and entities using the draw system
	g.drawSystem.Update(g.backgroundTiles, g.entitySlice, screen, g.player)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	resolutionWidth := config.RESOLUTION_WIDTH
	resolutionHeight := config.RESOLUTION_HEIGHT

	return resolutionWidth, resolutionHeight
}

func main() {
	// Load tiles
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
		{0, 0, 0, 1, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 2, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 1, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 2, 2, 0, 0, 0, 0, 0},
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

	movementSystem := &systems.MovementSystem{}
	damageSystem := &systems.DamageSystem{}
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
		tickManager:            systems.NewTickManager(300 * time.Millisecond),
	}

	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Brightwells")

	// Set a fixed frame rate of 60 FPS

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
