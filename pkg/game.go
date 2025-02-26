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
	dots   []*Dot

	gameOver bool
	count    int
}

func NewGame() *Game {
	board := NewBoard()
	dots := board.PrepareBoard()
	game := &Game{
		board:    board,
		player:   NewPlayer(*board.pacmanSpawn.position.Multiply(32)),
		ghost:    NewGhost(board.ghostSpawnPositions[0], assets.BlinkyImage),
		dots:     dots,
		gameOver: false,
		count:    0,
	}

	return game
}

func (game *Game) checkColision() {
	if game.player.CheckColision(game.ghost.entity.colider) {
		game.player.entity.alive = false
		game.gameOver = true
		return
	}
	for idx, dot := range game.dots {
		if game.player.CheckColision(dot.entity.colider) {
			game.player.AddPoints(dot.points)
			game.dots = append(game.dots[:idx], game.dots[idx+1:]...)
		}
	}
}

func (game *Game) Update() error {
	if !game.gameOver {
		game.count += 1
		game.player.Update(*game.board)
		game.ghost.Update(*game.board)
		game.checkColision()
	}
	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	game.board.Draw(screen)
	for _, dot := range game.dots {
		dot.entity.Draw(screen, game.count, DotFrameCount)
	}
	game.ghost.Draw(screen, game.count)
	if !game.gameOver {
		game.player.Draw(screen, game.count)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
