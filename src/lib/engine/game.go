package engine

import (
	"fmt"
	"go-rogue/src/lib/components"
	"go-rogue/src/lib/entities"
	"go-rogue/src/lib/interfaces"
	"go-rogue/src/lib/maps"
	"go-rogue/src/lib/scenes"
	"go-rogue/src/lib/userInterface"
	"math/rand"
	"os"
	"strconv"
	"time"
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
	userInterface.DrawTitleText("Go Rogue")
	// Initial Map Generation
	fmt.Println("Generating map...")
	g.World.AddZone(0, 0, 0, 0, true)
	fmt.Println("Writting map...")
	WriteDotFile("graph.dot", g.World.GetCurrentZone().GetSceneGraph())

	ticker := time.NewTicker(time.Duration(g.TickRate*1000) * time.Millisecond)
	defer ticker.Stop()

	var input string
	fmt.Printf("\033[13;50HPress Enter to continue or type 'exit' to quit: ")
	fmt.Scanln(&input)

	// Check if the user wants to exit
	if input == "exit" {
		os.Exit(0)
	}

	for range ticker.C {
		userInterface.DrawPlayerAttributes(g.Player)
		scenes.RenderScene(
			g.GameScenes.GetScene(GetRandomNumer()),
		)

		fmt.Printf("\033[13;50HPress Enter to continue or type 'exit' to quit: ")
		//currentPosition := g.Player.GetCurrentPosition()
		movementOptions := g.Player.GetMovementOptions(g.World.GetCurrentZone().GetSceneGraph())
		fmt.Printf("\033[13;50HPress the number to enter the room: %d                 ", movementOptions.Keys())
		fmt.Scanln(&input)
		i, err := strconv.Atoi(input)
		if err != nil {
			fmt.Printf("\033[13;50HPress Enter to continue or type 'exit' to quit: ")
			continue
		}
		if movementOptions.Contains(i) {
			g.Player.SetCurrentPosition(i)
		}

		// Move between zones
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
					continue
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
		// Start combat Test
		// fmt.Printf("\033[13;50HStarting combat... %s", userInterface.Spaces(80))
		//g.Combat.Attack(g.Player, g.Enemy)
	}
}

func GetRandomNumer() int {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(100) % 3
	return randomNumber
}

func WriteDotFile(filename string, sceneGraph *maps.SceneGraph) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Start the DOT graph
	_, err = file.WriteString(fmt.Sprintf("graph G {\n  label=\"%s\";\n  labelloc=\"t\";\n  fontsize=\"20\";\n", sceneGraph.GetTheme().Name))
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	// Write nodes with labels and colors
	for _, node := range sceneGraph.GetAllNodes() {
		nodeMetaData := node.GetMetaData()
		_, err = file.WriteString(fmt.Sprintf("  %d [label=\"%s\", color=\"%s\"];\n", node.GetId(), nodeMetaData.Label, nodeMetaData.Color))
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	// Write edges
	for _, node := range sceneGraph.GetAllNodes() {
		for neighbor := range node.GetAllEdges() {
			if node.GetId() < neighbor { // Avoid duplicate edges
				edgeMetaData := node.GetEdge(neighbor).GetMetaData()
				_, err = file.WriteString(fmt.Sprintf("  %d -- %d [label=\"%s\", color=\"%s\", style=\"%s\", penwidth=\"%d\"];\n", node.GetId(), neighbor, edgeMetaData.Name, edgeMetaData.Color, edgeMetaData.Style, edgeMetaData.Width))
				if err != nil {
					fmt.Println("Error writing to file:", err)
					return
				}
			}
		}
	}

	// End the DOT graph
	_, err = file.WriteString("}\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Graph written to", filename)
}
