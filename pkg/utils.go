package pkg

type Direction uint8

const (
	NO    Direction = 0
	UP    Direction = 1
	RIGHT Direction = 2
	DOWN  Direction = 3
	LEFT  Direction = 4
)

func (direction Direction) Opposite() Direction {
	switch direction {
	case LEFT:
		return RIGHT
	case UP:
		return DOWN
	case RIGHT:
		return LEFT
	case DOWN:
		return UP
	default:
		panic("Unknown direction")
	}
}

type Vector struct {
	X float64
	Y float64
}

func NewVector(x float64, y float64) *Vector {
	return &Vector{
		X: x,
		Y: y,
	}
}

func (vector *Vector) Modify(dx float64, dy float64) {
	vector.X += dx
	vector.Y += dy
}

func (vector *Vector) SetValues(x float64, y float64) {
	vector.X = x
	vector.Y = y
}

func (vector Vector) Equals(other Vector) bool {
	return vector.X == other.X && vector.Y == other.Y
}

func (vector Vector) Multiply(multiplicator float64) *Vector {
	return NewVector(vector.X*multiplicator, vector.Y*multiplicator)
}

func (vector Vector) CheckFullTile() bool {
	xRemainder, yRemainder := int(vector.X)%TileSize, int(vector.Y)%TileSize
	return xRemainder == 0 && yRemainder == 0
}

type Collider struct {
	position Vector
	width    float64
	height   float64
}

func NewCollider(position Vector, width float64, height float64) *Collider {
	return &Collider{
		position: position,
		width:    width,
		height:   height,
	}
}

func (colider *Collider) UpdatePosition(dx float64, dy float64) {
	colider.position.Modify(dx, dy)
}

func (colider Collider) GetMaxXY() (float64, float64) {
	return colider.position.X + colider.width, colider.position.Y + colider.height
}

func (colider Collider) CheckColision(other Collider) bool {
	otherMaxX, otherMaxY := other.GetMaxXY()
	coliderMaxX, coliderMaxY := colider.GetMaxXY()

	return colider.position.X < otherMaxX && other.position.X < coliderMaxX &&
		colider.position.Y < otherMaxY && other.position.Y < coliderMaxY
}
