package pkg

import (
	"bytes"
	"fmt"
	"image/color"

	"github.com/BognaLew/Pacman/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	ScreenWidth  = 544
	ScreenHeight = 544
)

type Game struct {
	board *Board

	source *text.GoTextFaceSource

	player *Player
	ghost  *Ghost
	dots   []*Dot

	gameOver bool
	count    int
}

func NewGame() *Game {
	board := NewBoard()
	dots := board.PrepareBoard()

	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		panic(err)
	}
	source := s

	game := &Game{
		board:    board,
		source:   source,
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
	if len(game.dots) == 0 {
		game.gameOver = true
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

func (game *Game) drawText(screen *ebiten.Image, msg string, position Vector, fontSize float64) {
	op := &text.DrawOptions{}
	op.GeoM.Translate(position.X, position.Y)
	op.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, msg, &text.GoTextFace{
		Source: game.source,
		Size:   fontSize,
	}, op)
}

func (game *Game) Draw(screen *ebiten.Image) {
	game.board.Draw(screen)
	game.drawText(screen, fmt.Sprintf("Score: %d", game.player.score), *NewVector(10, 5), 22)
	for _, dot := range game.dots {
		dot.entity.Draw(screen, game.count, DotFrameCount)
	}
	game.ghost.Draw(screen, game.count)
	if !game.gameOver {
		game.player.Draw(screen, game.count)
	} else {
		var msg string
		if msg = "Game Over"; game.player.entity.alive {
			msg = "You won!"
		}
		game.drawText(screen, msg, *NewVector(170, 260), 40)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
