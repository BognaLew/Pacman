package assets

import (
	"embed"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed *
var assetsEmbed embed.FS

var PacmanImage = LoadImage("PacMan.png")
var BlinkyImage = LoadImage("ghosts/redGhost.png")
var DotImage = LoadImage("Coin.png")
var BigDotImage = LoadImage("BigCoin.png")

func LoadImage(filename string) *ebiten.Image {
	f, err := assetsEmbed.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}
