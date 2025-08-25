package classes

import (
	"slices"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (player *Entity) Move(tileSize float32) {
	// Basic movement
	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		player.PositionY -= tileSize * float32(player.SpeedMultiplier)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		player.PositionX -= tileSize * float32(player.SpeedMultiplier)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		player.PositionY += tileSize * float32(player.SpeedMultiplier)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		player.PositionX += tileSize * float32(player.SpeedMultiplier)
	}

	// Diagonal movement
	if player.LegalDiagonals {
		if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
			player.PositionX -= tileSize * float32(player.SpeedMultiplier)
			player.PositionY -= tileSize * float32(player.SpeedMultiplier)
		} else if inpututil.IsKeyJustPressed(ebiten.KeyE) {
			player.PositionX += tileSize * float32(player.SpeedMultiplier)
			player.PositionY -= tileSize * float32(player.SpeedMultiplier)
		} else if inpututil.IsKeyJustPressed(ebiten.KeyZ) {
			player.PositionX -= tileSize * float32(player.SpeedMultiplier)
			player.PositionY += tileSize * float32(player.SpeedMultiplier)
		} else if inpututil.IsKeyJustPressed(ebiten.KeyC) {
			player.PositionX += tileSize * float32(player.SpeedMultiplier)
			player.PositionY += tileSize * float32(player.SpeedMultiplier)
		}
	}
}

func (player Entity) Dash(keyNames []string) {
	if slices.Contains(keyNames, "Shift") && player.DashEnabled {
		player.SpeedMultiplier = 2
	} else {
		player.SpeedMultiplier = 1
	}
}
