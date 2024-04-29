package game

import (
	"fmt"
	"image/color"

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
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(1, 1.422)
	rect := ebiten.NewImage(16, 720)
	rect.Fill(color.White)

	op2 := &ebiten.DrawImageOptions{}
	op2.GeoM.Scale(1.422, 1)
	rect2 := ebiten.NewImage(720, 16)
	rect2.Fill(color.White)

	op3 := &ebiten.DrawImageOptions{}
	op3.GeoM.Scale(16, 1.422)
	op3.GeoM.Translate(1008, 0)
	rect3 := ebiten.NewImage(1, 720)
	rect3.Fill(color.White)

	op4 := &ebiten.DrawImageOptions{}
	op4.GeoM.Scale(1.422, 1)
	op4.GeoM.Translate(0, 1008)
	rect4 := ebiten.NewImage(720, 16)
	rect4.Fill(color.White)

	g.camera.clear()
	g.camera.draw(assets.Tile, &ebiten.DrawImageOptions{})

	g.player.Draw(screen, &g.camera, g)

	g.camera.draw(rect, op)
	g.camera.draw(rect2, op2)
	g.camera.draw(rect3, op3)
	g.camera.draw(rect4, op4)

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

	// if g.player.player.x <= 3150 && g.player.player.y <= 2350 {
	// 	g.camera.setPos(0, g.player.player.y/unit-240)
	// 	g.camera.setPos(0, 0)
	// }

	// if g.player.player.x >= 7040 && g.player.player.y >= 7845 {
	// 	g.camera.setPos(385, g.player.player.y/unit-240)
	// 	g.camera.setPos(g.player.player.x/unit, 240)
	// }

	vsync := ebiten.IsVsyncEnabled()

	if inpututil.IsKeyJustPressed(ebiten.KeyV) {
		ebiten.SetVsyncEnabled(!vsync)
	}

	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
