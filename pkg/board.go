package pkg

import (
	"image/color"
	"math/rand"

	"github.com/BognaLew/Pacman/maps"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	TilesPerX = 17
	TilesPerY = 17

	TileSize = 32
)

var PathColor = color.RGBA{0, 0, 0, 255}
var WallColor = color.RGBA{30, 30, 30, 255}

type TileType uint8

const (
	NONE         TileType = 0
	WALL         TileType = 1
	GHOST_SPAWN  TileType = 2
	PACMAN_SPAWN TileType = 3
)

type Tile struct {
	tileType TileType
	position Vector

	sprite *ebiten.Image
}

func NewTile(tileType TileType, position Vector, tileColor color.RGBA) *Tile {
	sprite := ebiten.NewImage(TileSize, TileSize)
	vector.DrawFilledRect(sprite, 0, 0, TileSize, TileSize, tileColor, false)

	return &Tile{
		tileType: tileType,
		position: position,
		sprite:   sprite,
	}
}

func (tile *Tile) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(tile.position.X*TileSize, tile.position.Y*TileSize)

	screen.DrawImage(tile.sprite, op)
}

type Board struct {
	tiles [][]*Tile

	pacmanSpawn         *Tile
	ghostSpawnPositions []Vector
}

func NewBoard() *Board {
	return &Board{
		tiles:               make([][]*Tile, 0),
		pacmanSpawn:         nil,
		ghostSpawnPositions: make([]Vector, 0),
	}
}

func (board *Board) Draw(screen *ebiten.Image) {
	for _, row := range board.tiles {
		for _, tile := range row {
			tile.Draw(screen)
		}
	}
}

func (board *Board) parseBoard(boardTemplate [][]uint8) []*Dot {
	tiles := make([][]*Tile, 0)
	dots := make([]*Dot, 0)

	for yIdx, row := range boardTemplate {
		tilesInRow := make([]*Tile, 0)
		for xIdx, item := range row {
			position := *NewVector(float64(xIdx), float64(yIdx))
			switch item {
			case 0:
				tilesInRow = append(tilesInRow, NewTile(NONE, position, PathColor))
				dots = append(dots, NewDot(*position.Multiply(TileSize)))
			case 1:
				tilesInRow = append(tilesInRow, NewTile(WALL, position, WallColor))
			case 2:
				tilesInRow = append(tilesInRow, NewTile(GHOST_SPAWN, position, PathColor))
				board.ghostSpawnPositions = append(board.ghostSpawnPositions, *position.Multiply(TileSize))
			case 3:
				tile := NewTile(PACMAN_SPAWN, position, PathColor)
				tilesInRow = append(tilesInRow, NewTile(PACMAN_SPAWN, position, PathColor))
				board.pacmanSpawn = tile
			default:
				panic("Unknown tile type")
			}
		}
		tiles = append(tiles, tilesInRow)
	}

	board.tiles = tiles
	return dots
}

func (board *Board) PrepareBoard() []*Dot {
	dots := board.parseBoard(maps.LoadMap1())
	numOfBigDots, numOfDots := 4, len(dots)
	for numOfBigDots != 0 {
		idx := rand.Intn(numOfDots)
		if !dots[idx].bigDot {
			dots[idx].SetBigDot()
			numOfBigDots -= 1
		}
	}

	return dots
}

func (board Board) GetTileTypeAtPosition(position Vector) TileType {
	x, y := int(position.X/TileSize), int(position.Y/TileSize)

	return board.tiles[y][x].tileType
}

func (board Board) GetAvailableDirections(position Vector) []Direction {
	directions := make([]Direction, 0)
	ghostSpawnDirections := make([]Direction, 0)
	x, y := int(position.X/TileSize), int(position.Y/TileSize)
	xDim, yDim := len(board.tiles[0]), len(board.tiles)

	if x > 0 {
		tileType := board.tiles[y][x-1].tileType
		if tileType == NONE {
			directions = append(directions, LEFT)
		} else if tileType == GHOST_SPAWN {
			ghostSpawnDirections = append(ghostSpawnDirections, LEFT)
		}
	}
	if x < xDim-1 {
		tileType := board.tiles[y][x+1].tileType
		if tileType == NONE {
			directions = append(directions, RIGHT)
		} else if tileType == GHOST_SPAWN {
			ghostSpawnDirections = append(ghostSpawnDirections, RIGHT)
		}
	}

	if y > 0 {
		tileType := board.tiles[y-1][x].tileType
		if tileType == NONE {
			directions = append(directions, UP)
		} else if tileType == GHOST_SPAWN {
			ghostSpawnDirections = append(ghostSpawnDirections, UP)
		}
	}
	if y < yDim-1 {
		tileType := board.tiles[y+1][x].tileType
		if tileType == NONE {
			directions = append(directions, DOWN)
		} else if tileType == GHOST_SPAWN {
			ghostSpawnDirections = append(ghostSpawnDirections, DOWN)
		}
	}

	if len(directions) == 0 {
		return ghostSpawnDirections
	}

	return directions
}
