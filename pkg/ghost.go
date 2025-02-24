package pkg

import (
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	GhostFrameCount   = 8
	InitialGhostSpeed = 1
)

type Ghost struct {
	entity Entity
}

func NewGhost(position Vector, sprite *ebiten.Image) *Ghost {
	entity := *NewEntity(position, sprite, InitialGhostSpeed)
	ghost := &Ghost{
		entity: entity,
	}

	return ghost
}

func (ghost *Ghost) Update(board Board) {
	fullTile := ghost.entity.position.CheckFullTile()
	availableDirections := board.GetAvailableDirections(ghost.entity.centre)
	if ghost.entity.direction == NO {
		idx := rand.IntN(len(availableDirections))
		ghost.entity.ChangeDirection(availableDirections[idx])
		ghost.entity.move(board)
		return
	}

	if fullTile {
		newDirections := make([]Direction, 0)
		currentDirectionInArray := false
		for _, value := range availableDirections {
			if value != ghost.entity.direction && value != ghost.entity.direction.Opposite() {
				newDirections = append(newDirections, value)
			}
			if value == ghost.entity.direction {
				currentDirectionInArray = true
			}
		}
		if len(newDirections) != 0 {
			randomDirection := newDirections[rand.IntN(len(newDirections))]
			if rand.IntN(100) > 55 {
				ghost.entity.ChangeDirection(randomDirection)
			}
		} else if len(newDirections) == 0 && !currentDirectionInArray {
			ghost.entity.ChangeDirection(ghost.entity.direction.Opposite())
		}
	}
	ghost.entity.move(board)
}

func (ghost *Ghost) Draw(screen *ebiten.Image, count int) {
	ghost.entity.Draw(screen, count, GhostFrameCount)
}
