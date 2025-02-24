package pkg

import (
	"github.com/BognaLew/Pacman/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 544
	ScreenHeight = 544
)

type Game struct {
	board *Board

	player *Player
	ghost  *Ghost

	gameOver bool
	count    int
}

func NewGame() *Game {
	board := NewBoard()
	board.PrepareBoard()
	game := &Game{
		board:    board,
		player:   NewPlayer(*board.pacmanSpawn.position.Multiply(32)),
		ghost:    NewGhost(board.ghostSpawnPositions[0], assets.BlinkyImage),
		gameOver: false,
		count:    0,
	}

	return game
}

func (game *Game) checkColision() {
	colided := game.player.CheckColision(game.ghost.entity.colider)
	if colided {
		game.player.entity.alive = false
		game.gameOver = true
	}
}

func (game *Game) Update() error {
	game.count += 1
	game.player.Update(*game.board)
	game.ghost.Update(*game.board)
	game.checkColision()
	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	game.board.Draw(screen)
	if !game.gameOver {
		game.player.Draw(screen, game.count)
		game.ghost.Draw(screen, game.count)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
