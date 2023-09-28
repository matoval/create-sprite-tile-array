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
	windowWidth         = 640
	windowHeight        = 480
	ticksPerSecond      = 10
	spriteFile          = "Fox Sprite Sheet.png"
	spriteWidth         = 448
	spriteHeight        = 224
	tileWidth           = 32
	tileHeight          = 32
	outsideScreenWidth  = 320
	outsideScreenHeight = 240
	startTile           = 28
	tileCount           = 7
)

var img *ebiten.Image

func init() {
	var err error
	img, _, err = ebitenutil.NewImageFromFile(spriteFile)
	if err != nil {
		log.Fatal(err)
	}
}

type TileSize struct {
	height int
	width  int
}

type SpriteSize struct {
	height int
	width  int
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
	if spriteSize.width%tileSize.width != 0 {
		return nil, fmt.Errorf("TileSize width of %v doesn't evenly fit in the sprint sheet width of %v", tileSize.height, spriteSize.height)
	}
	// Frist loop loops over each row of the sprite sheet
	for i := 0; i < spriteSize.height/tileSize.height; i++ {
		// Second loop loops over each column of the sprite sheet
		for j := 0; j < spriteSize.width/tileSize.width; j++ {
			sheet = append(sheet, image.Rectangle{image.Pt(j*tileSize.width, i*tileSize.height), image.Pt((j+1)*tileSize.width, (i+1)*tileSize.height)})
		}
	}
	return sheet, nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	spriteSheetArr, err := SpriteSheet(
		TileSize{height: tileHeight, width: tileWidth},
		SpriteSize{height: spriteHeight, width: spriteWidth},
	)
	if err != nil {
		log.Fatal(err)
	}
	// Centers the image on the screen
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(tileWidth)/2, -float64(tileHeight)/2)
	op.GeoM.Translate(outsideScreenWidth/2, outsideScreenHeight/2)
	// Draws tile to the screen starting with the tile with the index of startTile and continues until
	// tileCount has been added to startTile then it resets
	screen.DrawImage(img.SubImage(spriteSheetArr[((g.count)%tileCount)+startTile]).(*ebiten.Image), op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideScreenWidth, outsideScreenHeight
}

func main() {
	ebiten.SetTPS(ticksPerSecond)
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Animation!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
