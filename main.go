package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lidldev/NotNamedYet/game"
)

func main() {
	g := game.NewGame()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
