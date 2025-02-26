package pkg

import (
	"fmt"

	"github.com/BognaLew/Pacman/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	PacmanFrameCount   = 8
	InitialPacmanSpeed = 2
)

type Player struct {
	entity *Entity

	nextDirection Direction

	score int
}

func NewPlayer(position Vector) *Player {
	player := &Player{
		entity:        NewEntity(position, assets.PacmanImage, InitialPacmanSpeed),
		nextDirection: NO,
		score:         0,
	}

	return player
}

func (player *Player) Update(board Board) {
	fullTile := player.entity.position.CheckFullTile()
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		if player.canChangeDirection(UP, board) && fullTile {
			player.entity.ChangeDirection(UP)
			player.nextDirection = NO
		} else {
			player.nextDirection = UP
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		if player.canChangeDirection(RIGHT, board) && fullTile {
			player.entity.ChangeDirection(RIGHT)
			player.nextDirection = NO
		} else {
			player.nextDirection = RIGHT
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		if player.canChangeDirection(DOWN, board) && fullTile {
			player.entity.ChangeDirection(DOWN)
			player.nextDirection = NO
		} else {
			player.nextDirection = DOWN
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		if player.canChangeDirection(LEFT, board) && fullTile {
			player.entity.ChangeDirection(LEFT)
			player.nextDirection = NO
		} else {
			player.nextDirection = LEFT
		}
	} else if fullTile && player.canChangeDirection(player.nextDirection, board) && player.nextDirection != NO {
		player.entity.ChangeDirection(player.nextDirection)
		player.nextDirection = NO
	}
	player.entity.move(board)
}

func (player *Player) Draw(screen *ebiten.Image, count int) {
	player.entity.Draw(screen, count, PacmanFrameCount)
}

func (player *Player) CheckColision(other Collider) bool {
	return player.entity.colider.CheckColision(other)
}

func (player Player) canChangeDirection(direction Direction, board Board) bool {
	switch direction {
	case UP:
		return board.GetTileTypeAtPosition(*NewVector(player.entity.centre.X, player.entity.centre.Y-TileSize)) != WALL
	case RIGHT:
		return board.GetTileTypeAtPosition(*NewVector(player.entity.centre.X+TileSize, player.entity.centre.Y)) != WALL
	case DOWN:
		return board.GetTileTypeAtPosition(*NewVector(player.entity.centre.X, player.entity.centre.Y+TileSize)) != WALL
	case LEFT:
		return board.GetTileTypeAtPosition(*NewVector(player.entity.centre.X-TileSize, player.entity.centre.Y)) != WALL
	}
	return false
}

func (player *Player) AddPoints(points int) {
	player.score += points
	fmt.Println(player.score)
}
