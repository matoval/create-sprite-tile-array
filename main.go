package main

import (
	"fmt"
	"image"
	"log"

	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	startTile = 28
	tileCount = 7
)

var img *ebiten.Image

func init() {
	var err error
	img, _, err = ebitenutil.NewImageFromFile("Fox Sprite Sheet.png")
	if err != nil {
		log.Fatal(err)
	}
}

type TileSize struct {
	height int
	length int
}

type SpriteSize struct {
	height int
	length int
}

type Game struct {
	count int
}

func (g *Game) Update() error {
	g.count++
	return nil
}

func SpriteSheet(tileSize TileSize, spriteSize SpriteSize) ([]image.Rectangle, error) {
	var sheet []image.Rectangle
	if spriteSize.height%tileSize.height != 0 {
		return nil, fmt.Errorf("TileSize height of %v doesn't evenly fit in the sprint sheet height of %v", tileSize.height, spriteSize.height)
	}
	if spriteSize.length%tileSize.length != 0 {
		return nil, fmt.Errorf("TileSize length of %v doesn't evenly fit in the sprint sheet length of %v", tileSize.height, spriteSize.height)
	}
	for i := 0; i < spriteSize.height/tileSize.height; i++ {
		for j := 0; j < spriteSize.length/tileSize.length; j++ {
			sheet = append(sheet, image.Rectangle{image.Pt(j*tileSize.length, i*tileSize.height), image.Pt((j+1)*tileSize.length, (i+1)*tileSize.height)})
		}
	}

	return sheet, nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	spriteSheetArr, err := SpriteSheet(
		TileSize{height: 32, length: 32},
		SpriteSize{height: 224, length: 448},
	)
	if err != nil {
		log.Fatal(err)
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(32)/2, -float64(32)/2)
	op.GeoM.Translate(320/2, 240/2)
	screen.DrawImage(img.SubImage(spriteSheetArr[((g.count/6)%tileCount)+startTile]).(*ebiten.Image), op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
