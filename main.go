package main

import (
	"Brightwells/config"
	"Brightwells/data"
	"Brightwells/entities"
	"Brightwells/systems"
	"image/color"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

// Define game states
const (
	LoginState = iota
	GameState
)

type Game struct {
	state                  int
	backgroundTiles        []*entities.Entity
	entitySlice            []*entities.Entity
	deadEntitySlice        []*entities.Entity
	player                 *entities.Entity
	collisionSystem        *systems.CollisionSystem
	triggerCollisionSystem *systems.TriggerCollisionSystem
	drawSystem             *systems.DrawSystem
	userInputSystem        *systems.UserInputSystem
	tickManager            *systems.TickManager
	loginSystem            *systems.LoginSystem
}

func (g *Game) Update() error {
	switch g.state {
	case LoginState:
		// Update login screen
		if g.loginSystem.Update() {
			// Transition to game state if login is successful
			g.state = GameState
		}
	case GameState:
		// Always update user input for responsiveness
		g.userInputSystem.Update(g.player, g.entitySlice)

		// Check if it's time for a tick update
		if g.tickManager.ShouldUpdate() {
			// Update systems that should be tick-based
			systems.UpdateMovement(g.entitySlice)
			g.triggerCollisionSystem.Update(g.entitySlice)
			// Check if entities hit 0 hp
			systems.UpdateCheckZeroHP(g.entitySlice)
			systems.UpdateResourceTime(g.entitySlice)

			// Update state
			systems.UpdateState(g.entitySlice)
			// Reload only existing entities. Check Respawn system here
			g.entitySlice, g.deadEntitySlice = entities.GetExistEntitySlice(g.entitySlice, g.deadEntitySlice)
		}

		// Always update camera for smooth following
		systems.UpdateCamera(g.player)

		// Check for game exit
		// Check if the window is being closed
		if ebiten.IsWindowBeingClosed() {
			log.Println("Window close event triggered. Performing cleanup.")
			// Perform any necessary cleanup before exiting
			// ...
		}

	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Clear the screen
	screen.Clear()
	// Set primer background
	screen.Fill(color.Black)

	switch g.state {
	case LoginState:
		// Draw login screen
		g.loginSystem.Draw(screen)
	case GameState:
		// Draw the background tiles and entities using the draw system
		g.drawSystem.Update(g.backgroundTiles, g.entitySlice, screen, g.player)
	}
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
	triggerCollisionSystem := &systems.TriggerCollisionSystem{
		CollisionSystem: collisionSystem,
	}
	loginSystem := systems.NewLoginSystem()

	// Load player in
	player := entities.GetPlayerEntity(entitySlice)
	game := &Game{
		state:                  LoginState,
		player:                 player,
		backgroundTiles:        backgroundTiles,
		entitySlice:            entitySlice,
		triggerCollisionSystem: triggerCollisionSystem,
		collisionSystem:        collisionSystem,
		tickManager:            systems.NewTickManager(300 * time.Millisecond),
		loginSystem:            loginSystem,
	}

	ebiten.SetWindowSize(config.WINDOW_WIDTH, config.WINDOW_HEIGHT)
	ebiten.SetWindowTitle("Brightwells")

	// Set a fixed frame rate of 60 FPS
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
