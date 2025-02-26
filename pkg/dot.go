package pkg

import (
	"github.com/BognaLew/Pacman/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	DotFrameCount = 8

	DotPoints    = 10
	BigDotPoints = 50
)

type Dot struct {
	entity *Entity

	bigDot bool
	points int
}

func NewDot(position Vector) *Dot {
	var sprite *ebiten.Image
	var points int
	sprite = assets.DotImage
	points = DotPoints
	entity := NewEntity(position, sprite, 0)

	return &Dot{
		entity: entity,
		bigDot: false,
		points: points,
	}
}

func (dot *Dot) SetBigDot() {
	dot.points = BigDotPoints
	dot.entity.sprite = assets.BigDotImage
}
