package engine

import (
	"fmt"
	"go-rogue/src/lib/components"
	"go-rogue/src/lib/config"
	"go-rogue/src/lib/entities"
	"go-rogue/src/lib/interfaces"
	"go-rogue/src/lib/maps"
	"go-rogue/src/lib/scenes"
	"go-rogue/src/lib/userInterface"
	"go-rogue/src/lib/utilities"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/eiannone/keyboard"
)

type GameScenes struct {
	scenes []scenes.IScene
}

func NewGameScenes() *GameScenes {
	return &GameScenes{
		scenes: []scenes.IScene{
			scenes.NewCabinScene(),
			scenes.NewForestEntranceScene(),
			scenes.NewForestEntrance2Scene(),
		},
	}
}

func (gs *GameScenes) GetScene(index int) scenes.IScene {
	if index < 0 || index >= len(gs.scenes) {
		return nil
	}
	return gs.scenes[index]
}

type Game struct {
	TickRate   float32
	Player     *entities.Player
	Enemy      interfaces.IEntity
	Combat     *components.Combat
	GameScenes *GameScenes
	SceneGraph *maps.SceneGraph
	World      *entities.World
}

func NewGame(player *entities.Player, enemy interfaces.IEntity, tickRate float32) *Game {
	return &Game{
		TickRate:   tickRate,
		Player:     player,
		Enemy:      enemy,
		Combat:     &components.Combat{TickRate: 0.5},
		GameScenes: NewGameScenes(),
		World:      entities.NewWorld(),
	}
}

func (g *Game) Run() {
	g.World.AddZone(0, 0, 0, 0, true)
	traversalProcessor := NewTraversalProcessor(&InputProcessor{})
	eventProcessor := NewEventProcessor()
	userInterface.DrawTitleText("Go Rogue")
	maps.WriteDotFile("graph.dot", g.World.GetCurrentZone().GetSceneGraph())

	ticker := time.NewTicker(time.Duration(g.TickRate*1000) * time.Millisecond)
	defer ticker.Stop()

	var input string
	if input == "exit" {
		os.Exit(0)
	}

	for range ticker.C {
		utilities.ClearScreenBelow(2, config.General.Offset)
		// Unpack the current node
		currentNode := g.World.GetCurrentZone().GetSceneGraph().GetNode(g.Player.GetCurrentPosition())

		currentLine := 3
		// Draw the event screen

		// Draw Traversal Options
		currentLine = eventProcessor.Execute(currentNode, g.Player, currentLine)
		currentLine = traversalProcessor.Execute(currentNode, g.Player, currentLine)

		currentLine += 2
		err := keyboard.Open()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("\033[%d;%dH%s\n", currentLine, config.General.Offset, "Press Enter to continue...")
		_, key, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		// Check if the Enter key is pressed
		if key == keyboard.KeyEnter {
			keyboard.Close()
			continue
		}
		// Resolve Traversal

	}
}

// Flow
// * Process input
// * Process movement
// * Present edge-traversal
// * Resolve edge-traversal
// * Present event
// * Process encounter | decision
// * Resolve event
// * Resolve turn

func (g *Game) TransitionBetweenZones() {
	forwardTraversal := true
	if g.Player.GetCurrentPosition() == g.World.GetCurrentZone().GetSceneGraph().GetTerminusNodeId() ||
		g.Player.GetCurrentPosition() == g.World.GetCurrentZone().GetSceneGraph().GetOrignId() {
		if g.Player.GetCurrentPosition() == g.World.GetCurrentZone().GetSceneGraph().GetOrignId() {
			forwardTraversal = false
		}
		currentZoneId := g.World.GetCurrentZoneId()
		zoneId, exists := g.World.GetCurrentZone().GetLink(g.Player.GetCurrentPosition())
		if exists {
			if zoneId == currentZoneId {
				return
			}
			g.World.SetCurrentZone(zoneId)
		} else {
			newZoneId := g.World.GetZoneCount() + 1
			g.World.AddZone(newZoneId, 0, g.Player.GetCurrentPosition(), currentZoneId, forwardTraversal)
			g.World.SetCurrentZone(newZoneId)
		}
		if forwardTraversal {
			g.Player.SetCurrentPosition(0)
		} else {
			g.Player.SetCurrentPosition(g.World.GetCurrentZone().GetSceneGraph().GetTerminusNodeId())
		}
	}
}

func GetRandomNumer() int {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(100) % 3
	return randomNumber
}
