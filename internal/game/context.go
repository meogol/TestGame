package game

import (
	graphics "github.com/quasilyte/ebitengine-graphics"
	input "github.com/quasilyte/ebitengine-input"
	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/gscene"
)

type Context struct {
	Input  *input.Handler
	Loader *resource.Loader
	Rand   gmath.Rand

	WindowWidth  int
	WindowHeight int

	graphicsCache *graphics.Cache

	scene gscene.GameRunner
}

func ChangeScene[ControllerAccessor any](ctx *Context, c gscene.Controller[ControllerAccessor]) {
	s := gscene.NewRootScene[ControllerAccessor](c)
	ctx.scene = s
}

func NewContext() *Context {
	return &Context{
		graphicsCache: graphics.NewCache(),
	}
}

func (ctx *Context) CurrentScene() gscene.GameRunner {
	return ctx.scene
}
