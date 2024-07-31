package main

import (
	"Brightwells/config"
	"Brightwells/data"
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

	stateSystem     *systems.StateSystem
	movementSystem  *systems.MovementSystem
	drawSystem      *systems.DrawSystem
	userInputSystem *systems.UserInputSystem
	deathSystem     systems.DeathSystem
	cameraSystem    *systems.CameraSystem
	tickManager     *systems.TickManager
}

func (g *Game) Update() error {

	// Update entites

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
		g.triggerCollisionSystem.Update(g.entitySlice)
		g.deathSystem.Update(g.entitySlice)
		// Update state
		g.stateSystem.Update(g.entitySlice)
		// Reload onl yexisting entities
		g.entitySlice = entities.GetExistEntitySlice(g.entitySlice)
	}

	// Always update camera for smooth following
	g.cameraSystem.Update(g.player)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Clear the screen
	screen.Clear()
	// Set primer background
	screen.Fill(color.Black)

	// Draw the background tiles and entities using the draw system
	g.drawSystem.Update(g.backgroundTiles, g.entitySlice, screen, g.player)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	resolutionWidth := config.RESOLUTION_WIDTH
	resolutionHeight := config.RESOLUTION_HEIGHT

	return resolutionWidth, resolutionHeight
}

func main() {
	// Load Entity Sprites
	spriteMap := data.SQL_query(data.Select_all_Sprite)
	spriteImages := systems.LoadSprites(spriteMap)

	// Load Background map
	backgroundMap := systems.ReadMap("assets/maps/map_1.csv")

	// Load Entity map
	entityMap := data.SQL_query(data.Select_all_Entity) // Any non-tile

	// Initialize background tiles and entities
	tileSystem := &systems.TileSystem{
		BackgroundMap: backgroundMap,
		EntityMap:     entityMap,
		SpriteImages:  spriteImages,
	}
	backgroundTiles, entitySlice := tileSystem.InitializeTiles()

	// Init systems
	collisionSystem := &systems.CollisionSystem{}

	foodRespawnSystem := &systems.FoodRespawnSystem{}

	triggerCollisionSystem := &systems.TriggerCollisionSystem{
		FoodRespawnSystem: foodRespawnSystem,
		CollisionSystem:   collisionSystem,
	}

	// Load player in
	player := entities.GetPlayerEntity(entitySlice)
	game := &Game{
		player:                 player,
		backgroundTiles:        backgroundTiles,
		entitySlice:            entitySlice,
		triggerCollisionSystem: triggerCollisionSystem,
		collisionSystem:        collisionSystem,
		tickManager:            systems.NewTickManager(300 * time.Millisecond),
	}

	ebiten.SetWindowSize(config.WINDOW_WIDTH, config.WINDOW_HEIGHT)
	ebiten.SetWindowTitle("Brightwells")

	// Set a fixed frame rate of 60 FPS

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
