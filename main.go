package main

import (
	"Brightwells/components"
	"Brightwells/config"
	"Brightwells/data"
	"Brightwells/entities"
	"Brightwells/systems"
	"fmt"
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
	username               string
}

func (g *Game) Update() error {
	switch g.state {
	case LoginState:
		// Update login screen
		username, loginStatus := g.loginSystem.Update()
		if loginStatus {
			// Transition to game state if login is successful
			g.username = username
			g.initializeGameData() // Load game data after login
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
			// Perform updates before closing to save game
			position := g.player.GetComponent(components.PositionComponentID).(*components.PositionComponent)
			// Save player position
			savePlayerPositionQuery := fmt.Sprintf("UPDATE Player SET x = %f, y = %f WHERE name = '%s';", position.TileX, position.TileY, g.username)
			log.Println(position.TileX, ",", position.TileY)
			data.SQL_exec(savePlayerPositionQuery)
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
		// -----------UI ELEMENTS---------
		// Right click options
		if systems.RightClickTriggerOptions != nil {
			systems.DrawRightClickOptions(screen, systems.RightClickTriggerOptions, systems.RightClickX, systems.RightClickY)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	resolutionWidth := config.RESOLUTION_WIDTH
	resolutionHeight := config.RESOLUTION_HEIGHT

	return resolutionWidth, resolutionHeight
}

func (g *Game) initializeGameData() {
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
	g.backgroundTiles, g.entitySlice = tileSystem.InitializeTiles()

	// Initialize the player
	g.player = entities.GetPlayerEntity(g.entitySlice)

	// Initialize game systems
	g.collisionSystem = &systems.CollisionSystem{}
	g.triggerCollisionSystem = &systems.TriggerCollisionSystem{
		CollisionSystem: g.collisionSystem,
	}
	g.drawSystem = &systems.DrawSystem{}
}

func main() {
	loginSystem := systems.NewLoginSystem()

	game := &Game{
		state:       LoginState,
		loginSystem: loginSystem,
		tickManager: systems.NewTickManager(300 * time.Millisecond),
	}

	ebiten.SetWindowSize(config.WINDOW_WIDTH, config.WINDOW_HEIGHT)
	ebiten.SetWindowTitle("Brightwells")

	// Set a fixed frame rate of 60 FPS
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
