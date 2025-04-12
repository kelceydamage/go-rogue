package main

import (
	"go-rogue/src/lib/engine"
	"go-rogue/src/lib/entities"
)

func main() {
	// Example initialization
	player := entities.NewPlayer()
	enemy := entities.NewEnemy()
	game := engine.NewGame(
		player,
		enemy,
		0.01,
	)
	game.Run()
}
