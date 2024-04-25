package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/lidldev/NotNamedYet/assets"
)

type Game struct {
	camera camera
	player Player
}

func NewGame() *Game {
	g := &Game{}
	g.camera.init()
	return g
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.camera.clear()
	g.camera.draw(assets.Tile, &ebiten.DrawImageOptions{})

	g.player.Draw(screen, &g.camera, g)

	g.camera.render(screen)

	msg2 := fmt.Sprintf("PosX: %d PosY: %d",
		g.player.player.x,
		g.player.player.y,
	)
	ebitenutil.DebugPrint(screen, msg2)
}

func (g *Game) Update() error {
	g.player.Update()
	g.camera.setPos(g.player.player.x/unit-320, g.player.player.y/unit-240)

	if g.player.player.x <= 3150 {
		g.camera.setPos(0, g.player.player.y/unit-240)
	}

	if g.player.player.x >= 7040 {
		g.camera.setPos(385, g.player.player.y/unit-240)
	}

	if g.player.player.y <= 2350 {
		g.camera.setPos(0, 0)
	}

	if g.player.player.y >= 7845 {
		g.camera.setPos(g.player.player.x/unit, 240)
	}

	vsync := ebiten.IsVsyncEnabled()

	if inpututil.IsKeyJustPressed(ebiten.KeyV) {
		ebiten.SetVsyncEnabled(!vsync)
	}

	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
