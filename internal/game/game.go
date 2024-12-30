package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	input "github.com/quasilyte/ebitengine-input"
)

type Game struct {
	InputSystem input.System
	Ctx         *Context
}

func (g *Game) Update() error {
	g.InputSystem.Update()
	g.Ctx.CurrentScene().UpdateWithDelta(1.0 / 60.0)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Ctx.CurrentScene().Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Ctx.WindowWidth, g.Ctx.WindowHeight
}
