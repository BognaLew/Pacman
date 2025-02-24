package pkg

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	FrameWidth  = 16
	FrameHeight = 16
	FrameOX     = 0
	FrameOY     = 0
)

type Entity struct {
	position          Vector
	centre            Vector
	speed             float64
	direction         Direction
	rotation          float64
	rotationTransform Vector
	colider           Collider
	alive             bool

	sprite *ebiten.Image
}

func NewEntity(position Vector, sprite *ebiten.Image, speed float64) *Entity {
	centre := *NewVector(position.X+TileSize/2, position.Y+TileSize/2)
	entity := &Entity{
		position:          position,
		centre:            centre,
		speed:             speed,
		direction:         NO,
		rotation:          0,
		rotationTransform: *NewVector(0, 0),
		colider:           *NewCollider(position, FrameWidth*2, FrameHeight*2),
		alive:             true,

		sprite: sprite,
	}

	return entity
}

func (entity *Entity) Draw(screen *ebiten.Image, count int, frameCount int) {
	op := &ebiten.DrawImageOptions{}

	var yScale float64 = 2
	if entity.rotation == math.Pi {
		yScale = -2
	}

	op.GeoM.Rotate(entity.rotation)
	op.GeoM.Scale(2, yScale)
	op.GeoM.Translate(entity.rotationTransform.X, entity.rotationTransform.Y)
	op.GeoM.Translate(entity.position.X, entity.position.Y)

	i := (count / 5) % frameCount
	sx, sy := FrameOX+i*FrameWidth, FrameOY
	screen.DrawImage(entity.sprite.SubImage(
		image.Rect(sx, sy, sx+FrameWidth, sy+FrameHeight)).(*ebiten.Image),
		op)
}

func (entity *Entity) move(board Board) {
	var dx, dx2, dy, dy2 float64
	switch entity.direction {
	case UP:
		dy = -entity.speed
		dy2 = dy
	case RIGHT:
		dx = entity.speed
		dx2 = TileSize
	case DOWN:
		dy = entity.speed
		dy2 = TileSize
	case LEFT:
		dx = -entity.speed
		dx2 = dx
	default:
		return
	}

	x, y := entity.position.X+dx2, entity.position.Y+dy2
	if board.GetTileTypeAtPosition(*NewVector(x, y)) != WALL {
		entity.position.Modify(dx, dy)
		entity.centre.Modify(dx, dy)
		entity.colider.UpdatePosition(dx, dy)
	}
}

func (entity *Entity) ChangeDirection(direction Direction) {
	switch direction {
	case UP:
		entity.direction = UP
		entity.rotation = math.Pi * 3 / 2
		entity.rotationTransform.SetValues(0, 2*FrameHeight)
	case RIGHT:
		entity.direction = RIGHT
		entity.rotation = 0
		entity.rotationTransform.SetValues(0, 0)
	case DOWN:
		entity.direction = DOWN
		entity.rotation = math.Pi / 2
		entity.rotationTransform.SetValues(2*FrameWidth, 0)
	case LEFT:
		entity.direction = LEFT
		entity.rotation = math.Pi
		entity.rotationTransform.SetValues(2*FrameWidth, 0)

	}
}
