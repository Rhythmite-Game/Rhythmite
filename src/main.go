package main

import (
	"image/color"
	"log"

	"github.com/Rhythmite-Game/Rhythmite/classes"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 640
	screenHeight = 370
)

var (
	player *classes.Entity
)

func init() {
	player = &classes.Entity{
		// Data
		Name:      "Player",
		PositionX: screenWidth / 2,
		PositionY: screenHeight / 2,

		// Stats
		Health:            10,
		BaseAttack:        1,
		AttackMultiplier:  1,
		BaseDefense:       0,
		DefenseMultiplier: 1,

		// Movement
		TileSplits:     0,
		LegalDiagonals: false,
	}
}

type Game struct {
	keys []ebiten.Key
}

func (g *Game) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
	vector.DrawFilledCircle(screen, player.PositionX, player.PositionY, 20, color.RGBA{255, 50, 50, 100}, false)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Rhythmite")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
