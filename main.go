package main

import (
	"github.com/BognaLew/Pacman/pkg"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := pkg.NewGame()

	err := ebiten.RunGame(game)
	if err != nil {
		panic(err)
	}
}
